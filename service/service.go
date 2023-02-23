package service

import (
	"MS1/config"
	"MS1/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	Publisher *config.Publisher
}

// service to insert a new product into productdb
func (s *Service) CreateProduct(c *gin.Context, input *model.Product) (product model.Product, err error) {
	log.Infoln("service.go, CreateProduct function started.....")
	product = model.Product{Title: input.Title, Description: input.Description, Price: input.Price, DiscountPercentage: input.DiscountPercentage,
		Rating: input.Rating, Stock: input.Stock, Brand: input.Brand, Category: input.Category, Thumbnail: input.Thumbnail}

	db := getDB(c)
	db.Create(&product)
	s.Publisher.PublishKafka(&product) //replace with new function
	log.Infoln("service.go, CreateProduct function ended.....")
	return product, nil
}

// service to update existing product with new data
func (s *Service) UpdateProduct(c *gin.Context, updateInput *model.Product, id uint) (product model.Product, err error) { //change input to addProduct
	log.Infoln("service.go, UpdateProduct function started.....")
	db := getDB(c)

	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		return product, err
	}

	db.Model(&product).Updates(updateInput)
	s.Publisher.PublishKafka(&product)
	log.Infoln("service.go, UpdateProduct function ended.....")
	return product, nil
}

func getDB(c *gin.Context) *gorm.DB {
	log.Infoln("service.go, getDB function started.....")
	db := c.MustGet("db").(*gorm.DB)
	log.Infoln("service.go, getDB function ended.....")
	return db
}

//publishProductMsg
//s.publisher.publishkafka(&model)
