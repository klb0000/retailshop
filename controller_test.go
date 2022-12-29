package retail_shop

import (
	"fmt"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var samples = []Product{

	{"c2d766ca982eca8304150849735ffef9", "Alisha Solid Women's Cycling Shorts", 300},
	{"7f7036a6d550aaa89d34c77bd39a5e48", "FabHomeDecor Fabric Double Sofa Bed", 32157},
	{"c275ee5ac19f774a3ef7da71b40aabd8", "Style Foot Bellies", 899},
}

var testDBFile = "data/testdata.db"
var testDB = getDB()

func getDB() *gorm.DB {
	os.Remove(testDBFile)
	var testDB, err = gorm.Open(sqlite.Open(testDBFile), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("unable to set up  testdatabase")
	}
	return testDB

}

func TestMain(m *testing.M) {

	testDB.AutoMigrate(&Product{})
	for i := range samples {
		err := SaveProduct(&samples[i], testDB)
		if err != nil {
			fmt.Println(err)
			panic("error seting up test db")
		}
	}
	
	os.Exit(m.Run())

}

func TestGetById(t *testing.T) {
	s := samples[0]
	//basic sample check
	row, err := GetById(s.ID, testDB)
	if err != nil {
		t.Error(err)
		return
	}
	if row.ID != s.ID {
		t.Error("invalid row returned")
	}
	fmt.Println(row, err)

	// should not find the query
	_, err = GetById(s.ID+"adssdakfl", testDB)
	if err == nil {
		t.Error(err)
		return
	}

}

func TestGetByName(t *testing.T) {
	s := samples[1]
	//basic sample check
	row, err := GetByName(s.Name, testDB)
	if err != nil {
		t.Error(err)
		return
	}
	if row.Name != s.Name {
		t.Error("invalid row returned")
	}
	fmt.Println(row.Name, err)

	// should not find the query
	_, err = GetByName(s.Name+"adssdakfl", testDB)
	if err == nil {
		t.Error(err)
		return
	}

}

func TestDoesNameExist(t *testing.T) {
	name := samples[0].Name
	fmt.Println(name)
	if !ProductNameExists(name, testDB) {
		t.Errorf("%s exists\n", name)
	}
	name = "sdkfjsdkfjkasfasd"
	if ProductNameExists(name, testDB) {
		t.Errorf("%s does not exist", name)
	}
}

func TestGetPriceAbove(*testing.T) {
	infos := GetPriceAbove(600, testDB)
	fmt.Println(infos)
	fmt.Println(infos[0])
}

func TestSaveProduct(t *testing.T) {
	invalidPd := Product{
		ID:    "124",
		Name:  "productInvalid",
		Price: 10,
	}
	err := SaveProduct(&invalidPd, testDB)
	if err == nil {
		t.Error("should not save invalid product")
	}

	validProductWhichExists := samples[1]
	err = SaveProduct(&validProductWhichExists, testDB)
	if err == nil {
		t.Error("should throw error if product alerady exists")
	}

	validProduct := Product{
		ID:    "1234",
		Name:  "product1",
		Price: 100,
	}
	err = SaveProduct(&validProduct, testDB)
	if err != nil {
		t.Error(err)
	}
	DeleteById(validProduct.ID, testDB)
}

func TestDeleteById(t *testing.T) {
	id := samples[2].ID
	err := DeleteById(id, testDB)
	if err == nil {
		fmt.Println("Deleted Properly")

	} else {
		t.Error(err)

	}

}

func TestUpdateProduct(t *testing.T) {
	s := samples[0]
	s2 := samples[0]
	s2.Name = "Name changed"
	err := UpdateProduct(&s, &s2, testDB)
	if err != nil {
		t.Error(err)
	}
	p, err := GetById(s.ID, testDB)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(p.Name)
}
