package retailshop

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// var _db = initDB()

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./data/data.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func TestGetProduct(t *testing.T) {
	db := initDB()

	defer db.Close()
	id := "1001"
	pd, err := QProductByID(id, db)
	// pd, err := QProductByID(id, db)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(pd)

}

func TestInsertAndDeleteProduct(t *testing.T) {
	db := initDB()
	defer db.Close()

	pd := Product{"9988", "product 9988", 9988}
	if err := InsertProduct(&pd, db); err != nil {
		t.Error(err)
	}
	if err := DeleteProductByID(pd.ID, db); err != nil {
		t.Error(err)
	}
}

func TestQAllProducts(t *testing.T) {
	db := initDB()
	defer db.Close()

	products, err := QAllProduct(db)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(products[0])

}

func TestQProductByName(t *testing.T) {
	db := initDB()
	defer db.Close()

	name := "product 2"
	pd, err := QProductByName(name, db)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(pd)
}

func TestUpdatePrice(t *testing.T) {
	db := initDB()
	defer db.Close()

	// get oldprice of product
	id := "0001"
	pd, err := QProductByID(id, db)
	if err != nil {
		t.Error(err)
	}
	oldPrice := pd.Price

	//update price
	newPrice := 888
	if err := UpdateProductPrice(id, newPrice, db); err != nil {
		t.Error(err)
		return
	}

	// check if price has changed
	pd, err = QProductByID(id, db)
	if err != nil {
		t.Error(err)
	}
	if pd.Price != newPrice {
		t.Error("price not changed")
	}

	// update price to old price
	if err := UpdateProductPrice(id, oldPrice, db); err != nil {
		t.Error("err")
	}
}

// func TestUpdateProduct(t *testing.T) {
// 	db := initDB()
// 	defer db.Close()
// 	id := "0001"
// 	pd, err := QProductByID(id, db)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	pd2 := pd

// 	// change name of product
// 	pd2.Name = "product1-c"
// 	if err := UpdateProduct(pd, pd2, db); err != nil {
// 		t.Error(err)
// 	}

// }

func TestMakeTable(t *testing.T) {
	db := initDB()
	defer db.Close()
	CreateTableProduct(db)
}
