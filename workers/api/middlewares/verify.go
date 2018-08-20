package middlewares

import (
	"net/http"
	"github.com/go-chi/chi"
	"context"
)

func VerifySomething() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lang := chi.URLParam(r, "some_param")

			r = r.WithContext(context.WithValue(r.Context(), "some_param", lang))
			next.ServeHTTP(w, r)
			return

		})
	}
}