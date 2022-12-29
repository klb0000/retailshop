package retail_shop

import (
	"fmt"
	"testing"
)

var samples = []Product{

	{"c2d766ca982eca8304150849735ffef9", "Alisha Solid Women's Cycling Shorts", 999},
	{"7f7036a6d550aaa89d34c77bd39a5e48", "FabHomeDecor Fabric Double Sofa Bed", 32157},
}

func TestGetById(t *testing.T) {
	fmt.Println(len(samples[0].ID))
	//basic sample check
	row, err := GetById(samples[0].ID)
	if err != nil {
		t.Error(err)
		return
	}
	if row.ID != samples[0].ID {
		t.Error("invalid row returned")
	}
	fmt.Println(row, err)

	// should not find the query
	row, err = GetById(samples[0].ID + "adssdakfl")
	if err == nil {
		t.Error(err)
		return
	}

}

func TestGetPriceAbove(*testing.T) {
	infos := GetPriceAbove(6000)
	fmt.Println(infos[0])
}
