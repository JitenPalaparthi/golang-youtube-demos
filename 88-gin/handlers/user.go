package handlers

import (
	"demo/db"
	"demo/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	IUserDB db.IUserDB
}

type IUserhandler interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetAllBy(ctx *gin.Context)
	DeleteByID(ctx *gin.Context)
}

func NewUserHandler(iUserDB db.IUserDB) IUserhandler {
	return &UserHandler{IUserDB: iUserDB}
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

	// user.Status = "active"
	// user.LastModified = time.Now().Unix()
	// user.Id = uint(rand.IntN(10000))

	// err = utils.SaveToFile("user.dat", user)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
	// 	return
	// }

	user, err = u.IUserDB.Create(user)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
		return
	}

	ctx.JSON(201, user)
	//ctx.String(201, "user successfully saved in the file")
}

func (u *UserHandler) Update(ctx *gin.Context) {
	user := new(models.User)
	err := ctx.Bind(user) // read the request data, and bind it into the user object, that means unmarshal the data into user object

	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
		return
	}
	if user.Id == 0 {
		log.Println("invalid id")
		ctx.String(http.StatusBadRequest, "id is must to udpate")
		return
	}

	// err = user.Validate()
	// if err != nil {
	// 	log.Println(err.Error())
	// 	ctx.String(http.StatusBadRequest, err.Error())
	// 	return
	// }

	user, err = u.IUserDB.Update(user)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
		return
	}

	ctx.JSON(201, user)
	//ctx.String(201, "user successfully saved in the file")
}

func (u *UserHandler) GetByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		log.Println("id path param not foundf")
		ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
		return
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "invalid id.Id should be a digit")
		return
	}

	user, err := u.IUserDB.GetByID(id)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "seems something went wrong.Probably invalid input")
		return
	}

	ctx.JSON(200, user)

}

func (u *UserHandler) GetAll(ctx *gin.Context) {
	users, err := u.IUserDB.GetAll()
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "seems something went wrong.Probably invalid input")
		return
	}
	ctx.JSON(200, users)
}

func (u *UserHandler) GetAllBy(ctx *gin.Context) {

	limit, ok := ctx.Params.Get("limit")
	if !ok {
		log.Println("limit path param not foundf")
		ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
		return
	}

	_limit, err := strconv.Atoi(limit)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "invalid limit.limit should be a digit")
		return
	}

	offset, ok := ctx.Params.Get("offset")
	if !ok {
		log.Println("offset path param not foundf")
		ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
		return
	}

	_offset, err := strconv.Atoi(offset)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "invalid offset.offset should be a digit")
		return
	}

	users, err := u.IUserDB.GetAllBy(_limit, _offset)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "seems something went wrong.Probably invalid input")
		return
	}
	ctx.JSON(200, users)
}

func (u *UserHandler) DeleteByID(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		log.Println("id path param not foundf")
		ctx.String(http.StatusBadRequest, "Oops.Something went wrong with the request data")
		return
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "invalid id.Id should be a digit")
		return
	}

	err = u.IUserDB.DeleteByID(id)
	if err != nil {
		log.Println(err.Error())
		ctx.String(http.StatusBadRequest, "seems something went wrong.Probably invalid input or id to delete is not available")
		return
	}

	ctx.String(http.StatusNoContent, "record successfully deleted")

}
