package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	verbose = flag.Bool("verbose", false, "verbose")
)

func mkworkdir() error {
	dir, err := NewWorkDir()
	if err != nil {
		return err
	}

	if *verbose {
		log.Printf("workdir: %s\n", dir.Path)
	}

	if err := dir.Create(); err != nil {
		return err
	}

	if dir.IsOpen() {
		if *verbose {
			log.Printf("open %s\n", dir.Path)
		}
		if err := dir.Open(); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()

	if err := mkworkdir(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}
}
