package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	// EnvWorkDir is root working directory path if it sets.
	EnvWorkDir = "WORKDIR"
	// EnvWorkDirOpen is for opening the working directory if it sets.
	EnvWorkDirOpen = "WORKDIR_OPEN"
)

// A WorkDir is working directory.
type WorkDir struct {
	Path string
}

func getWorkDir() (path string) {
	var keys = []string{EnvWorkDir, "TEMP", "TMP"}
	for _, key := range keys {
		path = os.Getenv(key)
		if path != "" {
			return
		}
	}
	return
}

// NewWorkDir returns WorDir that creates and opens working directory.
func NewWorkDir() (*WorkDir, error) {
	dir := getWorkDir()
	if dir == "" {
		return nil, errors.New("Fail to retrieve value of environment variable 'WORKDIR'")
	}

	t := time.Now()
	daily := fmt.Sprintf("%d%c%02d%c%02d", t.Year(), filepath.Separator, t.Month(), filepath.Separator, t.Day())
	path := filepath.Join(dir, daily)

	return &WorkDir{Path: path}, nil
}

func (w *WorkDir) Create() error {
	return os.MkdirAll(w.Path, 0777)
}

func (w *WorkDir) Open() error {
	if !w.IsOpen() {
		return nil
	}
	return open(w.Path)
}

// IsOpen returns true if EnvWorkDirOpen is set in environment variables.
func (w *WorkDir) IsOpen() bool {
	return "" != os.Getenv(EnvWorkDirOpen)
}
