package application

import (
	"os"
	"path/filepath"
)

var Workdir string = "."

func initWorkdir() error {
	var absPath string
	var err error

	if Workdir == "." {
		absPath, err = os.Getwd()
	} else {
		absPath, err = filepath.Abs(Workdir)
	}

	if err != nil {
		return err
	}

	Workdir = absPath
	return nil
}
