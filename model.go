package retail_shop

type Product struct {
	//gorm.Model
	ID    string `gorm:"primaryKey"`
	Name  string `gorm:"unique;not null;default:null"`
	Price int    `gorm:"default:null; not null"`
}
