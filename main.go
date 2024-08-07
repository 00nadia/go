package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "dbname"
)

func main() {

	psglInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psglInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	//update
	sqlStatement := `
	UPDATE test 
	SET age = $1, pet =$2, last_name = $3
	WHERE first_name = $4;`

	_, err = db.Exec(sqlStatement, 35, "bat", "Wayne", "Bruce")
	if err != nil {
		panic(err)
	}
	//print
	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var firstName string
		var lastName string
		var age int
		var pet string
		err = rows.Scan(&firstName, &lastName, &age, &pet)
		if err != nil {
			panic(err)
		}
		fmt.Println(firstName, lastName, age, pet)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
