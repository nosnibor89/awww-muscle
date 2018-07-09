package repository

import "github.com/nosnibor89/awww-muscle/models"

type userRepo struct { }

func (userRepo) Find() (error, []models.User){
	var users []models.User

	db.Find(&users)
	//defer DB.Close()


	return nil, users
}