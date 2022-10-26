package main

import (
	"net/http"
)

func changeOrderStatRef(inData *reservingStruct, currUser *idBalance) error {
	newResBalance := currUser.res_balance - inData.Value
	newBalance := currUser.balance + inData.Value
	_, err := DB.Exec(`UPDATE Users SET balance = $1, res_balance = $2 WHERE id_user = $3`,
		newBalance, newResBalance, currUser.id)
	if err != nil {
		return err
	}
	_, err = DB.Exec(`UPDATE Orders SET status = $1 WHERE id_order = $2`,
		2, inData.Id_order)
	return err
}

func refuseFromReserve(w http.ResponseWriter, r *http.Request) {
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
	order, err := getOrderById(DB, inData.Id_order)
	if err != nil ||
		order.Id_order != inData.Id_order ||
		order.Id_service != inData.Id_service ||
		order.Id_user != inData.Id_user ||
		order.Value != inData.Value {
		sendErrorJSON(w, http.StatusUnauthorized, "invalid_order_values",
			"Client sent a wrong order values")
		return
	}
	if currUser.res_balance < order.Value ||
		currUser.balance+order.Value > 2147483647 {
		sendErrorJSON(w, http.StatusPreconditionFailed, "balance_limit",
			"Value exceed user balance limit")
		return
	}
	err = changeOrderStatRef(inData, currUser)
	if err != nil {
		sendErrorJSON(w, http.StatusInternalServerError, "", "")
		return
	}
}
