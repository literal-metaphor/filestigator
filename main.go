package main

import (
	FA "literal-metaphor/filestigator/analysis/filename"
	"os"
)

func main() {
	var filenames []string = make([]string, 0);

	filenames = append(filenames, os.Args[1:]...)

	for _, filename := range filenames {
		FA.TestGeneric(filename);
	}
}