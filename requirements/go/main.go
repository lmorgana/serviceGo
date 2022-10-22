package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

type topUpStruct struct {
	Id_user int `json:"Id_user"`
	Value   int `json:"Value"`
}

var db *sql.DB

func write(w http.ResponseWriter, r *http.Request) {
	var data topUpStruct
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
}

func topUp(w http.ResponseWriter, r *http.Request) {
	var content topUpStruct
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result is", content)
	cont := `INSERT INTO users VALUES ( $1, $2 )`
	_, err = db.Exec(cont, content.Id_user, content.Value)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	connStr := "user=postgres password=postgres host=host.docker.internal port=5432 dbname=dbname sslmode=disable"
	var err error
	db, err := sql.Open("postgres", connStr)
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
