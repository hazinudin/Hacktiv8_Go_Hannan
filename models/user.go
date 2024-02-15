package models

import (
	"go_final_project/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	Id           uint   `gorm:"primaryKey"`
	Username     string `gorm:"not null" json:"username" form:"username" form:"username" valid:"required~Username is required"`
	Email        string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required, email~Invalid email format"`
	Password     string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required, minstringlength(6)~Password minimal character length is 6 characters"`
	Age          uint   `gorm:"not null" json:"age" form:"age" valid:"required~Age is required, range(8|200)"`
	Created_at   time.Time
	Updated_at   time.Time
	Comments     []Comment
	Photos       []Photo
	Socialmedias []Socialmedia
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	u.Created_at = time.Now()
	u.Updated_at = time.Now()

	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Updated_at = time.Now()
	err = nil
	return
}
