package handler

import (
"net/http"

"gitlab.inn4science.com/gophers/service-kit/api/render"
)

func GetValueFromMiddleware(w http.ResponseWriter, r *http.Request) {
	testParam := r.Context().Value("some_param")
	render.Success(w, testParam)
}
