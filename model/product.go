package model

type Product struct {
	ID                 uint    `json:"id" gorm:"primary_key"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Price              float32 `json:"price"`
	DiscountPercentage float32 `json:"discountPercentage"`
	Rating             float32 `json:"rating"`
	Stock              int32   `json:"stock"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
	Thumbnail          string  `json:"thumbnail"`
}

type AddProductInput struct {
	ID                 uint    `json:"id" gorm:"primary_key"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Price              float32 `json:"price"`
	DiscountPercentage float32 `json:"discountPercentage"`
	Rating             float32 `json:"rating"`
	Stock              int32   `json:"stock"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
	Thumbnail          string  `json:"thumbnail"`
}

//kafka model -> 2 fields endpoint string,  data message byte data type

//structure pass to publisher
