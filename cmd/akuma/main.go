package main

import (
	"net/http"

	"github.com/luitel777/akuma/internal/interface/sqlite"
	"github.com/luitel777/akuma/internal/server"
)

func main() {
	sqlite.ClearHashmap()
	sqlite.GenerateUniqueIdenfitiers()
	http.ListenAndServe(":3333", server.SetupRoutes())
}
