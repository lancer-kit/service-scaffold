package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/models"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	type inputData struct {
		Id int64 `json:"id"`
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
	err = dataQ.DeleteBuzzByID(data.Id)
	if err != nil {
		render.BadRequest(w, err)
		return
	}
	log.Default.Info("Data has been deleted successfully")
	render.Success(w, data)
}
