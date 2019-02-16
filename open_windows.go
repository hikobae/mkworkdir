// +build windows

package main

import (
	"os/exec"
)

func open(path string) error {
	return exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", path).Run()
}
