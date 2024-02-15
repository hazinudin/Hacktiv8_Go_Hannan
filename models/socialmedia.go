package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Socialmedia struct {
	Id               uint   `gorm:"primaryKey"`
	Name             string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	Social_media_url string `gorm:"not null" json:"social_media_url" form:"url" valid:"url~Invalid social media URL"`
	UserID           uint
	Created_at       time.Time
	Updated_at       time.Time
}

func (u *Socialmedia) BeforeCreate(tx *gorm.DB) (err error) {
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

func (u *Socialmedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Updated_at = time.Now()

	err = nil
	return
}
