package main

import (
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("godror", "IXUDBTEST/uBqffxYzciU9fFR3@10.82.71.188:1521/tcbssit")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	fmt.Println("Connected to Oracle")
}

