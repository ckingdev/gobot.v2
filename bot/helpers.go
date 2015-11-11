package bot

import (
	"euphoria.io/scope"
	"github.com/Sirupsen/logrus"
)

type ctxKey int

const (
	loggerKey ctxKey = iota
)

func SetLogger(ctx scope.Context, logger *logrus.Logger) {
	ctx.Set(loggerKey, logger)
}

func GetLogger(ctx scope.Context) *logrus.Logger {
	return ctx.Get(loggerKey).(*logrus.Logger)
}