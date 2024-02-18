package akuma

import (
	"embed"
	"log"
)

//go:generate rm -rf assets
//go:generate mkdir assets
//go:generate cp -r ../../../web ../../../static assets
//go:embed assets
var Content embed.FS

func VerifyEmbed() {
	_, err := Content.ReadFile("assets/web/home.html")
	if err != nil {
		log.Panic(err)
	}
	_, err = Content.ReadFile("assets/static/reader.css")
	if err != nil {
		log.Panic(err)
	}
}
