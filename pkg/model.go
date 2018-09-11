package pkg

import "database/sql"

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	return []product{},nil
}
