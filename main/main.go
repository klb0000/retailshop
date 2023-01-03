package main

import (
	"fmt"
	"log"

	"github.com/klb0000/retailshop"
	"github.com/klb0000/retailshop/server"
)

var DefaultDB = retailshop.DefaultDB

var sampleId = "c2d766ca982eca8304150849735ffef9"

func main() {
	pd := retailshop.Product{
		ID:    "234324",
		Name:  "screwdriver",
		Price: 332,
	}
	fmt.Println(pd)

	products, err := retailshop.GetAllProducts(DefaultDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(products))
	server.Serve("localhost:8080")

}
