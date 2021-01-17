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

type FindByConditions struct {
	 Where map[string]interface{}
	 Limit uint
	 Offset uint
	 Sort string
	 Order string
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

//поиск товаров
func (r *repository) FindBy(cond FindByConditions) []Product {
	var products []Product

	err := connection.Connection.GetDB().Table("products").Where(cond.Where).Order(cond.Sort + " " + cond.Order).Limit(cond.Limit).Offset(cond.Offset).Find(&products).Error
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