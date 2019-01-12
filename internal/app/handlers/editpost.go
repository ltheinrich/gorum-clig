package handlers

import (
	"errors"
	"fmt"

	"github.com/ltheinrich/clig/pkg/db"
)

// EditPost handler
func EditPost() (err error) {
	// check variables
	if (*id == -1) || (*thread == -1 && *author == -1 && *created == -1 && *content == "") {
		return errors.New("Use -id argument and one or many of -thread, -author, -created, -content arguments")
	}

	// update thread
	if *thread != -1 {
		_, err = db.DB.Exec(`UPDATE posts SET thread = $1 WHERE id = $2`, *thread, *id)
		if err != nil {
			return err
		}
	}

	// update author
	if *author != -1 {
		_, err = db.DB.Exec(`UPDATE posts SET author = $1 WHERE id = $2`, *author, *id)
		if err != nil {
			return err
		}
	}

	// update created
	if *created != -1 {
		_, err = db.DB.Exec(`UPDATE posts SET created = $1 WHERE id = $2`, *created, *id)
		if err != nil {
			return err
		}
	}

	// update content
	if *content != "" {
		_, err = db.DB.Exec(`UPDATE posts SET content = $1 WHERE id = $2`, *content, *id)
		if err != nil {
			return err
		}
	}

	// print success and return
	fmt.Println("Post updated")
	return nil
}
