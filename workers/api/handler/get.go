package handler

import (
	"net/http"

	"gitlab.inn4science.com/gophers/service-kit/api/render"
)

func GetValueFromMiddleware(w http.ResponseWriter, r *http.Request) {
	type resStruct struct {
		value interface{} `json:"value"`
	}
	res := new(resStruct)
	testParam := r.Context().Value("some_param")
	res.value = testParam

	render.Success(w, res)
}
