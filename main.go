package main

import (
	"log"
	"os/exec"
)

func main() {
	version, err := exec.Command("git", "describe", "--tags").CombinedOutput()
	if err != nil {
		log.Fatal("Running git:", err)
	}
	log.Println("version:", string(version))
}
