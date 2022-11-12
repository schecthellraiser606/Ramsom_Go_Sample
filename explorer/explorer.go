package explorer

import (
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

func MapFiles(dir string) []string {
	var files []string
	var root string
	var env_name string = "ProgramFiles(x86)"

	if runtime.GOOS == "windows" {
		root = os.Getenv(env_name)
	} else {
		root = os.Getenv("HOME")
	}

	if dir == "" {
		root += "/Lhasa"
	} else {
		root += dir
	}

	error := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		files = append(files, path)

		return nil
	})

	if error != nil {
		panic(error)
	}

	return files
}
