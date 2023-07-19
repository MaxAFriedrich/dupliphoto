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

func main() {
	base := "/mnt/ImpSSD/Development/dupliphoto"
	target := base + "/imgGen/deep"
	hashFile := base + "/targetHash"
	allPaths := getPaths(target)
	hashed := getHashes(allPaths)
	outputHashes(hashed, hashFile)
}
