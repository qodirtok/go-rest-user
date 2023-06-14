package config

import (
	"go-rest-user/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Http(router *gin.Engine) {

	router.GET("/getUser", controller.GetAll)
	router.POST("/getUserById", controller.GetById)
	router.POST("/getUserByLoginId", controller.GetByLoginid)
	router.POST("/userAttribute", controller.CreateUser)
	router.PUT("/userUpdate", controller.UpdateUser)
	router.DELETE("/userDelete", controller.Delete)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "OK"})
	})
}
