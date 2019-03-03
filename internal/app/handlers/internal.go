package handlers

import "flag"

// Handler function
type Handler func() (err error)

var (
	// HandlersMap name -> handler
	HandlersMap = map[string]Handler{
		"v":  VersionHandler,
		"cb": CreateBoard,
		"cc": CreateCategory,
		"db": DeleteBoard,
		"dc": DeleteCategory,
		"dp": DeletePost,
		"dt": DeleteThread,
		"du": DeleteUser,
		"eb": EditBoard,
		"ec": EditCategory,
		"ep": EditPost,
		"et": EditThread,
		"eu": EditUser,
		"lb": ListBoards,
		"lc": ListCategories,
		"lp": ListPosts,
		"lt": ListThreads,
		"lu": ListUsers,
	}

	// int variables
	id       = flag.Int("id", -1, "ID (Board, Category, User, Post)")
	sort     = flag.Int("sort", -1, "Sort ID (Board, Category)")
	board    = flag.Int("board", -1, "Board ID (Thread)")
	thread   = flag.Int("thread", -1, "Thread ID (Post)")
	author   = flag.Int("author", -1, "Author ID (Post, Thread)")
	created  = flag.Int("created", -1, "Created Timestamp (Post, Thread)")
	category = flag.Int("category", -1, "Category ID (Board)")

	// string variables
	icon        = flag.String("icon", "forum", "Icon (Board)")
	description = flag.String("description", "", "Description (Board)")
	password    = flag.String("password", "", "Password (User)")
	content     = flag.String("content", "", "Content (Post, Thread)")
	name        = flag.String("name", "", "Name (Board, Category, Thread)")
)
