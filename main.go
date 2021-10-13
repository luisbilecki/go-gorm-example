package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
}

type User struct {
	gorm.Model
	Name        string
	Company     Company
	CompanyID   int
	CreditCards []CreditCard
}

type Company struct {
	gorm.Model
	Name string
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	// Open in-memory connection
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})

	// Create the database schema
	db.AutoMigrate(&Product{}, &User{}, &Company{}, &CreditCard{})

	// Batch insertion of products
	productsData := []Product{
		{Name: "Eggs", Description: "Eggs", Price: 10},
		{Name: "Aspargus", Description: "Aspargus", Price: 18.95},
		{Name: "Strawberries", Description: "Fresh Strawberries", Price: 10},
	}
	db.Create(productsData)

	// Query examples
	var product Product
	db.First(&product, 1)
	fmt.Println(product)

	var products []Product
	db.Find(&products)
	db.Where("price > ?", 15).Find(&products)

	// Update
	db.Model(&product).Update("Price", 12)
	fmt.Println(product)

	// Delete
	db.Delete(&Product{}, 2)

	// Associations
	user := User{
		Name:        "John",
		Company:     Company{Name: "Foo"},
		CreditCards: []CreditCard{{Number: "12345"}},
	}
	db.Create(&user)

	db.First(&user)
	fmt.Println(user)
}
