package models

import (
	"fmt"

	"github.com/lancer-kit/armory/db"
	cdb "github.com/leesper/couchdb-golang"
	"github.com/pkg/errors"

	"lancer-kit/service-scaffold/config"
)

type CustomDocument struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	cdb.Document
}

type CustomDocumentQ struct {
	dbInstance *cdb.Database
}

func CreateCustomDocumentQ(cfg *config.Cfg) (*CustomDocumentQ, error) {
	newDocInstance := new(CustomDocumentQ)

	dbInstance, err := cdb.NewDatabase(cfg.CouchDB)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to couchdb")
	}

	newDocInstance.dbInstance = dbInstance
	return newDocInstance, nil
}

func (d *CustomDocumentQ) AddDocument(doc *CustomDocument) error {
	err := cdb.Store(d.dbInstance, doc)
	if err != nil {
		return errors.Wrap(err, "Unable to write into couchdb")
	}

	return nil
}

func (d *CustomDocumentQ) GetAllDocument(pQ db.PageQuery) ([]CustomDocument, error) {
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

func (d *CustomDocumentQ) GetDocument(userID int) ([]CustomDocument, error) {
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

func (d *CustomDocumentQ) UpdateDocument(userID int, doc *CustomDocument) error {
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

	doc.ID = int64(userID)
	err = cdb.Store(d.dbInstance, doc)
	if err != nil {
		return errors.Wrap(err, "Unable to write into couchdb")
	}
	return nil
}

func (d *CustomDocumentQ) DeleteDocument(userID int) error {
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
