package actions

import (
	"github.com/nosnibor89/awww-muscle/models"
	"github.com/nosnibor89/awww-muscle/services/repository"
	"log"
	"github.com/labstack/echo"
	"net/http"
)

var user models.User
var repo repository.Repo

func GetUsers(c echo.Context) error  {
	err, users := repo.User.Find()
	if err == nil{
		log.Print("Error finding users")
	}

	return c.JSON(http.StatusOK, users)

}
