package models

type UserInfo struct {
	UserID     int64  `json:"userId"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
}
