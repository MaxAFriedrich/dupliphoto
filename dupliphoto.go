package main

import (
	"path/filepath"
)

func checkTargetNames(basePath string, paths []string, hashes [][]string, isDryRun bool) ([]string, [][]string) {
	for index, path := range paths {
		if isImage(path) {
			filename := filepath.Base(path)
			_, vaildName := getStandardName(filename)
			if !vaildName {
				newName := buildFilename(path, paths)
				newPath := filepath.Join(basePath, newName)
				if filepath.Dir(path) != basePath {
					syncFile(path, newPath, isDryRun)
				} else {
					renameFile(path, newPath, isDryRun)
				}
				paths[index] = newPath
				hashes[index][0] = newPath
			}
		}
	}
	return paths, hashes
}

func inTarget(targetHashes [][]string, sourceHash string) bool {
	for _, target := range targetHashes {
		targetHash := target[1]
		if targetHash == sourceHash {
			return true
		}
	}
	return false
}

func syncBlock(block Block, targetPaths []string, targetHashes [][]string, isDryRun bool) {
	for _, sources := range block.Sources {
		sourcePaths := getPaths(sources)
		sourceHashes := getHashes(sourcePaths)
		for _, source := range sourceHashes {
			sourcePath := source[0]
			sourceHash := source[1]
			if isImage(sourcePath) {
				if !inTarget(targetHashes, sourceHash) {
					newName := buildFilename(sourcePath, targetPaths)
					newPath := filepath.Join(block.Target, newName)
					syncFile(sourcePath, newPath, isDryRun)
					targetHashes = append(targetHashes, []string{newPath, sourceHash})
					targetPaths = append(targetPaths, newPath)
				}
			}
		}
	}
}

func main() {
	configPath := "/mnt/ImpSSD/Development/dupliphoto/test.yml"
	config := getConfig(configPath)
	for _, block := range config.Blocks {
		targetPaths := getPaths(block.Target)
		targetHashes := getHashes(targetPaths)
		targetPaths, targetHashes = checkTargetNames(block.Target, targetPaths, targetHashes, true)
		syncBlock(block, targetPaths, targetHashes, true)
	}
}
