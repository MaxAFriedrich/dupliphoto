package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/djherbis/times"
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

func syncFile(sourcePath string, targetPath string, isDryRun bool, verbose bool) {
	if verbose {
		fmt.Printf("copy: %s -> %s\n", sourcePath, targetPath)
	}
	if isDryRun {
		return
	}
	err := exec.Command("mkdir", "-p", filepath.Dir(targetPath)).Run()
	if err != nil {
		panic("Failed to create folder:" + err.Error())
	}
	err = exec.Command("cp", sourcePath, targetPath).Run()
	if err != nil {
		panic("Failed to write file:" + err.Error())
	}
}

func renameFile(oldPath string, newPath string, isDryRun bool, verbose bool) {
	if verbose {
		fmt.Printf("rename: %s -> %s\n", oldPath, newPath)
	}
	if isDryRun {
		return
	}
	err := os.Rename(oldPath, newPath)
	if err != nil {
		panic("Failed to rename file: " + err.Error())
	}
}

func isImage(path string) bool {
	if len(filepath.Ext(path)) < 1 {
		return false
	}
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

func getDate(path string) standardName {
	var out standardName
	out.index = 0
	out.day = 0
	out.month = 0
	out.year = 0
	out.fullDate = "YYYY_MM_DD"

	t, err := times.Stat(path)
	if err != nil {
		return out
	}

	var birthTime time.Time
	if t.HasBirthTime() {
		birthTime = t.BirthTime()
	} else if t.HasChangeTime() {
		birthTime = t.ChangeTime()
	} else {
		return out
	}

	out.fullDate = birthTime.Format("2006_01_02")
	out.day = birthTime.Day()
	out.month = int(birthTime.Month())
	out.year = birthTime.Year()

	return out
}
