package models

import (
	"fmt"

	"github.com/lancer-kit/armory/db"
	cdb "github.com/leesper/couchdb-golang"
	"github.com/pkg/errors"

	"github.com/lancer-kit/service-scaffold/config"
)

type CustomDocument struct {
	Id         int64  `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	cdb.Document
}

type customDocumentQ struct {
	dbInstance *cdb.Database
}

func CreateCustomDocumentQ() (*customDocumentQ, error) {
	newDocInstance := new(customDocumentQ)

	dbInstance, err := cdb.NewDatabase(config.Config().CouchDB)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to couchdb")
	}

	newDocInstance.dbInstance = dbInstance
	return newDocInstance, nil
}

func (d *customDocumentQ) AddDocument(doc *CustomDocument) error {
	err := cdb.Store(d.dbInstance, doc)
	if err != nil {
		return errors.Wrap(err, "Unable to write into couchdb")
	}

	return nil
}

func (d *customDocumentQ) GetAllDocument(pQ db.PageQuery) ([]CustomDocument, error) {
	fields := []string{"id", "firstName", "secondName"}

	res, err := d.dbInstance.Query(fields, `exists(id,true)`, nil, int(pQ.PageSize), int(pQ.PageSize*(pQ.Page-1)), nil)

	if err != nil {
		return nil, errors.Wrap(err, "Unable to write into couchdb")
	}

	resSlice := make([]CustomDocument, 0)
	for _, v := range res {
		obj := new(CustomDocument)
		err = cdb.FromJSONCompatibleMap(obj, v)
		if err != nil {
			return nil, errors.Wrap(err, "failed to constructs a document JSON-map")
		}
		resSlice = append(resSlice, *obj)
	}

	return resSlice, nil
}

func (d *customDocumentQ) GetDocument(userID int) ([]CustomDocument, error) {
	fields := []string{"id", "firstName", "secondName"}

	res, err := d.dbInstance.Query(fields, fmt.Sprintf("id == %d", userID), nil, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to write into couchdb")
	}

	resSlice := make([]CustomDocument, 0)
	for _, v := range res {
		obj := new(CustomDocument)
		err = cdb.FromJSONCompatibleMap(obj, v)
		if err != nil {
			return nil, errors.Wrap(err, "failed to constructs a document JSON-map")
		}
		resSlice = append(resSlice, *obj)
	}

	return resSlice, nil

}

func (d *customDocumentQ) UpdateDocument(userID int, doc *CustomDocument) error {
	fields := []string{"_rev", "_id"}
	selector := fmt.Sprintf("id == %d", userID)
	res, err := d.dbInstance.Query(fields, selector, nil, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "Unable to read from couchdb")
	}

	obj := new(CustomDocument)
	err = cdb.FromJSONCompatibleMap(obj, res[0])
	if err != nil {
		return errors.Wrap(err, "failed to constructs a document JSON-map")
	}

	doc.SetRev(obj.GetRev())
	err = doc.SetID(obj.GetID())
	if err != nil {
		return errors.Wrap(err, "failed to set id")
	}

	doc.Id = int64(userID)
	err = cdb.Store(d.dbInstance, doc)
	if err != nil {
		return errors.Wrap(err, "Unable to write into couchdb")
	}
	return nil
}

func (d *customDocumentQ) DeleteDocument(userID int) error {
	fields := []string{"_id"}
	selector := fmt.Sprintf("id == %d", userID)
	res, err := d.dbInstance.Query(fields, selector, nil, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "Unable to read from couchdb")
	}

	obj := new(CustomDocument)
	err = cdb.FromJSONCompatibleMap(obj, res[0])
	if err != nil {
		return errors.Wrap(err, "failed to constructs a document JSON-map")
	}

	err = d.dbInstance.Delete(obj.GetID())
	if err != nil {
		return errors.Wrap(err, "Unable to delete from couchdb")
	}

	return nil
}
