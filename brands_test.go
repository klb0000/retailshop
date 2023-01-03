package retailshop

import "testing"

func TestTmp(t *testing.T) {
	db := initDB()
	defer db.Close()

	if err := CreateBrandLookupTable(db); err != nil {
		t.Error(err)
	}

}
