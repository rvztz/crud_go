package main

import (
	"github.com/gin-gonic/gin"

	"crud_go/controller"
	"crud_go/model"
)

func main() {
	r := gin.Default()
	model.SetDatabase()
	r.GET("/locations", controller.GetLocations)
	r.POST("/locations", controller.CreateLocation)
	r.GET("/locations/:id", controller.GetLocation)
	r.GET("/locations/:id/denue", controller.GetDenue)
	r.PATCH("/locations/:id", controller.UpdateLocation)
	r.DELETE("/locations/:id", controller.DeleteLocation)
	err := r.Run(":5050")
	if err != nil {
		return
	}
}
