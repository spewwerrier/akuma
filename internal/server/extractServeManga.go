package server

import (
	"archive/zip"
	"encoding/base64"
	"fmt"
	"html/template"

	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/luitel777/akuma/internal/interface/akuma"
	"github.com/luitel777/akuma/internal/interface/sqlite"
)

func unzipCBZ(filePath string, fileName string, pagechan chan string) {
	reader, err := zip.OpenReader(filePath)
	if err != nil {
		log.Println("unzip cbz", err)
		pagechan <- ""
	}
	defer reader.Close()

	for _, file := range reader.File {
		f, err := file.Open()
		if err != nil {
			log.Println("reading file", err)
			continue
		}
		defer f.Close()

		buf, err := io.ReadAll(f)
		if err != nil {
			log.Println("reading buffer: ", err)
			continue
		}

		if buf[0] == 255 && buf[1] == 216 || buf[0] == 137 && buf[1] == 80 || buf[0] == 82 && buf[1] == 73 {
			// only checks 2 bytes but this filter outs a lot of non jpeg/png files
			encoded := base64.StdEncoding.EncodeToString(buf)
			pagechan <- encoded
		}
	}
	close(pagechan)
}

type Property struct {
	Title string
	Next  int
	Prev  int
}

// /manga/{hash}/{id}
func ExtractServeManga(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")
	volume := chi.URLParam(r, "id")

	vol, err := strconv.Atoi(volume)
	if err != nil {
		http.Error(w, "cannot convert URL from string to integer", http.StatusBadRequest)
		return
	}

	mangaName := sqlite.RetriveManga(hash, sqlite.NAME)
	entries := akuma.GetMangaChapters(mangaName)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if vol >= len(entries) {
		http.Error(w, "invalid volume ID", http.StatusBadRequest)
		return
	}

	fmt.Println(entries[vol].Name())

	filePath := "manga/" + mangaName + "/" + entries[vol].Name()

	pageChan := make(chan string)

	go unzipCBZ(filePath, entries[vol].Name(), pageChan)

	data := Property{
		Title: mangaName + string(rune(vol+1)),
	}
	data.Next, data.Prev = akuma.DoesNextChapterExists(vol, entries)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	for {
		page := <-pageChan
		if page == "" {
			break
		}
		time.Sleep(100)

		_, err := fmt.Fprintf(w, "<img src=\"data:image/png;base64,%s\">\n", page)
		if err != nil {
			log.Println("error writing", err)
			break
		}
	}
	tmpl := template.Must(template.ParseFiles("web/reader.html"))
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
