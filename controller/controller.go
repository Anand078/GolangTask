package controller

import (
	"MS1/model"
	"MS1/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Service *service.Service
}

func (controller *Controller) CreateProduct(c *gin.Context) {
	var input model.AddProductInput
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
	fmt.Println("path: ", c.Request.URL.Path)
	return
}

func (controller *Controller) UpdateProduct(c *gin.Context) {
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
	fmt.Println("input : ", updateInput)
	fmt.Println("uid: ", uid)
	product, err := controller.Service.UpdateProduct(c, &updateInput, uint(uid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
	fmt.Println("output data", product)
	fmt.Println("host: ", c.Request.URL.Host)
	fmt.Println("path: ", c.Request.URL.Path)
	return
}
