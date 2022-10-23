package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type reservingStruct struct {
	Id_user    int `json:"Id_user"`
	Id_service int `json:"Id_service"`
	Id_order   int `json:"Id_order"`
	Value      int `json:"Value"`
}

func decodeJSONRes(r *http.Request) (*reservingStruct, error) {
	var content reservingStruct
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&content)
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func addNewOrder(inData *reservingStruct, currUser *idBalance) error {
	if currUser.balance >= inData.Value {
		newBalance := currUser.balance - inData.Value
		_, err := DB.Exec(`INSERT INTO Orders VALUES ( $1, $2, $3, $4 )`,
			currUser.id, inData.Id_service, inData.Id_order, inData.Value)
		if err != nil {
			return err
		}
		_, err = DB.Exec(`UPDATE Users SET balance = $1, res_balance = $2 WHERE id_user = $3`,
			newBalance, inData.Value, currUser.id)
		if err != nil {
			_, err = DB.Exec(`DELETE FROM Users WHERE id_order = $1`,
				inData.Id_order)
		}
		return err
	}
	return errors.New("not enough money")
}

func reserving(w http.ResponseWriter, r *http.Request) {
	inData, err := decodeJSONRes(r)
	if err != nil || inData.Id_user < 0 || inData.Id_order < 0 ||
		inData.Id_service < 0 || inData.Value < 0 {
		sendErrorJSON(w, http.StatusBadRequest, "invalid_value",
			"Client sent an unsupported value")
		return
	}
	currUser, err := getUserById(DB, inData.Id_user)
	if err != nil || currUser.id == -1 {
		sendErrorJSON(w, http.StatusNotFound, "invalid_id_user",
			"Client provided an invalid User ID")
		return
	} else {
		err = addNewOrder(inData, currUser)
		if err != nil {
			//sql error
			return
		}
	}
}
