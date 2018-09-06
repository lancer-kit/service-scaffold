package handler

import (
	"net/http"

	"strconv"

	"github.com/go-chi/chi"
	"gitlab.inn4science.com/gophers/service-kit/api/render"
	"gitlab.inn4science.com/gophers/service-kit/db"
	"gitlab.inn4science.com/gophers/service-kit/log"
	"gitlab.inn4science.com/gophers/service-scaffold/models"
)

func AllBuzz(w http.ResponseWriter, r *http.Request) {
	dbQuery := models.NewQ(nil).BuzzFeed()
	pageQuery, err := db.ParsePageQuery(r.URL.Query())
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	dbQuery.SetPage(&pageQuery)

	ols, err := dbQuery.Select()
	if err != nil || len(ols) == 0 {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	log.Default.Info("Buzz instances was successfully obtained")
	render.RenderListWithPages(w, pageQuery, int64(len(ols)), ols)
}

func GetBuzz(w http.ResponseWriter, r *http.Request) {
	uid := chi.URLParam(r, "id")
	idINT, err := strconv.Atoi(uid)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	dataQ := models.NewBuzzFeedQ(models.NewQ(nil))
	res, err := dataQ.ByID(int64(idINT))
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	log.Default.Info("Buzz instance was successfully obtained")
	render.Success(w, res)
}

func GetAllDocument(w http.ResponseWriter, r *http.Request) {
	pageQuery, err := db.ParsePageQuery(r.URL.Query())
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

	res, err := docQ.GetAllDocument(pageQuery)
	if err != nil {
		log.Default.Error(err)
		render.ResultNotFound.SetError("Not found").Render(w)
		return
	}

	log.Default.Info("All documents were successfully obtained from couchdb")
	render.RenderListWithPages(w, pageQuery, int64(len(res)), res)
}
