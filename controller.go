package retail_shop

import (
	"errors"
)

var ErrInvalidIdString = errors.New("id string is not valid ")

func GetById(id string) (*Product, error) {
	if len(id) < 3 || len(id) > MaxIdLen {
		return nil, ErrInvalidIdString
	}
	var info = new(Product)
	err := DB.First(info, "id=?", id).Error

	return info, err
}

func GetPriceAbove(price int) []Product {
	var infos []Product
	DB.Where("price > ?", price).Find(&infos)
	return infos
}

func SaveProduct(p *Product) error {
		panic("not implemented")
		return  nil
}


