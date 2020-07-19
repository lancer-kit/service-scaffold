package repo

import (
	"fmt"

	"lancer-kit/service-scaffold/models"

	"github.com/lancer-kit/armory/db"
	cdb "github.com/leesper/couchdb-golang"
	"github.com/pkg/errors"
)

type CouchRepo struct {
	dbURL string
}

func NewCouchRepo(dbURL string) *CouchRepo {
	return &CouchRepo{dbURL: dbURL}
}

func (repo *CouchRepo) UserInfo() (UserInfoRepoI, error) {
	newDocInstance := new(UserInfoRepo)

	dbInstance, err := cdb.NewDatabase(repo.dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to connect to couchdb")
	}

	newDocInstance.dbInstance = dbInstance
	return newDocInstance, nil
}

type UserInfoRepoI interface {
	AddUserInfo(doc *models.UserInfo) error
	AllUserInfo(pQ db.PageQuery) ([]models.UserInfo, error)
	GetUserInfo(userID int64) ([]models.UserInfo, error)
	UpdateUserInfo(userID int64, doc *models.UserInfo) error
	DeleteUserInfo(userID int64) error
}

type UserInfo struct {
	models.UserInfo
	cdb.Document
}

type UserInfoRepo struct {
	dbInstance *cdb.Database
}

func (d *UserInfoRepo) AddUserInfo(doc *models.UserInfo) error {
	err := cdb.Store(d.dbInstance, doc)
	if err != nil {
		return errors.Wrap(err, "Unable to write into couchdb")
	}

	return nil
}

func (d *UserInfoRepo) AllUserInfo(pQ db.PageQuery) ([]models.UserInfo, error) {
	fields := []string{"id", "firstName", "secondName"}

	res, err := d.dbInstance.Query(fields, `exists(id,true)`, nil, int(pQ.PageSize), int(pQ.PageSize*(pQ.Page-1)), nil)

	if err != nil {
		return nil, errors.Wrap(err, "Unable to write into couchdb")
	}

	resSlice := make([]models.UserInfo, 0)
	for _, v := range res {
		obj := new(UserInfo)
		err = cdb.FromJSONCompatibleMap(obj, v)
		if err != nil {
			return nil, errors.Wrap(err, "failed to constructs a document JSON-map")
		}
		resSlice = append(resSlice, obj.UserInfo)
	}

	return resSlice, nil
}

func (d *UserInfoRepo) GetUserInfo(userID int64) ([]models.UserInfo, error) {
	fields := []string{"id", "firstName", "secondName"}

	res, err := d.dbInstance.Query(fields, fmt.Sprintf("id == %d", userID), nil, nil, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to write into couchdb")
	}

	resSlice := make([]models.UserInfo, 0)
	for _, v := range res {
		obj := new(UserInfo)
		err = cdb.FromJSONCompatibleMap(obj, v)
		if err != nil {
			return nil, errors.Wrap(err, "failed to constructs a document JSON-map")
		}
		resSlice = append(resSlice, obj.UserInfo)
	}

	return resSlice, nil

}

func (d *UserInfoRepo) UpdateUserInfo(userID int64, info *models.UserInfo) error {
	fields := []string{"_rev", "_id"}
	selector := fmt.Sprintf("id == %d", userID)
	res, err := d.dbInstance.Query(fields, selector, nil, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "Unable to read from couchdb")
	}

	obj := new(UserInfo)
	err = cdb.FromJSONCompatibleMap(obj, res[0])
	if err != nil {
		return errors.Wrap(err, "failed to constructs a document JSON-map")
	}
	doc := UserInfo{UserInfo: *info}
	doc.SetRev(obj.GetRev())
	err = doc.SetID(obj.GetID())
	if err != nil {
		return errors.Wrap(err, "failed to set id")
	}

	info.UserID = int64(userID)
	err = cdb.Store(d.dbInstance, info)
	if err != nil {
		return errors.Wrap(err, "unable to write into couchdb")
	}

	return nil
}

func (d *UserInfoRepo) DeleteUserInfo(userID int64) error {
	fields := []string{"_id"}
	selector := fmt.Sprintf("id == %d", userID)
	res, err := d.dbInstance.Query(fields, selector, nil, nil, nil, nil)
	if err != nil {
		return errors.Wrap(err, "Unable to read from couchdb")
	}

	obj := new(UserInfo)
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
