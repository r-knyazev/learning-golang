package categoryRepository

import (
	"github.com/jinzhu/gorm"
	"learning/connection"
)

type repository struct {
	db *gorm.DB
}

type Category struct {
	gorm.Model
	Name		string	`gorm:"type:varchar(128)"`
	Tags		string	`gorm:"type:varchar(255)"`
}

type FindByConditions struct {
	 Where map[string]interface{}
	 Limit uint
	 Offset uint
	 Sort string
	 Order string
}

func newCategoryRepository() CategoryRepositoryInterface {
	return &repository{
		db : connection.Connection.GetDB()}
}

//получение категории по id
//если категория не найдена, вернет nil
func (r *repository) GetById(id uint) *Category {
	category := &Category{}
	err := r.db.Table("categories").First(category, id).Error

	if err != nil {
		return nil
	}

	return category
}

//поиск категорий
func (r *repository) FindBy(cond FindByConditions) []Category {
	var category []Category

	err := r.db.Table("categories").Order(cond.Sort + " " + cond.Order).Limit(cond.Limit).Offset(cond.Offset).Find(&category).Error
	if err != nil {
		return nil
	}

	return category
}

//сохранить категорию
func (r *repository) Save(category *Category) (*Category, error) {
	error := r.db.Save(category).Error

	return category, error
}

//удалить категорию
func (r *repository) Delete(category *Category) {
	r.db.Delete(category)
}