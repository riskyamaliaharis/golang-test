package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type Financing struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
	Sub   string `json:"sub"`
}

type ConventionalOsf struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Tenor  string `json:"tenor"`
	Grade  string `json:"grade"`
	Rate   int    `json:"rate"`
}
type ConventionalInvoice struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Tenor  string `json:"tenor"`
	Grade  string `json:"grade"`
	Rate   int    `json:"rate"`
}

type ProductiveInvoice struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Grade  string `json:"grade"`
	Rate   int    `json:"rate"`
}

type Reksadana struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Raturn int    `json:"return"`
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func main() {
	db, err = gorm.Open("mysql", "root:@/golang_try?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection estabilished")
	}

	db.AutoMigrate(&Financing{})
	db.AutoMigrate(&ConventionalOsf{})
	db.AutoMigrate(&ConventionalInvoice{})
	db.AutoMigrate(&ProductiveInvoice{})
	db.AutoMigrate(&Reksadana{})
	handleRequests()
}

func handleRequests() {
	log.Println("Start the development server at http://172.0.0.1:8084")
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/api/finance", createFinancing).Methods(("POST"))
	myRouter.HandleFunc("/api/finance", getFinancing).Methods(("GET"))
	myRouter.HandleFunc("/api/finance/convinvoice", createConventionalInvoice).Methods(("POST"))
	myRouter.HandleFunc("/api/finance/convinvoice", getConventionalInvoice).Methods(("GET"))
	myRouter.HandleFunc("/api/finance/convosf", createConventionalOsf).Methods(("POST"))
	myRouter.HandleFunc("/api/finance/convosf", getConventionalOsf).Methods(("GET"))
	myRouter.HandleFunc("/api/finance/productiveinvoice", createProductiveInvoice).Methods(("POST"))
	myRouter.HandleFunc("/api/finance/productiveinvoice", getProductiveInvoice).Methods(("GET"))
	myRouter.HandleFunc("/api/finance/reksadana", createReksadana).Methods(("POST"))
	myRouter.HandleFunc("/api/finance/reksadana", getReksadana).Methods(("GET"))

	log.Fatal(http.ListenAndServe(":8084", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome !!")
}

// Products

func createReksadana(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var reksadana Reksadana
	json.Unmarshal(payloads, &reksadana)
	db.Create(&reksadana)

	res := Result{Code: 200, Data: reksadana, Message: "Success create reksadana"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func getReksadanas(w http.ResponseWriter, r *http.Request) {
	reksadanas := []Reksadana{}

	db.Find(&reksadanas)
	res := Result{Code: 200, Data: reksadanas, Message: "Success get reksadanas"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}
func getReksadana(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reksadanaID := vars["id"]

	var reksadana Reksadana
	db.First(&reksadana, reksadanaID)
	res := Result{Code: 200, Data: reksadana, Message: "Success get reksadanas"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// Financing

func createFinancing(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var financing Financing
	json.Unmarshal(payloads, &financing)
	db.Create(&financing)

	res := Result{Code: 200, Data: financing, Message: "Success create financing"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func getFinancing(w http.ResponseWriter, r *http.Request) {
	financing := []Financing{}

	db.Find(&financing)
	res := Result{Code: 200, Data: financing, Message: "Success get financing"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

// ConventionalInvoice

func createConventionalInvoice(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var convInvoice ConventionalInvoice
	json.Unmarshal(payloads, &convInvoice)
	db.Create(&convInvoice)

	res := Result{Code: 200, Data: convInvoice, Message: "Success create convInvoice"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func getConventionalInvoice(w http.ResponseWriter, r *http.Request) {
	convInvoice := []ConventionalInvoice{}

	db.Find(&convInvoice)
	res := Result{Code: 200, Data: convInvoice, Message: "Success get convInvoice"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

// ConventionalOsf

func createConventionalOsf(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var convOsf ConventionalOsf
	json.Unmarshal(payloads, &convOsf)
	db.Create(&convOsf)

	res := Result{Code: 200, Data: convOsf, Message: "Success create convOsf"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func getConventionalOsf(w http.ResponseWriter, r *http.Request) {
	convOsf := []ConventionalOsf{}

	db.Find(&convOsf)
	res := Result{Code: 200, Data: convOsf, Message: "Success get convOsf"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

// ConventionalInvoice

func createProductiveInvoice(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)
	var prodInvoice ProductiveInvoice
	json.Unmarshal(payloads, &prodInvoice)
	db.Create(&prodInvoice)

	res := Result{Code: 200, Data: prodInvoice, Message: "Success create prodInvoice"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)

}

func getProductiveInvoice(w http.ResponseWriter, r *http.Request) {
	prodInvoice := []ProductiveInvoice{}

	db.Find(&prodInvoice)
	res := Result{Code: 200, Data: prodInvoice, Message: "Success get prodInvoice"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}
