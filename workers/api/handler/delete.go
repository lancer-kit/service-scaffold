package handler

import (
	"net/http"

	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/auth"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(auth.KeyUID).(int64)

	dataQ := models.NewBuzzFeedQ(models.NewQ(nil))
	err := dataQ.DeleteBuzzByID(uid)
	if err != nil {
		render.BadRequest(w, err)
		return
	}
	log.Default.Info("Data has been deleted successfully")
	render.Success(w, "success")
}
