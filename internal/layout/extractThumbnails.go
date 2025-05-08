package layout

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	globals "github.com/luitel777/akuma/internal"
	"github.com/luitel777/akuma/internal/interface/sqlite"
)

func extractThumbnail(hash string) string {
	if !check_on_cache(hash) {
		generate_cache(hash)
	}
	return retriveFromCache(hash)
}

/*
get filehash
unzip the archive
cover is usually the 1st page (safest bet)
put the cover in hash/id on the cache folder
*/

func check_on_cache(hash string) bool {
	_, err := os.ReadFile(globals.HOME + "/.cache/akumathumbnails/" + hash)
	if err != nil {
		fmt.Println("generating thumbnails for ", sqlite.RetriveManga(hash, sqlite.NAME))
		return false
	}
	return true
}

func generate_cache(hash string) {
	dir := sqlite.RetriveManga(hash, sqlite.NAME)
	entries, err := os.ReadDir(globals.DIRECTORY + dir)
	if err != nil {
		log.Println(err)
	}
	entry := globals.DIRECTORY + dir + "/" + entries[0].Name()

	reader, err := zip.OpenReader(entry)
	if err != nil {
		log.Println("cannot open file", err)
	}
	defer reader.Close()

	file, err := reader.File[0].Open()
	if err != nil {
		log.Println("cannot open file of zip archive", err)
	}
	defer file.Close()

	buf, err := io.ReadAll(file)
	if err != nil {
		log.Println("error reading file contents")
	}
	// encoded := base64.StdEncoding.EncodeToString(buf)

	err = os.MkdirAll(globals.HOME+"/.cache/akumathumbnails", os.ModePerm)
	if err != nil {
		log.Println("error:", err)
	}
	err = os.WriteFile(globals.HOME+"/.cache/akumathumbnails/"+hash, buf, 0644)
	if err != nil {
		log.Println("error: ", err)
	}

}

func retriveFromCache(hash string) string {
	file, err := os.ReadFile(globals.HOME + "/.cache/akumathumbnails/" + hash)
	if err != nil {
		log.Println("cannot read thumbnail", err)
	}
	encoded := base64.StdEncoding.EncodeToString(file)
	return encoded
}
