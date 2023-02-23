package controller

import (
	"MS1/model"
	"MS1/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Controller struct {
	Service *service.Service
}

func (controller *Controller) CreateProduct(c *gin.Context) {
	log.Infoln("controller.go, CreateProduct function started....")
	var input model.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := controller.Service.CreateProduct(c, &input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
	log.Infoln("path: ", c.Request.URL.Path)
	log.Infoln("controller.go, CreateProduct function ended....")
	return
}

func (controller *Controller) UpdateProduct(c *gin.Context) {
	log.Infoln("controller.go, UpdateProduct function started....")
	var updateInput model.Product
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	log.Infoln("input : ", updateInput)
	log.Infoln("uid: ", uid)
	product, err := controller.Service.UpdateProduct(c, &updateInput, uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
	log.Infoln("output data", product)
	log.Infoln("host: ", c.Request.URL.Host)
	log.Infoln("path: ", c.Request.URL.Path)
	log.Infoln("controller.go, UpdateProduct function ended....")
	return
}
