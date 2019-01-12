package handlers

import (
	"errors"
	"fmt"

	"github.com/ltheinrich/clig/pkg/db"
)

// DeleteCategory handler
func DeleteCategory() (err error) {
	// check variables
	if *id == -1 {
		return errors.New("Use the -id argument")
	}

	// delete from database
	_, err = db.DB.Exec(`DELETE FROM categories WHERE id = $1;`, *id)
	if err != nil {
		return err
	}

	fmt.Println("Category deleted")

	return nil
}
