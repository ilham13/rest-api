package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ilham13/rest-api/models"
	"github.com/julienschmidt/httprouter"
)

type (
	TaxCodeController struct{}
)

type TaxCodeResponse struct {
	ResponseCode        string                `json:"response_code"`
	ResponseDescription string                `json:"response_description"`
	Data                []models.TaxCodeModel `json:"data"`
}

func NewTaxCodeController() *TaxCodeController {
	return &TaxCodeController{}
}

func (tc TaxCodeController) GetTaxCode(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	taxCode := models.GetTaxCode()

	d := TaxCodeResponse{"1", "success", taxCode}

	uj, _ := json.Marshal(d)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", uj)
}
