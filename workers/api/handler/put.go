package handler

import (
	"net/http"

	"encoding/json"
	"io/ioutil"

	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/auth"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/models"
)

func Put(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value(auth.KeyUID).(int64)
	type inputData struct {
		Description string `json:"description"`
	}
	data := new(inputData)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		render.ServerError(w)
		return
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		render.ServerError(w)
		return
	}

	dataQ := models.NewBuzzFeedQ(models.NewQ(nil))
	err = dataQ.UpdateBuzzDescription(uid, data.Description)
	if err != nil {
		render.ServerError(w)
		log.Default.WithError(err).Error("Can not insert data into database")
		return
	}

	log.Default.Info("Data has been written successfully")
	render.Success(w, data)
}
