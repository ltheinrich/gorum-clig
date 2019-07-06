package handlers

import (
	"errors"
	"fmt"

	"github.com/NathanNr/gorum-clig/internal/pkg/db"
)

// DeleteThread handler
func DeleteThread() (err error) {
	// check variables
	if *id == -1 {
		return errors.New("Use the -id argument")
	}

	// delete from database
	_, err = db.DB.Exec(`DELETE FROM threads WHERE id = $1;`, *id)
	if err != nil {
		return err
	}

	// print success and return
	fmt.Println("Thread deleted")
	return nil
}
