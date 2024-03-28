package models

import (
	"challenge-api/pkg/helper"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;varchar(200)" json:"username" form:"username" valid:"required~Your username is required"`
	Email    string `gorm:"not null;uniqueIndex;varchar(200)" json:"email" form:"email" valid:"required~Your email is required,email-Invalid email format"`
	Password string `gorm:"not null;varchar(200)" json:"password" fprm:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Orders []Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"orders"`
}

func (u *User) BeforeCreate (tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashPass(u.Password)

	err = nil
	return
}