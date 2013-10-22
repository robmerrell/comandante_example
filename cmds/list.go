package cmds

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var ListDoc = `
Usage: list [directory]

list all of the files either in the current directory, or the target directory
given as the first option.
`

func ListAction() error {
	dirToList, err := os.Getwd()
	if err != nil {
		return err
	}

	// get the first argument and treat it as the directory to list
	args := os.Args
	if len(args) >= 3 {
		dirToList, err = filepath.Abs(args[2])
		if err != nil {
			return err
		}
	}

	// print out a list of files in the directory
	files, err := ioutil.ReadDir(dirToList)
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}
