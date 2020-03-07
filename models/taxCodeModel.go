package models

import (
	"github.com/ilham13/rest-api/config"
)

// TaxCodeModel struct
type TaxCodeModel struct {
	TaxCodeID int    `json:"tax_code_id"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
}

// GetTaxCode get all data tax code
func GetTaxCode() []TaxCodeModel {
	db, err := config.Connect()
	checkErr(err)
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

// CreateTaxCode create new data tax code
func CreateTaxCode(tc TaxCodeModel) (TaxCodeModel, error) {
	db, err := config.Connect()
	checkErr(err)
	defer db.Close()

	sqlQuery := "INSERT INTO tax_code(tax_code_id, name, status) VALUES (?, ?, ?)"

	rows, err := db.Prepare(sqlQuery)
	checkErr(err)

	_, err = rows.Exec(tc.TaxCodeID, tc.Name, tc.Status)
	checkErr(err)

	return tc, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
