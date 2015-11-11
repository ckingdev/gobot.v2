package bot

import (
	"math"
	"time"

	"github.com/cpalone/gobot.v2/proto"

	hproto "euphoria.io/heim/proto"
	"euphoria.io/scope"
)

const maxRetries = 10

// RetryConn is a decorator on a proto.Connection that implements retrying after
// an error with exponential backoff.
type RetryConn struct {
	ctx  scope.Context
	conn proto.Connection
}

// Connect tries to connect, and if there is a failure, waits .5, 1, 2, 4...etc
// seconds before trying to connect again.
func (rc *RetryConn) Connect() error {
	rc.ctx.WaitGroup().Add(1)
	defer rc.ctx.WaitGroup().Done()
	var err error
	if err = rc.conn.Connect(); err == nil {
		return nil
	}
	for i := 0; i < maxRetries; i++ {
		time.Sleep(time.Duration(math.Pow(2.0, float64(i-1))) * time.Second)
		if err = rc.conn.Connect(); err == nil {
			return nil
		}
	}
	GetLogger(rc.ctx).Errorf("RetryConn: Connect: max retries reached: %s", err)
	return err
}

// Run wraps the underlying Connection's Run by looping over calls to Run if the
// Connection requests a restart. Exits with a nil error do not trigger a restart.
func (rc *RetryConn) Run() error {
	rc.ctx.WaitGroup().Add(1)
	defer rc.ctx.WaitGroup().Done()
	errChan := make(chan error)
	for {
		rc.ctx.WaitGroup().Add(1)
		go func() {
			defer rc.ctx.WaitGroup().Done()
			errChan <- rc.conn.Run()
		}()
		select {
		case err := <-errChan:
			// if Run exited cleanly, no need to restart
			if err == nil {
				return nil
			}
			// for a restart of Run to occur, the Connection must explicitly
			// specify proto.ErrRequestRestart. Otherwise, we return the error.
			if err == proto.ErrRequestRestart {
				GetLogger(rc.ctx).Infoln("RetryConn: Connection requested restart. Attemping...")
				rc.ctx.WaitGroup().Add(1)
				go func() {
					defer rc.ctx.WaitGroup().Done()
					errChan <- rc.conn.Run()
				}()
				continue
			}
			GetLogger(rc.ctx).Infof("RetryConn: Connection exited with error: %s. Exiting.", err)
			return err
		case <-rc.ctx.Done():
			GetLogger(rc.ctx).Debugln("RetryConn: Run: context is done, exiting.")
			return nil
		}
	}
}

// Kill calls the underlying Connection's Kill function, cancels its own context,
// and waits for its goroutines to exit.
func (rc *RetryConn) Kill() error {
	if err := rc.conn.Kill(); err != nil {
		GetLogger(rc.ctx).Errorf("RetryConn: Kill: error killing underlying Connection: %s", err)
		return err
	}
	rc.ctx.Cancel()
	rc.ctx.WaitGroup().Wait()
	return nil
}

// Incoming simply passes the underlying Connection's Incoming().
func (rc *RetryConn) Incoming() <-chan *hproto.Packet {
	return rc.conn.Incoming()
}

// Outgoing simply passes the underlying Connection's Outgoing().
func (rc *RetryConn) Outgoing() chan<- *hproto.Packet {
	return rc.conn.Outgoing()
}
