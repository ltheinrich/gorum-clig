package handlers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ltheinrich/clig/pkg/db"
)

// ListThreads handler
func ListThreads() (err error) {
	// query database
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT threads.id, threads.threadname, threads.board, threads.author, users.username, threads.created, threads.content
													FROM threads INNER JOIN users ON threads.author = users.id ORDER BY threads.created DESC;`)
	if err != nil {
		return err
	}

	// loop through rows
	for rows.Next() {
		// read from row
		var id, board, author int
		var created int64
		var name, username, content string
		err = rows.Scan(&id, &name, &board, &author, &username, &created, &content)
		if err != nil {
			return err
		}

		// print out
		fmt.Printf("ID: %v - Name: %v - Board: %v - Author: %v (%v) - Created: %v - Content: %v\n",
			id, name, board, author, username, time.Unix(created, 0).Format("02.01.2006 at 15:04:05"), content)
	}

	// return
	return nil
}
