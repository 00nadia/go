package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "dbname"
)

type info struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Pet       string `json:"pet"`
}

func main() {

	SQLInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", SQLInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	rows, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var infoTest []info = make([]info, 0)
	for rows.Next() {
		content := info{}
		err = rows.Scan(&content.FirstName, &content.LastName, &content.Age, &content.Pet)
		if err != nil {
			panic(err)
		}
		infoTest = append(infoTest, content)
	}
	fmt.Println(infoTest)
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	newinfoTest, err := json.Marshal(infoTest)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		w.Write(newinfoTest)
	})

	http.ListenAndServe(":4000", nil)

}
