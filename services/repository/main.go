package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


//
type Finder interface {
	FindAll () (error, []interface{})
	FindOne(id int) (error, interface{})
}

type Repo struct {
	User userRepo
}

var db *gorm.DB
var err error

func init()  {
	db, err = gorm.Open("mysql", "root:awwmuscle@/awwmuscle?charset=utf8&parseTime=True&loc=Local")

	//DB.

	if err != nil{
		log.Fatal("Error connecting to DB\n ", err)
	}

	//defer DB.Close()
}
