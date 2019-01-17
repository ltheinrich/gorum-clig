package handlers

import (
	"errors"
	"fmt"

	"github.com/ltheinrich/clig/internal/pkg/db"
)

// DeleteBoard handler
func DeleteBoard() (err error) {
	// check variables
	if *id == -1 {
		return errors.New("Use the -id argument")
	}

	// delete from database
	_, err = db.DB.Exec(`DELETE FROM boards WHERE id = $1;`, *id)
	if err != nil {
		return err
	}

	// print success and return
	fmt.Println("Board deleted")
	return nil
}
