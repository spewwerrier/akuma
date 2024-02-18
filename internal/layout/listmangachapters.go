package layout

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/luitel777/akuma/internal/interface/akuma"
	"github.com/luitel777/akuma/internal/interface/sqlite"
)

type MangaInformation struct {
	Title    string
	Hash     string
	Chapters []string
	URL      []int
}

func ListMangaChapters(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")
	fmt.Println(hash)

	mangaName := sqlite.RetriveManga(hash, sqlite.NAME)
	entries := akuma.GetMangaChapters(mangaName)

	mangaInformation := MangaInformation{}
	mangaInformation.Title = mangaName
	mangaInformation.Hash = hash

	for i, volumes := range entries {
		mangaInformation.Chapters = append(mangaInformation.Chapters, volumes.Name())
		mangaInformation.URL = append(mangaInformation.URL, i)
	}

	// tmpl := template.Must(template.ParseFiles("web/listMangaChapters.html"))

	tmpl := template.Must(template.ParseFS(akuma.Content, "assets/web/listMangaChapters.html"))
	err := tmpl.Execute(w, mangaInformation)
	if err != nil {
		log.Println(err)
	}
}
