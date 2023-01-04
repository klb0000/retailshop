package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/klb0000/retailshop"
)

var defaultDB = retailshop.DefaultDB

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "home page")
}

// func handleCompleteTransaction(w http.ResponseWriter, r *http.Request) {
// 	data, err := io.ReadAll(r)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	jDecoder := json.NewDecoder(bytes.NewBuffer(data))
// 	var m map[string]interface{}
// 	jDecoder.Decode(&m)
// }

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

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	var pd retailshop.Product
	if err := json.NewDecoder(bytes.NewReader(body)).
		Decode(&pd); err != nil {

		w.WriteHeader(http.StatusConflict)
		return
	}
	fmt.Println(pd)

	if err := retailshop.SaveProduct(&pd, defaultDB); err != nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func Serve(addr string) {
	// http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getAll", handleGetAllPrdocuts)
	http.HandleFunc("/getByID", GetById)
	http.HandleFunc("/createProduct", CreateProduct)
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("client root: %s", root)
	http.Handle("/", http.FileServer(http.Dir(root)))
	log.Fatal(http.ListenAndServe(addr, nil))
}
