package handlers

import "flag"

// Handler function
type Handler func() (err error)

var (
	// HandlersMap name -> handler
	HandlersMap = map[string]Handler{
		"cb": CreateBoard,
		"cc": CreateCategory,
		"db": DeleteBoard,
		"dc": DeleteCategory,
		"du": DeleteUser,
		"dp": DeletePost,
		"dt": DeleteThread,
		"lb": ListBoards,
		"lc": ListCategories,
		"lu": ListUsers,
		"lp": ListPosts,
		"lt": ListThreads,
	}

	// variables
	id          = flag.Int("id", -1, "ID (Board, Category, User, Post)")
	name        = flag.String("name", "", "Name (Board, Category)")
	description = flag.String("description", "", "Board description")
	icon        = flag.String("icon", "forum", "Board icon")
	sort        = flag.Int("sort", -1, "Sort ID (Board, Category)")
	category    = flag.Int("category", -1, "Board category")
)
