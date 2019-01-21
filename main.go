package main

import (
	"fmt"

	"github.com/ltheinrich/clig/cmd"
)

func main() {
	if err := cmd.Init(); err != nil {
		fmt.Println(err)
	}
}
