package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"tma.com.vn/lngochuy/rest/resources"
	"tma.com.vn/lngochuy/rest/services"
)

func GetUserList(ctx *gin.Context) {
	userList, err := services.GetUserList()
	if err != nil {
		response := NewAPIError(1, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := NewAPIResponse(userList)
	ctx.JSON(http.StatusOK, response)
}

func GetUser(ctx *gin.Context) {
	sid := ctx.Param("id")
	id64, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	id := uint(id64)

	user, err := services.GetUser(id)
	if err != nil {
		response := NewAPIError(1, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
	}

	response := NewAPIResponse(user)
	ctx.JSON(http.StatusOK, response)
}

func CreateUser(ctx *gin.Context) {
	request := new(resources.CreateUserRequest)
	err := ctx.BindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	id, err := services.CreateUser(request.Email, time.Time(request.DateOfBirth))
	if err != nil {
		response := NewAPIError(1, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := NewAPIResponse(id)
	ctx.JSON(http.StatusCreated, response)
}

func UpdateUser(ctx *gin.Context) {
	sid := ctx.Param("id")
	id64, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	id := uint(id64)

	request := new(resources.CreateUserRequest)
	err = ctx.BindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	err = services.UpdateUser(id, request.Email, time.Time(request.DateOfBirth))
	if err != nil {
		response := NewAPIError(1, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := NewAPIResponse(nil)
	ctx.JSON(http.StatusCreated, response)
}

func RemoveUser(ctx *gin.Context) {
	sid := ctx.Param("id")
	id64, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	id := uint(id64)

	err = services.RemoveUser(id)
	if err != nil {
		response := NewAPIError(1, err.Error())
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := NewAPIResponse(nil)
	ctx.JSON(http.StatusNoContent, response)
}
