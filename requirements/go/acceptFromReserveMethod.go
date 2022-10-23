package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

func changeOrderStat(inData *reservingStruct, currUser *idBalance) error {
	if currUser.balance >= inData.Value {
		newResBalance := currUser.res_balance - inData.Value
		_, err := DB.Exec(`UPDATE Users SET res_balance = $1 WHERE id_user = $2`,
			newResBalance, currUser.id)
		if err != nil {
			return err
		}
		_, err = DB.Exec(`UPDATE Orders SET status = $1 WHERE id_order = $2`,
			1, inData.Id_order)
		return err
	}
	return errors.New("not enough money")
}

func getOrderById(db *sql.DB, id_order int) (*reservingStruct, error) {
	order := reservingStruct{-1, -1, -1, -1}
	rows, err := db.Query(`SELECT * FROM Orders WHERE id_order = $1 AND status = 0`, id_order)
	if err != nil {
		fmt.Println(err)
		return &order, err
	}
	var i int
	if rows.Next() {
		err = rows.Scan(&order.Id_user, &order.Id_service, &order.Id_order, &order.Value, &i)
	}
	return &order, err
}

func acceptFromReserve(w http.ResponseWriter, r *http.Request) {
	inData, err := decodeJSONRes(r)
	if err != nil || inData.Id_user < 0 || inData.Id_order < 0 || inData.Id_service < 0 || inData.Value < 0 {
		sendErrorJSON(w, http.StatusBadRequest, "invalid_value",
			"Client sent an unsupported value")
		return
	}
	currUser, err := getUserById(DB, inData.Id_user)
	if err != nil || currUser.id == -1 {
		sendErrorJSON(w, http.StatusNotFound, "invalid_id_user",
			"Client provided an invalid User ID")
		return
	}
	order, err := getOrderById(DB, inData.Id_order)
	if err != nil ||
		order.Id_order != inData.Id_order ||
		order.Id_service != inData.Id_service ||
		order.Id_user != inData.Id_user ||
		order.Value != inData.Value {
		sendErrorJSON(w, http.StatusNotFound, "invalid_order_values",
			"Client sent a wrong order values")
		return
	} else {
		err = changeOrderStat(inData, currUser)
		if err != nil {
			//sql error
			return
		}
	}
}
