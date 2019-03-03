package cmd

import (
	"errors"
	"flag"

	"github.com/ltheinrich/clig/internal/app/handlers"
	"github.com/ltheinrich/clig/internal/pkg/db"
)

var (
	// database flags
	dbhost = flag.String("dbhost", "::1", "PostgreSQL database host")
	dbport = flag.String("dbport", "5432", "PostgreSQL database port")
	dbssl  = flag.String("dbssl", "disable", "PostgreSQL database ssl mode")
	dbname = flag.String("dbname", "gorum", "PostgreSQL database name")
	dbuser = flag.String("dbuser", "gorum", "PostgreSQL database username")
	dbpass = flag.String("dbpass", "gorum", "PostgreSQL database password")

	// Actions to handle
	Actions = map[string]*bool{
		"v":  flag.Bool("v", false, "Get version"),
		"cb": flag.Bool("cb", false, "Create a new board"),
		"cc": flag.Bool("cc", false, "Create a new category"),
		"db": flag.Bool("db", false, "Delete a board"),
		"dc": flag.Bool("dc", false, "Delete a category"),
		"dp": flag.Bool("dp", false, "Delete a post"),
		"dt": flag.Bool("dt", false, "Delete a thread"),
		"du": flag.Bool("du", false, "Delete a user"),
		"eb": flag.Bool("eb", false, "Edit a board"),
		"ec": flag.Bool("ec", false, "Edit a category"),
		"ep": flag.Bool("ep", false, "Edit a post"),
		"et": flag.Bool("et", false, "Edit a thread"),
		"eu": flag.Bool("eu", false, "Edit a user"),
		"lb": flag.Bool("lb", false, "List boards"),
		"lc": flag.Bool("lc", false, "List categories"),
		"lp": flag.Bool("lp", false, "List posts"),
		"lt": flag.Bool("lt", false, "List threads"),
		"lu": flag.Bool("lu", false, "List users"),
	}
)

// Init command startup
func Init() (err error) {
	flag.Parse()

	// connect to database
	if err := db.Connect(*dbhost, *dbport, *dbssl, *dbname, *dbuser, *dbpass); err != nil {
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
