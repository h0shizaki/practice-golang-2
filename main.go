package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Hello GO")
	fmt.Println("Driver : ", sql.Drivers())

	db, err := sql.Open("mysql", "root:@/test")

	if err != nil {
		fmt.Println("Can not connect to Database ")
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM user")
	if err != nil {
		fmt.Println("Error while fectching data")
	}

	//POST
	// stmt, err := db.Prepare("INSERT INTO user (f_name,l_name) VALUES(?,?)")
	// stmt.Exec("Yuuki", "Tachibana")

	// if err != nil {
	// 	fmt.Println("Error :", err)
	// }

	//GET all
	for results.Next() {
		var (
			id     int
			f_name string
			l_name string
		)
		err = results.Scan(&id, &f_name, &l_name)
		if err != nil {
			fmt.Println("Unable to parse row")
		}

		fmt.Printf("ID : %d Name: %s %s \n", id, f_name, l_name)
	}

}
