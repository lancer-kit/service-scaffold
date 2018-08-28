package handler

import (
	"net/http"

	"encoding/json"
	"io/ioutil"

	"strconv"

	"github.com/go-chi/chi"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/models"
)

func ChangeBuzz(w http.ResponseWriter, r *http.Request) {
	type inputData struct {
		Description string `json:"description"`
	}

	uid := chi.URLParam(r, "id")
	idINT, err := strconv.Atoi(uid)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("wrong id").Render(w)
		return
	}

	data := new(inputData)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Bad body structure").Render(w)
		return
	}

	err = json.Unmarshal(body, data)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Bad body structure").Render(w)
		return
	}

	dataQ := models.NewBuzzFeedQ(models.NewQ(nil))
	err = dataQ.UpdateBuzzDescription(int64(idINT), data.Description)
	if err != nil {
		render.ResultNotFound.SetError("Cant found buzz with such user Id").Render(w)
		log.Default.WithError(err).Error("Can not insert data into database")
		return
	}

	log.Default.Info("Data has been written successfully")
	render.Success(w, data)
}
