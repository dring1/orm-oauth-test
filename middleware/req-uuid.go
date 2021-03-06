package middleware

import (
	"context"
	"net/http"

	"github.com/dring1/jwt-oauth/lib/contextkeys"
	uuid "github.com/satori/go.uuid"
)

func AddUUID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		id := uuid.NewV4().String()
		ctx := context.WithValue(req.Context(), contextkeys.ReqId, id)
		req = req.WithContext(ctx)
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}
