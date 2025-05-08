package layout

import (
	"html/template"
	"log"
	"net/http"

	"github.com/luitel777/akuma/internal/interface/akuma"
	"github.com/luitel777/akuma/internal/interface/sqlite"
)

type MangaDirs struct {
	Name  []string
	Hash  []string
	Cover []string
}

func ListManga(w http.ResponseWriter, r *http.Request) {
	inst := MangaDirs{}

	for _, file := range akuma.GetMangaList() {
		hash := sqlite.RetriveManga(file.Name(), sqlite.HASH)

		inst.Name = append(inst.Name, file.Name())
		inst.Hash = append(inst.Hash, hash)
		inst.Cover = append(inst.Cover, extractThumbnail(hash))
	}

	tmpl := template.Must(template.ParseFS(akuma.Content, "assets/web/list.html"))
	err := tmpl.Execute(w, inst)
	if err != nil {
		log.Println("failed to render listing page: ", err)
	}
}
