package productRepository

import (
	"github.com/jinzhu/gorm"
	"new/src/connection"
)

type repository struct {}

type Product struct {
	gorm.Model
	CategoryID	uint	`json:"category_id"`
	SKU			string	`gorm:"type:varchar(64);unique;not null"`
	Name		string	`json:"name"`
}

func newProductRepository() ProductRepositoryInterface {
	return &repository{}
}

//получение товара по id
//если товар не найден, вернет nil
func (r *repository) GetById(id uint) *Product {
	product := &Product{}
	err := connection.Connection.GetDB().Table("products").First(product, id).Error

	if err != nil {
		return nil
	}

	return product
}

//поиск товаров по переданным параметрам, где:
//ключ массива = стобец
//значение = значение для выборки
func (r *repository) FindBy(criteria map[string]interface{}) []Product {
	var products []Product

	err := connection.Connection.GetDB().Table("products").Where(criteria).Find(&products).Error
	if err != nil {
		return nil
	}

	return products
}

//сохранить товар
func (r *repository) Save(product *Product) (*Product, error) {
	error := connection.Connection.GetDB().Save(product).Error

	return product, error
}