package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type getBalStruct struct {
	Id_user int `json:"Id_user"`
}

type balance struct {
	Value int `json:"Value"`
}

func decodeJSONBal(r *http.Request) (*getBalStruct, error) {
	var content getBalStruct
	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&content)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &content, nil
}

func responseJSON(w http.ResponseWriter, value int) error {
	data := balance{value}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	return err
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	inData, err := decodeJSONBal(r)
	if err != nil || inData.Id_user < 0 {
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
		responseJSON(w, currUser.balance)
	}
}
