package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	fa "filestigator/analysis/filename"

	"github.com/fatih/color"
)

func main() {
	// Declare flags
	flag_dir := flag.String("dir", "", "Read all files in the specified directory")
	flag_dir_recursive := flag.Bool("r", false, "Recursively read all files, must be used in pair with -dir")
	flag_files := flag.String("files", "", "Accept multiple input of files to be read, separated by comma (,) without space")

	// Parse flags
	flag.Parse()

	// * Flag edge cases
	// No essential flag specified
	if *flag_dir == "" && *flag_files == "" {
		color.Yellow("Exiting: No directory or file to be read specified")
		return
	}
	// -dir with -files
	if *flag_dir != "" && *flag_files != "" {
		color.Yellow("Exiting: -dir and -files cannot be used at the same time")
		return
	}
	// -r without -dir
	if *flag_dir == "" && *flag_dir_recursive {
		color.Yellow("Exiting: Recursive flag (-r) can only be used with -dir")
		return
	}
	// ? are there more?

	var filenames []string = make([]string, 0) // Keep track of filenames
	var unreadables []string = make([]string, 0) // Keep track on file read errors

	// Read on specified directory
	// * Issues:
	// 1. When getting files recursively, unreadable files are appended to unreadables, but with non-recursive option any unreadable will result in panic
	if *flag_dir != "" {

		if *flag_dir_recursive {
			// Recursively get files
			filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
				// Append unreadable files to unreadables
				if err != nil {
					unreadables = append(unreadables, path)
					return nil
				}

				// Get the filename
				if !info.IsDir() {
					filenames = append(filenames, filepath.Base(path))
				}

				return nil
			})

		} else {
			// Unrecursively get files
			files, err := os.ReadDir(*flag_dir)
			if err != nil {
				panic(err)
			}

			// Append filename
			for _, file := range files {
				if (!file.IsDir()) {
					filenames = append(filenames, file.Name())
				}
			}
		}
	}

	// Read specified files
	if *flag_files != "" {
		// TODO: Find a way to allow space between values
		filenamesList := strings.Split(*flag_files, ",")
		for _, filename := range filenamesList {
			// Check if filename exists
			file, err := os.Open(filename)
			if err != nil {
				unreadables = append(unreadables, filename)
			} else {
				filenames = append(filenames, filename)
			}

			// Close the file
			file.Close()
		}
	}

	// Print all unreadables
	for _, ur := range unreadables {
		color.Red("Cannot read from " + ur)
	}

	// Process each files
	for _, filename := range filenames {
		fa.Check(filename)
	}
}