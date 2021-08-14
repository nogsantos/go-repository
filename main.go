package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/nogsantos/repository/entity"
	"github.com/nogsantos/repository/repository"
	"github.com/nogsantos/repository/service"
)

func main() {
	memory()
	database()
	listAllCategoriesInDb()
}

func memory() {
	db := repository.CategoryMemoryDb{Categories: []entity.Category{}}
	repositoryMemory := repository.NewCategoryRepositoryMemory(db)

	service := service.NewCategoryService(repositoryMemory)
	category := "My category " + time.Now().Format(time.RubyDate)
	cat, error := service.Create(category)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Memory: %v", cat)
}

func database() {
	db, error := sql.Open("sqlite3", "./repository.sqlite")
	if error != nil {
		log.Fatal(error)
	}
	repositoryDatabase := repository.NewCategoryRepositorySqlite(db)

	service := service.NewCategoryService(repositoryDatabase)
	category := "My category " + time.Now().Format(time.RubyDate)
	cat, error := service.Create(category)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("\nDatabase: %v", cat)
}

func listAllCategoriesInDb() {
	db, error := sql.Open("sqlite3", "./repository.sqlite")
	if error != nil {
		log.Fatal(error)
	}
	repositoryDatabase := repository.NewCategoryRepositorySqlite(db)

	service := service.NewCategoryService(repositoryDatabase)
	services, error := service.All()
	if error != nil {
		log.Fatal(error)
	}

	for i, sv := range services {
		fmt.Printf("\nDatabase %v: %v",i, sv)
	}
}

