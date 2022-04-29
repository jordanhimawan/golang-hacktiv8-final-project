package models

import (
	"sesi-final-project/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null,uniqueIndex" json:"username" form:"username" valid:"required~Your username is required"`
	Email       string        `gorm:"not null,uniqueIndex" json:"email" form:"email" valid:"required~Your email is required"`
	Password    string        `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Your password has to have a minimum length of 6 characters"`
	Age         uint8         `gorm:"not null" json:"age" form:"age" valid:"gte=8~Your age must not be less than 8 years old"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_media"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	return
}

func (p *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	return
}
