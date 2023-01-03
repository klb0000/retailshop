package retailshop

import (
	"errors"

	"gorm.io/gorm"
)

var ErrInvalidIdString = errors.New("id string is not valid ")
var ErrInvalidProduct = errors.New("product detail not valid")
var ErrProductNameExists = errors.New("name of product already exists in database")
var ErrUnableToUpdateProduct = errors.New("unable to update product")

func isValidId(id string) bool {
	return len(id) > 3 && len(id) <= MaxIdLen
}

func GetProductById(id string, db *gorm.DB) (*Product, error) {
	if !isValidId(id) {
		return nil, ErrInvalidIdString
	}
	var product = new(Product)
	err := db.First(product, "id=?", id).Error

	return product, err
}

func GetProductByName(name string, db *gorm.DB) (*Product, error) {
	if len(name) < MinProductNameLen {
		return nil, ErrInvalidProduct
	}
	var pd = new(Product)
	err := db.First(pd, "name=?", name).Error
	return pd, err
}

func GetAllProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	result := db.Find(&products)
	return products, result.Error
}

func GetProductAbovePrice(price int, db *gorm.DB) []Product {
	var infos []Product
	db.Where("price > ?", price).Find(&infos)
	return infos
}

func isValidProduct(p *Product) bool {
	return isValidId(p.ID) &&
		len(p.Name) >= MinProductNameLen
}

func ProductNameExists(name string, db *gorm.DB) bool {
	pd, err := GetProductByName(name, db)
	if err == nil {
		return pd.Name == name
	}
	return false
}

func ProductIdExist(id string, db *gorm.DB) bool {
	pd, err := GetProductById(id, db)
	if err == nil {
		return pd.ID == id
	}
	return false
}

func SaveProduct(p *Product, db *gorm.DB) error {
	if !isValidProduct(p) {
		return ErrInvalidProduct
	}
	if ProductNameExists(p.Name, db) {
		return ErrProductNameExists
	}
	// if id already exist than function which creates product has bug
	if ProductIdExist(p.ID, db) {
		panic("product id alerady exists")
	}

	return db.Create(p).Error
}

func DeleteProductById(id string, db *gorm.DB) error {
	pd := Product{ID: id}
	return db.Delete(&pd).Error
}

// need to rewrite this
// Just was lazy to read gorm documentation
func UpdateProduct(oldP, newP *Product, db *gorm.DB) error {

	//delete old product
	err := DeleteProductById(oldP.ID, db)
	if err != nil {
		return err
	}
	// save new product
	err = SaveProduct(newP, db)

	// if unable to save product
	if err != nil {
		saveErr := SaveProduct(newP, db)

		//if unable to save previously deleted prodcut
		if saveErr != nil {
			panic("error updating prodcut")
		}
		return ErrUnableToUpdateProduct
	}

	return nil

}
