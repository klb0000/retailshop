package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/retailshop"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "home page")
}

func handleGetAllPrdocuts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	products, err := retailshop.GetAllProducts(retailshop.DefaultDB)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(products)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func GetById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	product, err := retailshop.GetProductById(id, retailshop.DefaultDB)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	data, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func Serve(addr string) {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getAll", handleGetAllPrdocuts)
	http.HandleFunc("/getByID", GetById)
	log.Fatal(http.ListenAndServe(addr, nil))
}
