package middlewares

import (
	"net/http"
	"context"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
)

func VerifySomething() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			someParam := r.URL.Query().Get("param")
			if someParam != "test" {
				render.ResultNotFound.Render(w)
				return
			}
			r = r.WithContext(context.WithValue(r.Context(), "some_param", someParam))
			next.ServeHTTP(w, r)
			return

		})
	}
}