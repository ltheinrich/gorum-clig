package cmd

import (
	"errors"
	"flag"

	"github.com/ltheinrich/clig/internal/app/handlers"
	"github.com/ltheinrich/clig/pkg/db"
)

var (
	// database flags
	host     = flag.String("host", "::1", "PostgreSQL database host")
	port     = flag.String("port", "5432", "PostgreSQL database port")
	ssl      = flag.String("ssl", "disable", "PostgreSQL database ssl mode")
	database = flag.String("database", "gorum", "PostgreSQL database name")
	username = flag.String("username", "gorum", "PostgreSQL database username")
	password = flag.String("password", "gorum", "PostgreSQL database password")

	// Actions to handle
	Actions = map[string]*bool{
		"cb": flag.Bool("cb", false, "Create a new board"),
		"cc": flag.Bool("cc", false, "Create a new category"),
		"db": flag.Bool("db", false, "Delete a board"),
		"dc": flag.Bool("dc", false, "Delete a category"),
		"du": flag.Bool("du", false, "Delete a user"),
		"dp": flag.Bool("dp", false, "Delete a post"),
		"dt": flag.Bool("dt", false, "Delete a thread"),
		"eb": flag.Bool("eb", false, "Edit a board"),
		"ec": flag.Bool("ec", false, "Edit a category"),
		"eu": flag.Bool("eu", false, "Edit a user"),
		"ep": flag.Bool("ep", false, "Edit a post"),
		"et": flag.Bool("et", false, "Edit a thread"),
		"lb": flag.Bool("lb", false, "List boards"),
		"lc": flag.Bool("lc", false, "List categories"),
		"lu": flag.Bool("lu", false, "List users"),
		"lp": flag.Bool("lp", false, "List posts"),
		"lt": flag.Bool("lt", false, "List threads"),
	}
)

// Init command startup
func Init() (err error) {
	flag.Parse()

	// connect to database
	if err := db.Connect(*host, *port, *ssl, *database, *username, *password); err != nil {
		return err
	}

	for action, used := range Actions {
		if *used {
			if handler, exists := handlers.HandlersMap[action]; exists {
				return handler()
			}
			return errors.New("Unknown action")
		}
	}

	return errors.New("Use the -help argument for help (e.g. ./clig -help)")
}
