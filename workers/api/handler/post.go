package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/models"
)

func AddBuzz(w http.ResponseWriter, r *http.Request) {
	data := new(models.BuzzFeed)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Default.Error(err)
		render.ServerError(w)
		return
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		log.Default.Error(err)
		render.ServerError(w)
		return
	}

	log.Default.Info("Trying to write data into database")
	dataQ := models.NewBuzzFeedQ(models.NewQ(nil))
	err = dataQ.Insert(*data)
	if err != nil {
		render.ServerError(w)
		log.Default.WithError(err).Error("Can not insert data into database")
		return
	}

	log.Default.Info("Data has been written successfully")
	render.Success(w, data)
}
