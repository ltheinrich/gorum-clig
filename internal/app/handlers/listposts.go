package handlers

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ltheinrich/clig/internal/pkg/db"
)

// ListPosts handler
func ListPosts() (err error) {
	// query database
	var rows *sql.Rows
	rows, err = db.DB.Query(`SELECT posts.id, posts.thread, posts.author, users.username, posts.created, posts.content
													FROM posts INNER JOIN users ON posts.author = users.id ORDER BY posts.created DESC;`)
	if err != nil {
		return err
	}

	// loop through rows
	for rows.Next() {
		// read from row
		var id, thread, author int
		var created int64
		var username, content string
		err = rows.Scan(&id, &thread, &author, &username, &created, &content)
		if err != nil {
			return err
		}

		// print out
		fmt.Printf("ID: %v - Thread: %v - Author: %v (%v) - Created: %v - Content: %v\n",
			id, thread, author, username, time.Unix(created, 0).Format("02.01.2006 at 15:04:05"), content)
	}

	// return
	return nil
}
