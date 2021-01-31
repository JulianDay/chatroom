package config

import (
	"os"
	"path/filepath"
)

var RootDir string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		if exists(d + "/config") {
			return d
		}

		return infer(filepath.Dir(d))
	}

	RootDir = infer(cwd)
	loadServer()
	loadSensitive()
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
