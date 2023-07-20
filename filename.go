package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
)

type standardName struct {
	year     int
	month    int
	day      int
	index    int
	fullDate string
}

func getStandardName(filename string) (standardName, bool) {
	var out standardName
	re := regexp.MustCompile(`(?m)((\d{4})_(\d{2})_(\d{2}))_(\d{4})\.\w*`)
	if !re.MatchString(filename) {
		return out, false
	}
	matches := re.FindStringSubmatch(filename)
	out.fullDate = matches[1]
	// ignoring these errors because it is certain that these are ints
	out.year, _ = strconv.Atoi(matches[2])
	out.month, _ = strconv.Atoi(matches[3])
	out.day, _ = strconv.Atoi(matches[4])
	out.index, _ = strconv.Atoi(matches[5])

	return out, true
}

func getMaxIndex(paths []string, targetDate string) int {
	currentMax := 0
	for _, path := range paths {
		filename := filepath.Base(path)
		breakdown, valid := getStandardName(filename)
		if !valid {
			continue
		}
		if currentMax < breakdown.index && targetDate == breakdown.fullDate {
			currentMax = breakdown.index
		}
	}
	return currentMax
}

func buildFilename(source string, allPaths []string) string {
	date := getDate(source)
	date.index = getMaxIndex(allPaths, date.fullDate) + 1
	return fmt.Sprintf("%s_%04d%s", date.fullDate, date.index, filepath.Ext(source))
}
