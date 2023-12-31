package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func hashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func getHashes(paths []string) [][]string {
	var out [][]string

	for _, path := range paths {
		hash, err := hashFile(path)
		if err != nil {
			fmt.Printf("Warning: file %s could not be hashed.\n", path)
			continue
		}
		newData := []string{path, hash}
		out = append(out, newData)
	}
	return out
}

func findPathHash(all [][]string, find string, searchPath bool) (string, string, bool) {
	outHash := ""
	outPath := ""
	found := false
	current := ""
	for _, file := range all {
		hash := file[1]
		path := file[0]
		if searchPath {
			current = path
		} else {
			current = hash
		}
		if current == find {
			outPath = path
			outHash = hash
			found = true
			break
		}
	}
	return outPath, outHash, found
}
