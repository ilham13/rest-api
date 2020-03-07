package models

import (
	"fmt"

	"github.com/ilham13/rest-api/config"
)

type TaxCodeModel struct {
	TaxCodeID int    `json:"tax_code_id"`
	Name      string `json:"name"`
}

func GetTaxCode() []TaxCodeModel {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT tax_code_id, name FROM tax_code")
	config.CheckErr(err)

	defer rows.Close()

	var taxes []TaxCodeModel
	for rows.Next() {
		tax := TaxCodeModel{}
		err := rows.Scan(&tax.TaxCodeID, &tax.Name)
		config.CheckErr(err)

		taxes = append(taxes, tax)
	}
	return taxes
}
