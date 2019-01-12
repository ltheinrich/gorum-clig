package handlers

import (
	"errors"
	"fmt"

	"github.com/ltheinrich/clig/pkg/db"
)

// EditThread handler
func EditThread() (err error) {
	// check variables
	if (*id == -1) || (*name == "" && *board == -1 && *author == -1 && *created == -1 && *content == "") {
		return errors.New("Use -id argument and one or many of -thread, -author, -created, -content arguments")
	}

	// update name
	if *name != "" {
		_, err = db.DB.Exec(`UPDATE threads SET threadname = $1 WHERE id = $2`, *name, *id)
		if err != nil {
			return err
		}
	}

	// update board
	if *board != -1 {
		_, err = db.DB.Exec(`UPDATE threads SET board = $1 WHERE id = $2`, *board, *id)
		if err != nil {
			return err
		}
	}

	// update author
	if *author != -1 {
		_, err = db.DB.Exec(`UPDATE threads SET author = $1 WHERE id = $2`, *author, *id)
		if err != nil {
			return err
		}
	}

	// update created
	if *created != -1 {
		_, err = db.DB.Exec(`UPDATE threads SET created = $1 WHERE id = $2`, *created, *id)
		if err != nil {
			return err
		}
	}

	// update content
	if *content != "" {
		_, err = db.DB.Exec(`UPDATE threads SET content = $1 WHERE id = $2`, *content, *id)
		if err != nil {
			return err
		}
	}

	// print success and return
	fmt.Println("Thread updated")
	return nil
}
