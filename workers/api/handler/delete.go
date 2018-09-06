package handler

import (
	"net/http"

	"strconv"

	"github.com/go-chi/chi"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/models"
)

func DeleteBuzz(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	idINT, err := strconv.Atoi(uid)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	dataQ := models.NewBuzzFeedQ(models.NewQ(nil))
	err = dataQ.DeleteBuzzByID(int64(idINT))
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	log.Default.Info("Data has been deleted successfully")
	render.Success(w, "success")
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	idINT, err := strconv.Atoi(uid)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	docQ, err := models.CreateCustomDocumentQ()
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	err = docQ.DeleteDocument(idINT)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	render.Success(w, "Document was successfully deleted")
}
