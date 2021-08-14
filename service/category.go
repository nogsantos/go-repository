package service

import (
	"github.com/nogsantos/repository/entity"
)

type CategorieService struct {
	Repository entity.CategoryRepository
}

func NewCategoryService(repository entity.CategoryRepository) *CategorieService {
	return &CategorieService{Repository: repository}
}

func (c *CategorieService) FindById(id string) (entity.Category, error) {
	category, error := c.Repository.Get(id)
	if error != nil {
		return entity.Category{}, error
	}
	return category, nil
}

func (c *CategorieService) All() ([]entity.Category, error) {
	category, error := c.Repository.All()
	if error != nil {
		return []entity.Category{}, error
	}
	return category, nil
}

func (c *CategorieService) Create(name string) (entity.Category, error) {
	category := entity.NewCategory()
	category.Name = name
	cat, error := c.Repository.Create(*category)
	if error != nil {
		return entity.Category{}, error
	}
	return cat, nil
}