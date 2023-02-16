package service

import (
	"MS1/config"
	"MS1/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Service struct {
	Publisher *config.Publisher
}

// service to insert a new product into productdb
func (s *Service) CreateProduct(c *gin.Context, input *model.AddProductInput) (product model.Product, err error) { //change input to addProduct
	product = model.Product{Title: input.Title, Description: input.Description, Price: input.Price, DiscountPercentage: input.DiscountPercentage,
		Rating: input.Rating, Stock: input.Stock, Brand: input.Brand, Category: input.Category, Thumbnail: input.Thumbnail}

	db := getDB(c)
	db.Create(&product)
	s.Publisher.PublishKafka(&product) //replace with new function
	return product, nil
}

// service to update existing product with new data
func (s *Service) UpdateProduct(c *gin.Context, updateInput *model.Product, id uint) (product model.Product, err error) { //change input to addProduct

	db := getDB(c)

	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		return product, err
	}

	db.Model(&product).Updates(updateInput)
	s.Publisher.PublishKafka(&product)
	return product, nil
}

func getDB(c *gin.Context) *gorm.DB {
	db := c.MustGet("db").(*gorm.DB)
	return db
}

//publishProductMsg
//s.publisher.publishkafka(&model)
