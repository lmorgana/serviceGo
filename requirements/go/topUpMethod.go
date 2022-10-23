package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type topUpStruct struct {
	Id_user int `json:"Id_user"`
	Value   int `json:"Value"`
}

func getUserById(db *sql.DB, id_user int) (*idBalance, error) {
	currUser := idBalance{-1, -1, -1}
	rows, err := db.Query(`SELECT * FROM Users WHERE id_user = $1`, id_user)
	if err != nil {
		return &currUser, err
	}
	if rows.Next() {
		err = rows.Scan(&currUser.id, &currUser.balance, &currUser.res_balance)
	}
	return &currUser, err
}

func decodeJSON(r *http.Request) (*topUpStruct, error) {
	var content topUpStruct
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func mkNewUser(content *topUpStruct) error {
	cont := `INSERT INTO Users VALUES ( $1, $2 )`
	_, err := DB.Exec(cont, content.Id_user, content.Value)
	if err != nil {
		return err
	}
	return nil
}

func updateBalance(currUser *idBalance) error {
	_, err := DB.Exec(`UPDATE Users SET balance = $1 WHERE id_user = $2`,
		currUser.balance, currUser.id)
	return err
}

func topUp(w http.ResponseWriter, r *http.Request) {
	inData, err := decodeJSON(r)
	if err != nil || inData.Id_user < 0 || inData.Value < 0 {
		sendErrorJSON(w, http.StatusBadRequest, "invalid_value",
			"Client sent an unsupported value")
		return
	}
	currUser, err := getUserById(DB, inData.Id_user)
	if err != nil {
		sendErrorJSON(w, http.StatusNotFound, "invalid_id_user",
			"Client provided an invalid User ID")
		return
	}
	if currUser.id != inData.Id_user {
		err = mkNewUser(inData)
		if err != nil {
			//sql return some errors
			return
		}
	} else if inData.Value != 0 {
		currUser.balance += inData.Value
		err = updateBalance(currUser)
		if err != nil {
			//sql return some errors
			return
		}
	}
}
