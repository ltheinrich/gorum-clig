package handlers

import (
	"fmt"
)

var (
	// Version string generated using Makefile
	Version string

	// BuildTime string generated using Makefile
	BuildTime string
)

// VersionHandler printer
func VersionHandler() (err error) {
	// print out version and build time
	fmt.Printf("clig %v (c) 2019 Lennart Heinrich\nBuild time %v\n", Version, BuildTime)
	return nil
}
