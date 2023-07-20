package main

import (
	"path/filepath"
)

func checkTargetNames(basePath string, paths []string, hashes [][]string, isDryRun bool, verbose bool) ([]string, [][]string) {
	for index, path := range paths {
		if isImage(path) {
			filename := filepath.Base(path)
			_, vaildName := getStandardName(filename)
			if !vaildName {
				newName := buildFilename(path, paths)
				newPath := filepath.Join(basePath, newName)
				if filepath.Dir(path) != basePath {
					syncFile(path, newPath, isDryRun, verbose)
				} else {
					renameFile(path, newPath, isDryRun, verbose)
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

func syncBlock(block Block, targetPaths []string, targetHashes [][]string, isDryRun bool, verbose bool) {
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
					syncFile(sourcePath, newPath, isDryRun, verbose)
					targetHashes = append(targetHashes, []string{newPath, sourceHash})
					targetPaths = append(targetPaths, newPath)
				}
			}
		}
	}
}

func main() {
	args := cli()
	config := getConfig(args.ConfigFile)
	for _, block := range config.Blocks {
		targetPaths := getPaths(block.Target)
		targetHashes := getHashes(targetPaths)
		targetPaths, targetHashes = checkTargetNames(block.Target, targetPaths, targetHashes, args.IsDryRun, args.Verbose)
		syncBlock(block, targetPaths, targetHashes, args.IsDryRun, args.Verbose)
	}
}
