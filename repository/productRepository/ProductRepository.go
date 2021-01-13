package productRepository

import (
	"github.com/jinzhu/gorm"
	"new/src/connection"
)

type repository struct {}

type Product struct {
	gorm.Model
	CategoryID	uint	`json:"category_id"`
	Articul		string	`json:"articul"`
	Name		string	`json:"name"`
}

func newProductRepository() ProductRepositoryInterface {
	return &repository{}
}

//получение товара по id
//если товар не найден, вернет nil
func (r *repository) GetById(id uint) *Product {
	product := &Product{}
	err := connection.Connection.GetDB().Table("products").Where("id = ?", id).First(product).Error

	if err != nil {
		return nil
	}

	return product
}