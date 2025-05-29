package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func checkError(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func (app *App) Initialise(BDuser string, DBpasswd string, DBhost string, DBname string) error {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v)/%v", BDuser, DBpasswd, DBhost, DBname)
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	checkError(err)
	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
	return nil

}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)

}

func sendError(w http.ResponseWriter, statusCode int, err string) {
	errorMessage := map[string]string{"error": err}
	sendResponse(w, statusCode, errorMessage)
}

func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	sendResponse(w, http.StatusOK, products)
}

func (app *App) getOneProduct(w http.ResponseWriter, r *http.Request) {
	key, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Invalid ID")
		return
	}
	product, err := getAProduct(app.DB, key)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusInternalServerError, "No Product with this ID")
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	sendResponse(w, http.StatusOK, product)
}
func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {
	var p product
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = addProduct(app.DB, &p)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}
	sendResponse(w, http.StatusCreated, p)
}

func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	key, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Invalid ID")
		return
	}
	var p product
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	p.ID = key
	updateAProduct(app.DB, &p)
	sendResponse(w, http.StatusOK, p)

}
func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	key, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		sendError(w, http.StatusInternalServerError, "Invalid ID")
		return
	}
	deleteAProduct(app.DB, key)
	sendResponse(w, http.StatusOK, "Object is deleted")

}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/products", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/products/{id}", app.getOneProduct).Methods("GET")
	app.Router.HandleFunc("/product", app.createProduct).Methods("POST")
	app.Router.HandleFunc("/products/{id}", app.updateProduct).Methods("PUT")
	app.Router.HandleFunc("/products/{id}", app.deleteProduct).Methods("DELETE")
}
