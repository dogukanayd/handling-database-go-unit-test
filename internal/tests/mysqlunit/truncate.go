package mysqlunit

import (
	"database/sql"
	"fmt"
)

type Table struct {
	Name string `db:"TABLE_NAME"`
}

type Tables []*Table

// added table name will never truncated by this package
var truncateExceptions = map[string]bool{
	"productTypes": true,
}

// Truncate the tables in tablesToBeTruncated tables
func Truncate(informationSchemaConnection, db *sql.DB) error {
	var t Tables
	_, err := db.Exec("SET FOREIGN_KEY_CHECKS=0")

	if err != nil {
		return err
	}

	rows, err := informationSchemaConnection.Query("SELECT TABLE_NAME FROM TABLES WHERE TABLE_SCHEMA = 'test_database'")

	if err != nil {
		return err
	}

	for rows.Next() {
		table := new(Table)

		_ = rows.Scan(&table.Name)

		t = append(t, table)
	}

	for _, table := range t {
		if truncateExceptions[table.Name] {
			continue
		}

		_, err := db.Exec(fmt.Sprintf("TRUNCATE %v", table.Name))

		if err != nil {
			return err
		}
	}

	return nil
}
