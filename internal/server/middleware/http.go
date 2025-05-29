package middleware

import (
	"context"
	"maps"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"time"

	"github.com/gaurishhs/dav-server/internal/server"
	"github.com/gaurishhs/gor"
	"github.com/rs/zerolog/log"
)

func Logger() gor.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := NewWrapResponseWriter(w, r.ProtoMajor)
			rec := httptest.NewRecorder()
			ctx := r.Context()
			path := r.URL.EscapedPath()
			reqData, _ := httputil.DumpRequest(r, false)
			logger := log.Log().Timestamp().Str("path", path).Bytes("request_data", reqData)

			defer func(begin time.Time) {
				status := ww.Status()
				tookMs := time.Since(begin).Milliseconds()

				logger.Int64("took", tookMs).Int("status_code", status).Msgf("[%d] %s http request for %s took %dms", status, r.Method, path, tookMs)
			}(time.Now())

			ctx = context.WithValue(ctx, server.LoggerCtxKey, logger)
			next.ServeHTTP(rec, r.WithContext(ctx))

			maps.Copy(ww.Header(), rec.Header())
			ww.WriteHeader(rec.Code)
			rec.Body.WriteTo(ww)
		})
	}
}
