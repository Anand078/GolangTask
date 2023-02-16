package main

import (
	"MS1/config"
	"MS1/controller"
	"MS1/model"
	"MS1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	db := config.InitDB()
	db.AutoMigrate(&model.Product{})
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	publisher := config.InitKafka()
	service := initService(publisher)
	var ctr = initController(service)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.POST("/product", ctr.CreateProduct)
	r.PATCH("/product/:id", ctr.UpdateProduct)
	r.Run()
}

func initController(service *service.Service) (ctr *controller.Controller) {
	return &controller.Controller{
		Service: service,
	}
}

func initService(publisher *config.Publisher) (svc *service.Service) {
	return &service.Service{
		Publisher: publisher,
	}
}
