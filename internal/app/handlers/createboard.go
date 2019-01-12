package handlers

import (
	"errors"
	"fmt"

	"github.com/ltheinrich/clig/pkg/db"
)

// CreateBoard handler
func CreateBoard() (err error) {
	// check variables
	if *name == "" || *description == "" || *category == -1 || *sort == -1 {
		return errors.New("Use -name, -description, -category, -sort arguments")
	}

	// insert into database
	_, err = db.DB.Exec(`INSERT INTO boards (boardname, boarddescription, boardicon, sort, category) VALUES ($1, $2, $3, $4, $5);`,
		*name, *description, *icon, *sort, *category)
	if err != nil {
		return err
	}

	fmt.Println("Board created")

	return nil
}
