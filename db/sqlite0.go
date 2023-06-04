package db

// https://gorm.io/docs/

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-faker/faker/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func createFakeRecord(db *gorm.DB) {
	db.Create(&Product{Code: faker.Word(), Price: genRandomPrice()})
}

// pass a ptr via argument
func findFirstByPk(db *gorm.DB, p *Product, pk uint) {
	db.First(p, pk)
	fmt.Println(p)
}

func findFirstByProductCode(db *gorm.DB, p *Product, name string) {
	db.First(p, "code = ?", name)
	fmt.Println(p)
}

func findFirstByPrice(db *gorm.DB, p *Product, price uint) {
	db.First(p, "price = ?", 200)
	fmt.Println(p)
}

func updateProductPrice(db *gorm.DB, p *Product, price uint) {
	db.Model(p).Update("Price", price)
}

func RUN_sqlite0() {
	db, err := gorm.Open(sqlite.Open("sqlite3_test0.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&Product{})

	// create
	createFakeRecord(db)

	// Read
	var product1 Product
	findFirstByPk(db, &product1, 2)
	findFirstByProductCode(db, &product1, "D31")
	findFirstByPrice(db, &product1, 200)

	// Update - update product's price to 200
	updateProductPrice(db, &product1, 666)

	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product (soft delete)
	// db.Delete(&product, 1)

}

func genRandomPrice() uint {
	rand.Seed(time.Now().UnixNano())
	randPriceFrom10to1000 := rand.Intn(991) + 10
	return uint(randPriceFrom10to1000)
}
