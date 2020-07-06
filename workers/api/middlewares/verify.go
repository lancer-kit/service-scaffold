package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lancer-kit/armory/api/render"
)

type ContextKey string

const (
	SomeParam ContextKey = "some_param"
)

// VerifySomething is an example of custom middleware which checks parameter value from url
func VerifySomething() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mID := chi.URLParam(r, "mID")
			if mID != "test" {
				render.BadRequest(w, "Wrong param")
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), SomeParam, mID))
			next.ServeHTTP(w, r)
		})
	}
}
