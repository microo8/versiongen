package main

import (
	"flag"
	"fmt"
)

const VERSION = "v1.0.0-1-g147c865"

var versionFlag = flag.Bool("version", false, "prints version number")

func init() {
	flag.Parse()
	if *versionFlag {
		fmt.Println(VERSION)
	}
}
