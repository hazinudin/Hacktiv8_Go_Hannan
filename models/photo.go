package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	Id         uint   `gorm:"primaryKey"`
	Title      string `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption    string `gorm:"not nul" json:"caption" form:"caption"`
	Photo_url  string `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo URL is required, url~Invalid URL"`
	Created_at time.Time
	Updated_at time.Time
	UserID     uint
	Comments   []Comment
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Updated_at = time.Now()
	u.Created_at = time.Now()

	err = nil
	return
}

func (u *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Updated_at = time.Now()

	err = nil
	return
}
