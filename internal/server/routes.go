package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luitel777/akuma/internal/interface/akuma"
	"github.com/luitel777/akuma/internal/interface/sqlite"
	"github.com/luitel777/akuma/internal/layout"
)

func initAkuma(w http.ResponseWriter, r *http.Request) {
	sqlite.ClearHashmap()
	sqlite.GenerateUniqueIdenfitiers()
}

func SetupRoutes() chi.Router {
	routes := chi.NewRouter()
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)

	routes.Mount("/debug", middleware.Profiler())

	routes.Get("/", layout.Homepage)
	routes.Get("/manga", layout.ListManga)
	routes.Get("/init", initAkuma)
	routes.Get("/manga/{hash}", layout.ListMangaChapters)
	routes.Get("/manga/{hash}/{id}", ExtractServeManga)

	akuma.VerifyEmbed()
	routes.Handle("/*", http.FileServerFS(akuma.Content))
	fmt.Println("finished setting up routes")
	fmt.Println("serving on port 3333")

	return routes
}
