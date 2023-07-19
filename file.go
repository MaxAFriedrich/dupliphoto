package main

import (
	"io/fs"
	"path/filepath"
)

func getPaths(path string) []string {
	var allPaths []string

	filepath.WalkDir(path, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			allPaths = append(allPaths, s)
		}
		return nil
	})

	return allPaths
}
