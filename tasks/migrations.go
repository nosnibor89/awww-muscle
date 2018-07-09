package tasks

import (
	"fmt"

	"github.com/markbates/grift/grift"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/gommon/log"
	"github.com/nosnibor89/awww-muscle/models"
)
func migrateDB(c *grift.Context) error {
	fmt.Println("Lauching rocket...")

	db, err := gorm.Open("mysql", "root:awwmuscle@/awwmuscle?charset=utf8&parseTime=True&loc=Local")

	if err != nil{
		log.Fatal("Error connecting to DB")
	}

	defer db.Close()

	// Create table for model `User`
	db.CreateTable(&models.User{})

	// will append "ENGINE=InnoDB" to the SQL statement when creating table `users`
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&models.User{})

	db.Create(&models.User{Name:"Robinson", Email:"nosnibor1989@gmail.com", Role:"Admin"})

	return nil
}


