package mysqlunit

import (
	"github.com/dogukanayd/handling-database-go-unit-test/internal/platform/database/testmysql"
	"testing"
)

func TestTruncate(t *testing.T) {
	t.Helper()

	t.Helper() // provides more meaningful outputs
	c := testmysql.StartContainer(t)

	connection, err := testmysql.ConnectionList.Get("test_database", c.Host)

	if err != nil {
		t.Fatalf("can not open database connection: %v", err)
	}

	informationSchemaConnection, err := testmysql.ConnectionList.Get("information_schema", c.Host)

	if err != nil {
		t.Fatalf("Opening database connection to information schema: %v", err)
	}

	t.Run("it should truncate tables", func(t *testing.T) {
		err := Truncate(informationSchemaConnection, connection)

		if err != nil {
			t.Fatalf("error accured when truncate the tables: %v", err)
		}
	})

}
