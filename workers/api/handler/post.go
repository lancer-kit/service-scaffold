package handler

import (
	"net/http"
	"encoding/json"
	"gitlab.inn4science.com/gophers/service-scaffold/models"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"io/ioutil"
)

func Post(w http.ResponseWriter, r *http.Request) {
	data := new(models.BuzzFeed)
	//err := json.NewDecoder(r.Body).Decode(data)
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

	log.Default.Info("Trying to write data into database")
	dataQ := models.NewBuzzFeedQ(models.NewQ(nil))
	err = dataQ.Insert(*data)
	if err != nil {
		render.ServerError(w)
		log.Default.WithError(err).Error("Can not insert data into database")
		return
	}

	log.Default.Info("Data has been written successfully")
	render.Success(w,data)
}
