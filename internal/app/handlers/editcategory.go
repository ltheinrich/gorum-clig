package handlers

import (
	"errors"
	"fmt"

	"github.com/NathanNr/gorum-clig/internal/pkg/db"
)

// EditCategory handler
func EditCategory() (err error) {
	// check variables
	if (*id == -1) || (*name == "" && *sort == -1) {
		return errors.New("Use -id argument and one or many of -name, -sort arguments")
	}

	// update name
	if *name != "" {
		_, err = db.DB.Exec(`UPDATE categories SET categoryname = $1 WHERE id = $2`, *name, *id)
		if err != nil {
			return err
		}
	}

	// update sort
	if *sort != -1 {
		_, err = db.DB.Exec(`UPDATE categories SET sort = $1 WHERE id = $2`, *sort, *id)
		if err != nil {
			return err
		}
	}

	// print success and return
	fmt.Println("Category updated")
	return nil
}
