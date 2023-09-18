package route

import (
	"kp-golang-restfull-api/config"
	"kp-golang-restfull-api/controller"

	"github.com/gin-gonic/gin"
)

func User() {
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}
	route := gin.Default()

	route.GET("/", controller.Index)
	route.GET("/user/:id", inDB.GetUser)
	route.GET("/users", inDB.GetUsers)
	route.POST("/user", inDB.CreateUser)
	route.PUT("/user", inDB.UpdateUser)
	route.DELETE("/user/:id", inDB.DeleteUser)
	route.Run()
}
