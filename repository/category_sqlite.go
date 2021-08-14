package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nogsantos/repository/entity"
)

type CategoryRepositorySqlite struct {
	db *sql.DB
}

func NewCategoryRepositorySqlite(db *sql.DB) *CategoryRepositorySqlite {
	return &CategoryRepositorySqlite{db: db}
}

func (c *CategoryRepositorySqlite) Get(id string) (entity.Category, error) {
	var category entity.Category
	stmt, err := c.db.Prepare(`select id, name from categories where id=?`)
	if err != nil {
		return entity.Category{}, nil
	}
	err = stmt.QueryRow(id).Scan(&category.ID, &category.Name)
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (c *CategoryRepositorySqlite) All() ([]entity.Category, error) {
	var category []entity.Category
	row, err := c.db.Query(`select id, name from categories`)
	if err != nil {
		return []entity.Category{}, nil
	}
	defer row.Close()

	for row.Next() {
		var id string
		var name string
		row.Scan(&id, &name)
		category = append(category, entity.Category{ID: id, Name: name})
	}

	return category, nil
}

func (c *CategoryRepositorySqlite) Create(category entity.Category) (entity.Category, error) {
	stmt, err := c.db.Prepare(`insert into categories (id, name) values (?, ?)`)
	if err != nil {
		return entity.Category{}, nil
	}

	_, err = stmt.Exec(category.ID, category.Name)
	if err != nil {
		return entity.Category{}, nil
	}

	err = stmt.Close()
	if err != nil {
		return entity.Category{}, nil
	}

	return category, nil
}