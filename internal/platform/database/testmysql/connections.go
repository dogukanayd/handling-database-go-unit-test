package testmysql

import (
	"database/sql"
	"fmt"
	"sync"
)

// Connections is used for thread safe DB connection management
type Connections struct {
	list map[string]*sql.DB
	sync.Mutex
}

// ConnectionList is an initialization of connection list wrapper
var ConnectionList = &Connections{
	list: make(map[string]*sql.DB),
}

func (c *Connections) Get(database, host string) (*sql.DB, error) {
	if c.list == nil {
		c.list = make(map[string]*sql.DB)
	}

	db.host = fmt.Sprintf("tcp(%v)", host)

	dsn := fmt.Sprintf(
		"%v:%v@%v/%v?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true",
		db.user,
		db.password,
		db.host,
		database,
	)

	// lock thread
	c.Lock()
	defer c.Unlock()

	connection, ok := c.list[database]

	if !ok || connection.Ping() != nil {
		if db, err := sql.Open(db.driver, dsn); err != nil {
			return nil, err
		} else {
			c.list[database] = db

			return db, nil
		}
	}

	return connection, nil
}
