package middleware

import (
	"net/http"
	"runtime"

	"github.com/gaurishhs/dav-server/internal/server"
	"github.com/gaurishhs/gor"
)

func Recover() gor.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					if err == http.ErrAbortHandler {
						panic(err)
					}

					ctx := r.Context()
					logger := server.FromContext(ctx)
					stack := make([]byte, 8192)
					stack = stack[:runtime.Stack(stack, false)]
					logger.Bytes("stack", stack).Str("level", "fatal").Interface("error", err)
					hasRecovered, ok := ctx.Value(server.ContextKey("has_recovered")).(*bool)
					if ok {
						*hasRecovered = true
					}
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
