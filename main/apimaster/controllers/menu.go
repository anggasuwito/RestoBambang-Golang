package controllers

import (
	"encoding/json"
	"fmt"
	"gomux/main/apimaster/models"
	"gomux/main/apimaster/usecases"
	"gomux/utils"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
	r.HandleFunc("/getimage", MenuHandler.GetImage).Methods(http.MethodGet)
	r.HandleFunc("/postimage", MenuHandler.PostImage).Methods(http.MethodPost)
}

//GetImage AddImage
func (s MenuHandler) GetImage(w http.ResponseWriter, r *http.Request) {

	//D:\APP\BE1\RestoBambang-Golang-first-modify\main
	dir, err := os.Getwd()
	fmt.Println("dir = ", dir)
	fmt.Println("dir err = ", err)

	//D:\APP\BE1\RestoBambang-Golang-first-modify\main\uploads\arielTatum.jpg
	fileLocationWithDir := filepath.Join(dir, "uploads", `facebookIcon.png`)
	fmt.Println("fileLocationWithDir = ", fileLocationWithDir)

	//D:\APP\BE1\RestoBambang-Golang-first-modify\uploads\arielTatum.jpg
	fileLocation := filepath.Join("uploads", `arielTatum.jpg`)
	fmt.Println("file Location = ", fileLocation)

	//untuk menampilkan foto pada responsenya
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, fileLocationWithDir)
}

//PostImage AddImage
func (s MenuHandler) PostImage(w http.ResponseWriter, r *http.Request) {
	//membatasi max upload 1024
	r.ParseMultipartForm(1024)

	//mengambil key dari file dari formdata
	image, handlerImage, err := r.FormFile("image")
	if err != nil {
		log.Println(`Error while parsing file`, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println("image = ", image)
	fmt.Println("handler = ", handlerImage)
	fmt.Println("handler filename = ", handlerImage.Filename)
	fmt.Println("handler header = ", handlerImage.Header)
	fmt.Println("handler size = ", handlerImage.Size)

	//mengambil key dari text dari formdata
	var imageModel models.Image
	data := r.FormValue("data")
	_ = json.Unmarshal([]byte(data), &imageModel)
	fmt.Println("imagemodel ID = ", imageModel.IDImage)
	fmt.Println("imagemodel Name = ", imageModel.NameImage)

	//size adalah int64
	fmt.Printf("handlerSizeOriginalType = %v %T \n", handlerImage.Size, handlerImage.Size)

	newPhotoName := CopyImageAndRename(image, handlerImage)

	//convert int64 to int to String
	imageSize := strconv.Itoa(int(handlerImage.Size))
	response := "old image name = " + handlerImage.Filename + "\nold image size = " + imageSize + "\nnew photo name = " + newPhotoName

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

//CopyImageAndRename CopyImageAndRename
func CopyImageAndRename(photo multipart.File, handler *multipart.FileHeader) string {
	dir, err := os.Getwd()
	fmt.Println("dir = ", dir)
	fmt.Println("dir err = ", err)

	rand.Seed(time.Now().UnixNano())
	min := 11111111111
	max := 99999999999
	newPhotoName := "userPhoto-" + strconv.Itoa(rand.Intn(max-min+1)+min) + filepath.Ext(handler.Filename)
	fmt.Println("new Photo Name = ", newPhotoName)
	fileLocation := filepath.Join("uploads", newPhotoName)
	fmt.Println("File location = ", fileLocation)

	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	fmt.Println("targetFile = ", targetFile)
	fmt.Println("targetFile err = ", err)

	copyFile, err := io.Copy(targetFile, photo)
	fmt.Println("copyFile = ", copyFile)
	fmt.Println("copyFile err = ", err)

	return newPhotoName
}

//AllMenus AllMenus
func (s MenuHandler) AllMenus(w http.ResponseWriter, r *http.Request) {
	allMenus, err := s.MenuUseCase.GetAllMenus()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("Data Menu Not Found"))
		log.Println("Data Menu Not Found")
	}
	menuResponse := utils.Response{Status: http.StatusOK, Messages: "Data of Menu", Data: allMenus}
	byteOfMenuResponse, err := json.Marshal(menuResponse)
	if err != nil {
		w.Write([]byte("Oops something when wrong from Menu"))
		log.Println("Oops something when wrong from Menu")
	}
	log.Println("get all menu success")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteOfMenuResponse)
}

//AddMenu InsertMenu
func (s MenuHandler) AddMenu(w http.ResponseWriter, r *http.Request) {
	var newMenu models.Menu
	_ = json.NewDecoder(r.Body).Decode(&newMenu)
	err := s.MenuUseCase.AddMenu(newMenu)
	if err != nil {
		w.Write([]byte("Insert Menu Failed Cannot null"))
		log.Print(err)
	} else {
		w.Write([]byte("Insert Menu Success"))
		log.Println("Insert Menu Success")
	}
}

//MenuByID MenuById
func (s MenuHandler) MenuByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idMenu := param["id"]
	menu, err := s.MenuUseCase.GetMenuByID(idMenu)
	if err != nil {
		w.Write([]byte("Data Menu Not Found"))
		log.Println("Data Menu not found")
	}
	menuResponse := utils.Response{Status: http.StatusOK, Messages: "Data of Menu", Data: menu}
	byteOfMenuResponseByID, err2 := json.Marshal(menuResponse)
	if err2 != nil {
		w.Write([]byte("Oops something when wrong from Menu"))
		log.Println("Oops something when wrong from Menu")
	} else if err == nil {
		log.Println("get menu from id success")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(byteOfMenuResponseByID)
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
		w.Write([]byte("Id Menu Not Found"))
		log.Println("Data Menu not found")
	} else {
		w.Write([]byte("Data Menu Updated"))
		log.Println("Data Menu Updated")
	}

}

//DeleteMenu DeleteById
func (s MenuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idMenu := param["id"]
	_ = json.NewDecoder(r.Body).Decode(&s)
	err := s.MenuUseCase.DeleteMenusByID(idMenu)
	if err != nil {
		w.Write([]byte("Data Menu Not Found"))
		log.Println("Data Menu not found")
	} else {
		w.Write([]byte("Data Menu Deleted"))
		log.Println("Data Menu Deleted")
	}
}
