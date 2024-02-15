package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	Id         uint   `gorm:"primaryKey" form:"id"`
	Message    string `gorm:"not null" form:"message" valid:"required~Message is required"`
	UserID     uint
	PhotoID    uint `gorm:"not null" form:"photo_id" json:"photo_id" valid:"required~Photo ID is required"`
	Created_at time.Time
	Updated_at time.Time
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
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

func (u *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Updated_at = time.Now()

	err = nil
	return
}
