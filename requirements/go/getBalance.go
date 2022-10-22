package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type getBalStruct struct {
	Id_user int `json:"Id_user"`
	sdf     int
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
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(data)
	return err
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	inData, err := decodeJSONBal(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	if inData.Id_user < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	currUser, err := getUserById(db, inData.Id_user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	if currUser.id == -1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	} else {
		err = responseJSON(w, currUser.balance)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
	}
}
