package retailshop

import (
	"database/sql"
)

type Product struct {
	//gorm.Model
	ID    string `gorm:"primaryKey"`
	Name  string `gorm:"unique;not null;default:null"`
	Price int    `gorm:"default:null; not null"`
}

var ProductTableName = "products"

func CreateTableProduct(db *sql.DB) {
	db.Exec(
		`CREATE TABLE products2 (
		id TEXT PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		price NUMBER
		);`,
	)

}

func scanProduct(row *sql.Row, pd *Product) error {
	return row.Scan(&pd.ID, &pd.Name, &pd.Price)
}

func QProductByID(id string, db *sql.DB) (*Product, error) {
	if !isValidId(id) {
		return nil, ErrInvalidIdString
	}

	row := db.QueryRow("SELECT * FROM products WHERE id=?", id)
	var pd = new(Product)
	return pd, scanProduct(row, pd)
}

func InsertProduct(pd *Product, db *sql.DB) error {
	if !isValidProduct(pd) {
		return ErrInvalidProduct
	}

	q := `INSERT INTO products 
		(id, name, price) 
		VALUES (?, ?, ?);`

	_, err := db.Exec(q, pd.ID, pd.Name, pd.Price)
	return err

}

func UpdateProductPrice(id string, newPrice int, db *sql.DB) error {
	q := `
	UPDATE products
	SET price=?
	WHERE 
		id=?;`
	_, err := db.Exec(q, newPrice, id)
	return err
}

// func UpdateProduct(oldPd, newPd *Product, db *sql.DB) error {
// 	if err := DeleteProductByID(oldPd.ID, db); err != nil {
// 		return err
// 	}
// 	if err := InsertProduct(newPd, db); err != nil {
// 		return err
// 	}
// 	return nil
// }

func DeleteProductByID(id string, db *sql.DB) error {
	q := `DElETE FROM products
	WHERE id=?`
	_, err := db.Exec(q, id)
	return err
}

func QProductByName(name string, db *sql.DB) (*Product, error) {
	row := db.QueryRow("SELECT * FROM products WHERE name=?", name)
	var pd = new(Product)
	return pd, scanProduct(row, pd)
}
