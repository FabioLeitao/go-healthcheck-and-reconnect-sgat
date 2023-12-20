package main

import (
	"database/sql"
	"fmt"

	"github.com/captain-corgi/golang-oracledb-example/pkg/oracle"
)

func corrige_cosetting() {
	// Express intensions
	//	fmt.Println("Resetting VALUE to correct cosetting behaviour")

	// Create connection
	conn := oracle.NewConnection()
	defer conn.Close()

	// Create statement for update
	sqlStatement := `
			UPDATE
				CORE.cosetting 
				SET VALUE = 'TRUE'
			WHERE 
				id = 681
				`
//	stmt, err := conn.Prepare(sqlStatement)

//	res, err := conn.Exec(sqlStatement)
	_, err := conn.Exec(sqlStatement)
	if err != nil {
	  panic(err)
	}
//	count, err := res.RowsAffected()
//	if err != nil {
//	  panic(err)
//	}
//	fmt.Println(count)
}

func main() {
	// Greeting
//	fmt.Println("Hello, Oracle DB!")

	// Create connection
	conn := oracle.NewConnection()
	defer conn.Close()

	// Create statement for select
	stmt, err := conn.Prepare(`SELECT
									ID,
									KEY,
									VALUE,
									UPDATED
								FROM 
									CORE.cosetting 
								WHERE 
									id = 681
				`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute statement
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}

	// Iterate over results
	for rows.Next() {
		// Define columns
		var ID,
			KEY,
			VALUE,
			UPDATED sql.NullString

		// Scan columns
		err = rows.Scan(&ID,
			&KEY,
			&VALUE,
			&UPDATED)
		if err != nil {
			panic(err)
		}

		// Use results for logic
		if VALUE.String == "false" {
//			fmt.Println("Deveria ser resetado para TRUE")
			corrige_cosetting()
		}
		if VALUE.String == "falso" {
//			fmt.Println("Deveria ser resetado para TRUE")
			corrige_cosetting()
		}


		// Print results
		fmt.Println(ID.String,
			KEY.String,
			VALUE.String,
			UPDATED.String)

	}

}
