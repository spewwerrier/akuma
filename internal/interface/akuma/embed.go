package akuma

import (
	"embed"
	"log"
)

//go:embed assets/*
var Content embed.FS

func VerifyEmbed() {
	_, err := Content.ReadFile("assets/web/home.html")
	if err != nil {
		log.Println(err)
	}
}
