package handlers

import (
	"database/sql"
	"fmt"

	"github.com/ltheinrich/gorum-clig/internal/pkg/db"
)

// ListCategories handler
func ListCategories() (err error) {
	// query database
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT id, categoryname, sort FROM categories ORDER BY sort;`)
	if err != nil {
		return err
	}

	// loop through rows
	for rows.Next() {
		// read from row
		var id, sort int
		var categoryname string
		err = rows.Scan(&id, &categoryname, &sort)
		if err != nil {
			return err
		}

		// print out
		fmt.Printf("ID: %v - Name: %v - Sort: %v\n", id, categoryname, sort)
	}

	// return
	return nil
}
