package main

import (
	"encoding/json"
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
	newBalance := currUser.balance - inData.Value
	newResBalance := currUser.res_balance + inData.Value
	_, err := DB.Exec(`INSERT INTO Orders VALUES ( $1, $2, $3, $4 )`,
		currUser.id, inData.Id_service, inData.Id_order, inData.Value)
	if err != nil {
		return err
	}
	_, err = DB.Exec(`UPDATE Users SET balance = $1, res_balance = $2 WHERE id_user = $3`,
		newBalance, newResBalance, currUser.id)
	if err != nil {
		_, err = DB.Exec(`DELETE FROM Users WHERE id_order = $1`,
			inData.Id_order)
	}
	return err
}

func reserving(w http.ResponseWriter, r *http.Request) {
	inData, err := decodeJSONRes(r)
	if err != nil || !checkSliceForInterval(2147483648, 0,
		inData.Id_user, inData.Id_order, inData.Id_service, inData.Value) {
		sendErrorJSON(w, http.StatusBadRequest, "invalid_value",
			"Client sent an unsupported value")
		return
	}
	currUser, err := getUserById(DB, inData.Id_user)
	if err != nil || currUser.id == -1 {
		sendErrorJSON(w, http.StatusUnauthorized, "invalid_id_user",
			"Client provided an invalid User ID")
		return
	}
	order, err := getOrderById(DB, inData.Id_order, 4)
	if err != nil || order.Id_order >= 0 {
		sendErrorJSON(w, http.StatusUnauthorized, "invalid_order_values",
			"Client sent a wrong order values")
		return
	}
	if currUser.balance < inData.Value ||
		currUser.res_balance+inData.Value > 2147483647 {
		sendErrorJSON(w, http.StatusPreconditionFailed, "balance_limit",
			"Value exceed user balance limit")
		return
	}
	err = addNewOrder(inData, currUser)
	if err != nil {
		sendErrorJSON(w, http.StatusInternalServerError, "", "")
		return
	}
}
