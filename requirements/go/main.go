package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type idBalance struct {
	id      int
	balance int
}

var db *sql.DB

func main() {
	connStr := "user=postgres password=postgres host=host.docker.internal port=5432 dbname=dbname sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("All good database works")

	http.HandleFunc("/topup", topUp)
	http.HandleFunc("/getBalance", getBalance)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
