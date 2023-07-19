package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func outputHashes(hashes [][]string, path string) {
	fp, err := os.Create(path)
	if err != nil {
		fmt.Printf("Warning: could not open file for writing hash list %s.\n", path)
	}
	w := csv.NewWriter(fp)

	for _, hash := range hashes {
		if err := w.Write(hash); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func syncBlock(block Block) {
	targetPaths := getPaths(block.Target)
	targetHashes := getHashes(targetPaths)
	for _, sourceRoot := range block.Sources {
		sourcePaths := getPaths(sourceRoot)
		sourceHashes := getHashes(sourcePaths)
		for _, sourceFile := range sourceHashes {
			sourcePath := sourceFile[0]
			if !isImage(sourcePath) {
				continue
			}
			sourceHash := sourceFile[1]
			_, _, found := findPathHash(targetHashes, sourceHash, false)
			if !found {
				targetPath := buildFilename(sourcePath, targetPaths)
				syncFile(sourcePath, targetPath)
				targetHashes = append(targetHashes, []string{targetPath, sourceHash})
			}
		}
	}
}

func main() {
	configPath := "/mnt/ImpSSD/Development/dupliphoto/test.yml"
	config := getConfig(configPath)
	for _, block := range config.Blocks {
		syncBlock(block)
	}
}
