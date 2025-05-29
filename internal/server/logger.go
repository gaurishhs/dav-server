package server

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type contextKey struct {
	name string
}

func ContextKey(key string) contextKey {
	return contextKey{name: key}
}

func (k *contextKey) String() string {
	return "dav-server context value " + k.name
}

var LoggerCtxKey = &contextKey{"LogEntry"}

func FromContext(ctx context.Context) *zerolog.Event {
	logger, ok := ctx.Value(LoggerCtxKey).(*zerolog.Event)
	if ok {
		return logger
	}
	return log.Log().Timestamp()
}
