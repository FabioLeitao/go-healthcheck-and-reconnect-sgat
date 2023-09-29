package oracle

import (
	"database/sql"
)

func NewConnection() *sql.DB {
	conn, err := sql.Open("oracle", "oracle://tosp:Athenas2018@10.129.48.68/tosprd")
	if err != nil {
		panic(err)
	}

	return conn
}
