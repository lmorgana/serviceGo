package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type topUpStruct struct {
	Id_user int `json:"Id_user"`
	Value   int `json:"Value"`
}

func getUserById(db *sql.DB, id_user int) (*idBalance, error) {
	currUser := idBalance{-1, -1}
	rows, err := db.Query(`SELECT * FROM users WHERE id_user = $1`, id_user)
	if err != nil {
		fmt.Println(err)
		return &currUser, err
	}
	if rows.Next() {
		err = rows.Scan(&currUser.id, &currUser.balance)
	}
	return &currUser, err
}

func decodeJSON(r *http.Request) (*topUpStruct, error) {
	var content topUpStruct
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&content)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &content, nil
}

func mkNewUser(content *topUpStruct) error {
	cont := `INSERT INTO users VALUES ( $1, $2 )`
	res, err := db.Exec(cont, content.Id_user, content.Value)
	fmt.Println("exec's results", res)
	if err != nil {
		return err
	}
	return nil
}

func updateBalance(currUser *idBalance) error {
	res, err := db.Exec(`UPDATE users SET balance = $1 WHERE id_user = $2`,
		currUser.balance, currUser.id)
	fmt.Println("exec's results", res)
	return err
}

func topUp(w http.ResponseWriter, r *http.Request) {
	inData, err := decodeJSON(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	if inData.Id_user < 0 || inData.Value <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	currUser, err := getUserById(db, inData.Id_user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	fmt.Println(currUser.id, inData.Id_user)
	if currUser.id != inData.Id_user {
		err = mkNewUser(inData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	} else {
		currUser.balance += inData.Value
		err = updateBalance(currUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}
}
