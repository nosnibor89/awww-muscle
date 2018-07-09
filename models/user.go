package models

import (
	"github.com/jinzhu/gorm"
)

func init() {

}

type User struct {
	gorm.Model
	Name         string  `json:"name"`
	Email        string  `gorm:"type:varchar(100);unique_index" json:"email"`
	Role         string  `gorm:"size:255" json: "role"`
}

//func (u User) Find() User[]  {
//
//}