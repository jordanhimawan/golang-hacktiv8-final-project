package models

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" valid:"required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" valid:"required"`
	UserId         string
	User           *User
}
