package handlers

import (
	"errors"
	"fmt"

	"github.com/ltheinrich/gorum-clig/internal/pkg/db"
)

// EditBoard handler
func EditBoard() (err error) {
	// check variables
	if (*id == -1) || (*name == "" && *description == "" && *icon == "" && *category == -1 && *sort == -1) {
		return errors.New("Use -id argument and one or many of -name, -description, -icon, -category, -sort arguments")
	}

	// update name
	if *name != "" {
		_, err = db.DB.Exec(`UPDATE boards SET boardname = $1 WHERE id = $2`, *name, *id)
		if err != nil {
			return err
		}
	}

	// update description
	if *description != "" {
		_, err = db.DB.Exec(`UPDATE boards SET boarddescription = $1 WHERE id = $2`, *description, *id)
		if err != nil {
			return err
		}
	}

	// update icon
	if *icon != "" {
		_, err = db.DB.Exec(`UPDATE boards SET boardicon = $1 WHERE id = $2`, *icon, *id)
		if err != nil {
			return err
		}
	}

	// update category
	if *category != -1 {
		_, err = db.DB.Exec(`UPDATE boards SET category = $1 WHERE id = $2`, *category, *id)
		if err != nil {
			return err
		}
	}

	// update sort
	if *sort != -1 {
		_, err = db.DB.Exec(`UPDATE boards SET sort = $1 WHERE id = $2`, *sort, *id)
		if err != nil {
			return err
		}
	}

	// print success and return
	fmt.Println("Board updated")
	return nil
}
