package handlers

import (
	"demo/models"
	"demo/utils"
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

type IUserhandler interface {
	Create(ctx *gin.Context)
}

func NewUserHandler() IUserhandler {
	return &UserHandler{}
}

func (u *UserHandler) Create(ctx *gin.Context) {
	user := new(models.User)
	err := ctx.Bind(user) // read the request data, and bind it into the user object, that means unmarshal the data into user object

	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
		return
	}

	err = user.Validate()
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	user.Status = "active"
	user.LastModified = time.Now().Unix()
	user.Id = uint(rand.IntN(10000))

	err = utils.SaveToFile("user.dat", user)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
		return
	}

	ctx.String(201, "user successfully saved in the file")

}
