package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type idBalance struct {
	id          int
	balance     int
	res_balance int
}

var DB *sql.DB

func main() {
	connStr := "user=postgres password=postgres host=host.docker.internal port=5432 dbname=dbname sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("All good")

	http.HandleFunc("/topUp.json", topUp)
	http.HandleFunc("/getBalance.json", getBalance)
	http.HandleFunc("/reserving.json", reserving)
	http.HandleFunc("/acceptFromReserve.json", acceptFromReserve)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
