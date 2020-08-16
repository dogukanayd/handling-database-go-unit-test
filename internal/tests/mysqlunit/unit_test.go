package mysqlunit

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

type TestTable struct {
	Name string `db:"name"`
}

func TestNewUnit(t *testing.T) {
	connection, tearDown := NewUnit(t)
	defer tearDown()

	name := fmt.Sprintf("name-%v", time.Now().Unix())

	t.Run("it should create record on test database", func(t *testing.T) {
		var u TestTable

		tti, err := connection.Prepare("INSERT INTO test_table VALUES(?, ?)")

		if err != nil {
			t.Fatalf("can not prepare to test db: %v", err)
		}

		_, err = tti.Exec(1, name)

		if err != nil {
			t.Fatalf("can not insert to test db: %v", err)
		}

		row := connection.QueryRow("SELECT name FROM test_table WHERE name = ?", name)

		if err := row.Scan(&u.Name); err != nil {
			t.Fatalf("can not scan test_table name value")
		}

		if u.Name != name {
			t.Errorf("Name expected: %v, got: %v", name, u.Name)
		}
	})

}
