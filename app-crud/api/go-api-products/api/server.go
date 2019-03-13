package api

import (
	"net/http"
	"fmt"
	"log"
	"go-api-products/api/routes"
	"go-api-products/api/models"
)

func Run() {

	db := models.Connect()
	db.DropTableIfExists(&models.Product{})
	if !db.HasTable(&models.Product{}) {
		db.Debug().CreateTable(&models.Product{})
	}
	db.Close()
	listen(9000)
}

func listen(p int) {
	port := fmt.Sprintf(":%d", p)
	fmt.Printf("\nListening Port %s...\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(port, routes.LoadCors(r)))
}