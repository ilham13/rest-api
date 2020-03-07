package main

import (
	"log"
	"net/http"

	"github.com/ilham13/rest-api/controllers"
	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := httprouter.New()

	// Get Controller instance
	tc := controllers.NewTaxCodeController()

	/** REST API **/
	// Get a tax code resource
	router.GET("/tax_code", tc.GetTaxCode)
	router.POST("/tax_code", tc.CreateTaxCode)

	log.Println("Starting server.....")
	// http.ListenAndServe(":3000", router)
	log.Fatal(http.ListenAndServe(":3000", router))
}
