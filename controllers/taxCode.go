package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ilham13/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

// TaxCodeController struct
type (
	TaxCodeController struct{}
)

// GetTaxCodeResponse response all tax code
type GetTaxCodeResponse struct {
	ResponseCode        string                `json:"response_code"`
	ResponseDescription string                `json:"response_description"`
	Data                []models.TaxCodeModel `json:"data"`
}

// CreateTaxCodeResponse response create tax code
type CreateTaxCodeResponse struct {
	ResponseCode        string              `json:"response_code"`
	ResponseDescription string              `json:"response_description"`
	Data                models.TaxCodeModel `json:"data"`
}

// NewTaxCodeController assign controller
func NewTaxCodeController() *TaxCodeController {
	return &TaxCodeController{}
}

// GetTaxCode method get all tax code
func (tc TaxCodeController) GetTaxCode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	taxCode := models.GetTaxCode()

	d := GetTaxCodeResponse{"1", "success", taxCode}

	uj, _ := json.Marshal(d)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}

// CreateTaxCode method create tax code
func (tc TaxCodeController) CreateTaxCode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.TaxCodeModel{}

	json.NewDecoder(r.Body).Decode(&u)

	create, err := models.CreateTaxCode(u)
	checkErr(err)

	/**
	 * TODO
	 * add response error, validation
	 **/

	d := CreateTaxCodeResponse{"1", "success", create}

	uj, _ := json.Marshal(d)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", uj)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
