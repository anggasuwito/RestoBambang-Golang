package controllers

import (
	"encoding/json"
	"gomux/main/master/models"
	"gomux/main/master/usecases"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//MenuHandler MenuHandler
type MenuHandler struct {
	MenuUseCase usecases.MenuUseCase
}

//MenuController MenuController
func MenuController(r *mux.Router, service usecases.MenuUseCase) {
	MenuHandler := MenuHandler{service}
	r.HandleFunc("/allmenus", MenuHandler.AllMenus).Methods(http.MethodGet)
	r.HandleFunc("/menu/{id}", MenuHandler.MenuByID).Methods(http.MethodGet)
	r.HandleFunc("/menu", MenuHandler.AddMenu).Methods(http.MethodPost)
	r.HandleFunc("/menu/{id}", MenuHandler.UpdateMenu).Methods(http.MethodPut)
	r.HandleFunc("/menu/{id}", MenuHandler.DeleteMenu).Methods(http.MethodDelete)
	r.HandleFunc("/alljenis", MenuHandler.AllJenis).Methods(http.MethodGet)

}

//AllMenus AllMenus
func (s MenuHandler) AllMenus(w http.ResponseWriter, r *http.Request) {
	allMenus, err := s.MenuUseCase.GetAllMenus()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfAllMenus, err := json.Marshal(allMenus)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(byteOfAllMenus)
}

//AddMenu InsertMenu
func (s MenuHandler) AddMenu(w http.ResponseWriter, r *http.Request) {
	var newMenu models.Menu
	_ = json.NewDecoder(r.Body).Decode(&newMenu)
	err := s.MenuUseCase.AddMenu(newMenu)
	if err != nil {
		w.Write([]byte("Insert Failed Cannot null"))
		log.Print(err)
	} else {
		w.Write([]byte("Insert Success"))
	}
}

//MenuByID MenuById
func (s MenuHandler) MenuByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idMenu := param["id"]
	menu, err := s.MenuUseCase.GetMenuByID(idMenu)
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Println("Data not found")
	}
	byteOfMenuByID, err2 := json.Marshal(menu)
	if err2 != nil {
		w.Write([]byte("Oops something when wrong"))
	} else if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfMenuByID)
	}
}

// UpdateMenu UpdateMenu
func (s MenuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idMenu := param["id"]
	var changeMenu models.Menu
	_ = json.NewDecoder(r.Body).Decode(&changeMenu)
	err := s.MenuUseCase.UpdateMenusByID(idMenu, changeMenu)
	if err != nil {
		w.Write([]byte("Id Not Found"))
		log.Println("Data not found")
	} else {
		w.Write([]byte("Data Updated"))
	}

}

//DeleteMenu DeleteById
func (s MenuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idMenu := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&s)
	err := s.MenuUseCase.DeleteMenusByID(idMenu)
	if err != nil {
		w.Write([]byte("Data Not Found"))
		log.Println("Data not found")
	} else {
		w.Write([]byte("Data Deleted"))
	}
}

//AllJenis AllJenis
func (s MenuHandler) AllJenis(w http.ResponseWriter, r *http.Request) {
	allJenis, err := s.MenuUseCase.GetAllJenis()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	byteOfAllJenis, err := json.Marshal(allJenis)
	if err != nil {
		w.Write([]byte("Oops something when wrong"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(byteOfAllJenis)
}
