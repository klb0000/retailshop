package retailshop

import "database/sql"

type BrandLookup struct {
	PID       string
	BrandName string
}

func CreateBrandLookupTable(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE brandlookup (
		pid TEXT,
		brand_name TEXT,
		FOREIGN KEY (pid)
			REFERENCES products (id)
		);`,
	)
	return err
}
