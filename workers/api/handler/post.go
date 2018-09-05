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
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	log.Default.Info("Trying to write data into database")
	dataQ := models.NewBuzzFeedQ(models.NewQ(nil))
	err = dataQ.Insert(*data)
	if err != nil {
		log.Default.WithError(err).Error("Can not insert data into database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	log.Default.Info("Data has been written successfully")
	render.WriteJSON(w, 201, data)
}

func AddDocument(w http.ResponseWriter, r *http.Request) {
	data := new(models.CustomDocument)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	log.Default.Info("Trying to write data into couchdb")
	docQ, err := models.CreateCustomDocumentQ()
	if err != nil {
		log.Default.WithError(err).Error("Can not establish connection with database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	err = docQ.AddDocument(data)
	if err != nil {
		log.Default.WithError(err).Error("Can not insert data into database")
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	log.Default.Info("Data has been written successfully")
	render.WriteJSON(w, 201, data)
}
