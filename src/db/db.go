package db

import (
	"database/sql"
	"fmt"

	_ "github.com/golang/glog"
	_ "github.com/lib/pq"
)

type DB_manager struct {
	*sql.DB
}

func NewDB(str string) (*DB_manager, error) {
	cockroach, err := sql.Open("postgres", str)
	if err != nil {
		fmt.Println("failed to  connnect with database. Error ", err)
		return nil, err
	}
	// Set the maximum number of concurrently open connections (in-use + idle)
	// to 5. Setting this to less than or equal to 0 will mean there is no
	// maximum limit (which is also the default setting)
	cockroach.SetMaxOpenConns(0)
	dbObj := &DB_manager{
		DB: cockroach,
	}
	return dbObj, nil
}
