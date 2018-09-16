package middlewares

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
)

// VerifySomething is an example of custom middleware which checks parameter value from url
func VerifySomething() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mId := chi.URLParam(r, "mId")
			if mId != "test" {
				render.BadRequest(w, "Wrong param")
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), "some_param", mId))
			next.ServeHTTP(w, r)
			return

		})
	}
}
