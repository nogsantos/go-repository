package repository

import (
	"errors"

	"github.com/nogsantos/repository/entity"
)

type CategoryMemoryDb struct {
	Categories []entity.Category
}

type CategoryRepositoryMemory struct {
	db CategoryMemoryDb
}

func NewCategoryRepositoryMemory(db CategoryMemoryDb) *CategoryRepositoryMemory {
	return &CategoryRepositoryMemory{db: db}
}

func (c *CategoryRepositoryMemory) Get(id string) (entity.Category, error) {
	for _, category := range c.db.Categories{
		if category.ID == id {
			return category, nil
		}
	}
	return entity.Category{}, errors.New("no category found with this id")
}

func (c *CategoryRepositoryMemory) All() ([]entity.Category, error) {
	return []entity.Category{}, errors.New("no list all")
}

func (c *CategoryRepositoryMemory) Create(category entity.Category) (entity.Category, error) {
	c.db.Categories = append(c.db.Categories, category)
	return category, nil
}