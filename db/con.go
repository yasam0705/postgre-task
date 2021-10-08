package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const connStr = "user=postgres dbname=psql_task sslmode=disable"

var Db, _ = sql.Open("postgres", connStr)
