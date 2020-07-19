package controllers

import (
	"encoding/json"
	"gomux/main/apimaster/models"
	"gomux/main/apimaster/usecases"
	"gomux/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//TransaksiHandler TransaksiHandler
type TransaksiHandler struct {
	TransaksiUseCase usecases.TransaksiUseCase
}

//TransaksiController TransaksiController
func TransaksiController(r *mux.Router, service usecases.TransaksiUseCase) {
	TransaksiHandler := TransaksiHandler{service}
	r.HandleFunc("/alltransaksi", TransaksiHandler.AllTransaksi).Methods(http.MethodGet)
	r.HandleFunc("/transaksi/{id}", TransaksiHandler.TransaksiByID).Methods(http.MethodGet)
	r.HandleFunc("/transaksi", TransaksiHandler.AddTransaksi).Methods(http.MethodPost)
	r.HandleFunc("/transaksi/{id}", TransaksiHandler.UpdateTransaksi).Methods(http.MethodPut)
	r.HandleFunc("/transaksi/{id}", TransaksiHandler.DeleteTransaksi).Methods(http.MethodDelete)
}

//AllTransaksi AllTransaksi
func (s TransaksiHandler) AllTransaksi(w http.ResponseWriter, r *http.Request) {
	allTransaksi, err := s.TransaksiUseCase.GetAllTransaksi()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Println("Data Not Found")
	}
	transaksiResponse := utils.Response{Status: http.StatusOK, Messages: "Data of Transaksi", Data: allTransaksi}
	byteOfTransaksiResponse, err := json.Marshal(transaksiResponse)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
		log.Println("Oops something when wrong")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(byteOfTransaksiResponse)
}

//AddTransaksi InsertTransaksi
func (s TransaksiHandler) AddTransaksi(w http.ResponseWriter, r *http.Request) {
	var newTransaksi models.Transaksi
	_ = json.NewDecoder(r.Body).Decode(&newTransaksi)
	err := s.TransaksiUseCase.AddTransaksi(newTransaksi)
	if err != nil {
		w.Write([]byte("Insert Failed Cannot null"))
		log.Print(err)
	} else {
		w.Write([]byte("Insert Success"))
		log.Println("Insert Success")
	}
}

//TransaksiByID TransaksiById
func (s TransaksiHandler) TransaksiByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idTransaksi := param["id"]
	transaksi, err := s.TransaksiUseCase.GetTransaksiByID(idTransaksi)
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Println("Data not found")
	}
	transaksiResponseByID := utils.Response{Status: http.StatusOK, Messages: "Data of Transaksi", Data: transaksi}
	byteOfResponseTransaksiByID, err2 := json.Marshal(transaksiResponseByID)
	if err2 != nil {
		w.Write([]byte("Oops something when wrong"))
		log.Println("Oops something when wrong")
	} else if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfResponseTransaksiByID)
	}
}

// UpdateTransaksi UpdateTransaksi
func (s TransaksiHandler) UpdateTransaksi(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idTransaksi := param["id"]
	var changeTransaksi models.Transaksi
	_ = json.NewDecoder(r.Body).Decode(&changeTransaksi)
	err := s.TransaksiUseCase.UpdateTransaksiByID(idTransaksi, changeTransaksi)
	if err != nil {
		w.Write([]byte("Id Not Found"))
		log.Println("Data not found")
	} else {
		w.Write([]byte("Data Updated"))
		log.Println("Data Updated")
	}

}

//DeleteTransaksi DeleteById
func (s TransaksiHandler) DeleteTransaksi(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idTransaksi := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&s)
	err := s.TransaksiUseCase.DeleteTransaksiByID(idTransaksi)
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Println("Data not found")
	} else {
		w.Write([]byte("Data Deleted"))
		log.Println("Data Deleted")
	}
}
