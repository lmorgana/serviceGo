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

//func write(w http.ResponseWriter, r *http.Request) {
//	var data topUpStruct
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	json.NewEncoder(w).Encode(data)
//}

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
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
