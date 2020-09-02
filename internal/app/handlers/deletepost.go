package handlers

import (
	"errors"
	"fmt"

	"github.com/ltheinrich/gorum-clig/internal/pkg/db"
)

// DeletePost handler
func DeletePost() (err error) {
	// check variables
	if *id == -1 {
		return errors.New("Use the -id argument")
	}

	// delete from database
	_, err = db.DB.Exec(`DELETE FROM posts WHERE id = $1;`, *id)
	if err != nil {
		return err
	}

	// print success and return
	fmt.Println("Post deleted")
	return nil
}
