package product

import (
	"database/sql"
)

// Product ...
type Product struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

// CheckActivationStatusByName ...
func CheckActivationStatusByName(db *sql.DB, name string) bool {
	var p Product

	row := db.QueryRow("SELECT * FROM products WHERE name = ?", name)

	_ = row.Scan(&p.ID, &p.Name, &p.Status)

	return p.Status
}
