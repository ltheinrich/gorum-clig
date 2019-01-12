package handlers

import (
	"database/sql"
	"fmt"

	"github.com/ltheinrich/clig/pkg/db"
)

// ListBoards handler
func ListBoards() (err error) {
	// query database
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT boards.id, boards.boardname, boards.boarddescription, boards.boardicon, categories.id, categories.categoryname,
							           boards.sort + categories.sort FROM boards INNER JOIN categories ON boards.category = categories.id
												 ORDER BY boards.sort + categories.sort;`)
	if err != nil {
		return err
	}

	// loop through rows
	for rows.Next() {
		// read from row
		var id, sort, category int
		var name, description, icon, categoryname string
		err = rows.Scan(&id, &name, &description, &icon, &category, &categoryname, &sort)
		if err != nil {
			return err
		}

		// print out
		fmt.Printf("ID: %v - Name: %v - Description: %v - Icon: %v - Category: %v (%v) - Sort: %v\n", id, name, description, icon, category, categoryname, sort)
	}

	return nil
}
