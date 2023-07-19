package main

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
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

func syncFile(sourcePath string, targetPath string) {
	fmt.Println(sourcePath, targetPath)
	err := exec.Command("mkdir", "-p", filepath.Dir(targetPath)).Run()
	if err != nil {
		panic("Failed to create folder")
	}
	err = exec.Command("cp", sourcePath, targetPath).Run()
	if err != nil {
		panic("Failed to write file")
	}
}

func isImage(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))[1:]
	allExtensions := []string{"ase", "art", "bmp", "blp", "cd5", "cit", "cpt", "cr2", "cut", "dds", "dib", "djvu",
		"egt", "exif", "gif", "gpl", "grf", "icns", "ico", "iff", "jng", "jpeg", "jpg", "jfif", "jp2", "jps", "lbm",
		"max", "miff", "mng", "msp", "nef", "nitf", "ota", "pbm", "pc1", "pc2", "pc3", "pcf", "pcx", "pdn", "pgm", "PI1",
		"PI2", "PI3", "pict", "pct", "pnm", "pns", "ppm", "psb", "psd", "pdd", "psp", "px", "pxm", "pxr", "qfx", "raw",
		"rle", "sct", "sgi", "rgb", "int", "bw", "tga", "tiff", "tif", "vtf", "xbm", "xcf", "xpm", "3dv", "amf", "ai",
		"awg", "cgm", "cdr", "cmx", "dxf", "e2d", "egt", "eps", "fs", "gbr", "odg", "svg", "stl", "vrml", "x3d", "sxd",
		"v2d", "vnd", "wmf", "emf", "art", "xar", "png", "webp", "jxr", "hdp", "wdp", "cur", "ecw", "iff", "lbm", "liff",
		"nrrd", "pam", "pcx", "pgf", "sgi", "rgb", "rgba", "bw", "int", "inta", "sid", "ras", "sun", "tga", "heic", "heif"}
	for _, ext := range allExtensions {
		if ext == extension {
			return true
		}
	}
	return false
}

func buildFilename(source string, allPaths []string) string {
	usePlaceholder := false
	createdString, err := exec.Command("stat", source, "-c", "%W").Output()
	if err != nil {
		createdString = nil
	}
	unix, err := strconv.Atoi(string(createdString))
	if unix == 0 || err != nil {
		modifiedString, err := exec.Command("stat", source, "-c", "%Z").Output()
		if err != nil {
			modifiedString = nil
		}
		unix, err = strconv.Atoi(string(modifiedString))
		if err != nil {
			usePlaceholder = true
		}
	}

	date := "YYYY_DD_MM"
	if !usePlaceholder {
		date = time.Unix(int64(unix), 0).Format(date)
	}
	return fmt.Sprintf("%s_%d%s", date, getMaxCount(allPaths, date)+1, filepath.Ext(source))
}

func getMaxCount(allPaths []string, date string) int {
	out := 0
	for _, path := range allPaths {
		base := filepath.Base(path)
		pathDate := base[0:10]
		if pathDate == date {
			pathNumber := base[11:]
			newNum, err := strconv.Atoi(strings.Split(pathNumber, ".")[0])
			if err == nil && out < newNum {
				out = newNum
			}
		}
	}
	return out
}
