package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func main() {
	v, err := exec.Command("git", "describe", "--tags").CombinedOutput()
	if err != nil {
		log.Fatal("Running git:", err)
	}
	version := strings.TrimSpace(string(v))
	if err := ioutil.WriteFile("version.go",
		[]byte(fmt.Sprintf(`package main

import (
	"flag"
	"fmt"
	"os"
)

const VERSION = "%s"

var versionFlag = flag.Bool("version", false, "prints version number")

func printVersion() {
	flag.Parse()
	if *versionFlag {
		fmt.Println(VERSION)
		os.Exit(0)
	}
}
`, string(version))),
		0664); err != nil {
		fmt.Println("ERROR:", err)
	}
}
