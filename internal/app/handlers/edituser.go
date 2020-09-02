package handlers

import (
	"errors"
	"fmt"

	"github.com/ltheinrich/gorum-clig/internal/pkg/crypter"
	"github.com/ltheinrich/gorum-clig/internal/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

// EditUser handler
func EditUser() (err error) {
	// check variables
	if (*id == -1) || (*name == "" && *password == "") {
		return errors.New("Use -id argument and one or many of -name, -password arguments")
	}

	// update name
	if *name != "" {
		_, err = db.DB.Exec(`UPDATE users SET username = $1 WHERE id = $2`, *name, *id)
		if err != nil {
			return err
		}
	}

	// update password
	if *password != "" {
		// hash password
		var encryptedPassword []byte
		hashedPassword := crypter.Hash("gorum_" + crypter.Hash(*password))
		encryptedPassword, err = bcrypt.GenerateFromPassword([]byte(hashedPassword), bcrypt.DefaultCost+1)

		// update database
		_, err = db.DB.Exec(`UPDATE users SET passwordhash = $1 WHERE id = $2`, string(encryptedPassword), *id)
		if err != nil {
			return err
		}
	}

	// print success and return
	fmt.Println("User updated")
	return nil
}
