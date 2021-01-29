package productRepository

import (
	"github.com/jinzhu/gorm"
	"learning/connection"
)

type repository struct {
	db *gorm.DB
}

type Product struct {
	gorm.Model
	CategoryID	uint	`json:"category_id"`
	SKU			string	`gorm:"type:varchar(64);unique;not null"`
	Name		string	`gorm:"type:varchar(128);not null"`
}

type FindByConditions struct {
	 Where map[string]interface{}
	 Limit uint
	 Offset uint
	 Sort string
	 Order string
}

func newProductRepository() ProductRepositoryInterface {
	return &repository{
		db : connection.Connection.GetDB()}
}

//получение товара по id
//если товар не найден, вернет nil
func (r *repository) GetById(id uint) *Product {
	product := &Product{}
	err := r.db.Table("products").First(product, id).Error

	if err != nil {
		return nil
	}

	return product
}

//поиск товаров
func (r *repository) FindBy(cond FindByConditions) []Product {
	var products []Product

	err := r.db.Table("products").Where(cond.Where).Order(cond.Sort + " " + cond.Order).Limit(cond.Limit).Offset(cond.Offset).Find(&products).Error
	if err != nil {
		return nil
	}

	return products
}

//сохранить товар
func (r *repository) Save(product *Product) (*Product, error) {
	error := r.db.Save(product).Error

	return product, error
}

//удалить товар
func (r *repository) Delete(product *Product) {
	r.db.Unscoped().Delete(product)
}