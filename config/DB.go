package config

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func init()  {
	//db, err := gorm.Open("mysql", "test.db")
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//defer db.Close()
	//
	//// Migrate the schema
	//db.AutoMigrate(&Product{})
}