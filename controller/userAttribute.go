package controller

import (
	"go-rest-user/model"
	"go-rest-user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var input service.UserAttribute

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	userAttributes := service.UserAttribute{
		Id_login:  input.Id_login,
		Email:     input.Email,
		Firtsname: input.Firtsname,
		Midlename: input.Midlename,
		Lastname:  input.Lastname,
		Job:       input.Job,
	}

	savedAttibute, err := userAttributes.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"msg": "Success", "data": savedAttibute})
}

func GetAll(c *gin.Context) {

	var userAll []service.UserAttribute

	if err := model.Database.Find(&userAll).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"msg": "success", "data": userAll})
	}
}

func GetById(c *gin.Context) {
	var input service.UserAttribute

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	userId, err := service.FindUserById(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": userId})

}

func GetByLoginid(c *gin.Context) {
	var input service.UserAttribute

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
	}

	user, err := service.FindUserByLoginId(input.Id_login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": user})
}

func UpdateUser(c *gin.Context) {
	var input service.UserAttribute

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
	}

	userIdLogin, err := service.FindUserById(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	input.Id_login = userIdLogin.Id_login
	c.ShouldBindJSON(&input)

	err = service.Update(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": input})

}

func Delete(c *gin.Context) {
	var input service.UserAttribute

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	userDelete, err := service.Delete(input.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success", "data": userDelete})
}
