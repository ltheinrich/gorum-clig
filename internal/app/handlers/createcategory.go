package handlers

import (
	"errors"
	"fmt"

	"github.com/ltheinrich/clig/internal/pkg/db"
)

// CreateCategory handler
func CreateCategory() (err error) {
	// check variables
	if *name == "" || *sort == -1 {
		return errors.New("Use -name, -sort arguments")
	}

	// insert into database
	_, err = db.DB.Exec(`INSERT INTO categories (categoryname, sort) VALUES ($1, $2);`,
		*name, *sort)
	if err != nil {
		return err
	}

	// print success and return
	fmt.Println("Category created")
	return nil
}
