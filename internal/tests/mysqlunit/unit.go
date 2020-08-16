package mysqlunit

import (
	"database/sql"
	"github.com/dogukanayd/handling-database-go-unit-test/internal/platform/database/testmysql"
	"testing"
	"time"
)

func NewUnit(t *testing.T) (*sql.DB, func()) {
	t.Helper() // provides more meaningful outputs
	c := testmysql.StartContainer(t)

	connection, err := testmysql.ConnectionList.Get("test_database", c.Host)

	if err != nil {
		t.Fatalf("can not open database connection: %v", err)
	}

	healthCheck(connection, t, c)

	informationSchemaConnection, err := testmysql.ConnectionList.Get("information_schema", c.Host)

	if err != nil {
		t.Fatalf("Opening database connection to information schema: %v", err)
	}

	tearDown := func() {
		t.Helper()
		_ = Truncate(informationSchemaConnection, connection)
		_ = connection.Close()
	}

	return connection, tearDown
}

// Wait for the database to be ready. Wait 100ms longer between each attempt.
// Do not try more than 20 times.
func healthCheck(connection *sql.DB, t *testing.T, c *testmysql.Container) {
	var pingError error

	maxAttempts := 20

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		pingError = connection.Ping()

		if pingError == nil {
			break
		}

		time.Sleep(time.Duration(attempts) * 100 * time.Millisecond)
	}

	if pingError != nil {
		testmysql.DumpContainerLogs(t, c)
		t.Fatalf("waiting for database to be ready: %v", pingError)
	}
}
