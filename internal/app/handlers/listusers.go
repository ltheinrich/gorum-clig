package handlers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ltheinrich/clig/pkg/db"
)

// ListUsers handler
func ListUsers() (err error) {
	// query database
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT id, username, registered FROM users;`)
	if err != nil {
		return err
	}

	// loop through rows
	for rows.Next() {
		// read from row
		var id int
		var username, registered string
		err = rows.Scan(&id, &username, &registered)
		if err != nil {
			return err
		}

		// parse registered time
		var date time.Time
		date, err = time.Parse("2006-01-02T15:04:05", registered)
		if err != nil {
			return err
		}

		// print out
		fmt.Printf("ID: %v - Username: %v - Registered: %v\n", id, username, date.Format("02.01.2006 at 15:04:05"))
	}

	return nil
}
