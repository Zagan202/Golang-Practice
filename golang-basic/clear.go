package main

import (
	"os"
	"os/exec"
)

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
