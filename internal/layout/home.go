package layout

import (
	"fmt"
	"html/template"
	"net/http"
)

type Information struct {
	Name        string
	Description string
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	info := Information{
		Name:        "Akuma",
		Description: "Read locally archived manga archives",
	}

	tmpl := template.Must(template.ParseFiles("web/home.html"))

	err := tmpl.Execute(w, info)
	if err != nil {
		fmt.Print(err)
	}
}
