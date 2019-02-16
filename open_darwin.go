// +build darwin

package main

import (
	"os/exec"
)

func open(path string) error {
	return exec.Command("open", path).Run()
}
