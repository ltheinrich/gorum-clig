package main

import (
	"github.com/ltheinrich/clig/cmd"
)

func main() {
	if err := cmd.Init(); err != nil {
		panic(err)
	}
}
