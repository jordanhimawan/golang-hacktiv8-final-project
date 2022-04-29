package models

type Comment struct {
	GormModel
	Message string `json:"omitempty"`
	UserId  string
	PhotoId string
	User    *User
	Photo   *Photo
}
