package models

type Event struct {
	Kind string                 `json:"kind"`
	Info map[string]interface{} `json:"info"`
}

const (
	NATSTopic = "models.events"

	EventUserInfoAdd    = "user_info:add"
	EventUserInfoUpdate = "user_info:update"
	EventUserInfoDelete = "user_info:delete"

	EventBuzzAdd    = "buzz:add"
	EventBuzzUpdate = "buzz:update"
	EventBuzzDelete = "buzz:delete"
)
