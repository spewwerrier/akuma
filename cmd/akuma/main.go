package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	globals "github.com/luitel777/akuma/internal"
	"github.com/luitel777/akuma/internal/interface/sqlite"
	"github.com/luitel777/akuma/internal/server"
)

func init() {
	var err error
	globals.HOME, err = os.UserHomeDir()
	if err != nil {
		log.Fatalln("cannot get the home directory: ", err)
		return
	}

	flag.StringVar(&globals.DIRECTORY, "dir", globals.HOME+"/.local/share/akuma", "path to search for manga")
	flag.IntVar(&globals.PORT, "port", 3333, "port to run the server")
	flag.Parse()
}

func main() {
	sqlite.ClearHashmap()
	sqlite.GenerateUniqueIdenfitiers()

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(globals.PORT),
		Handler: server.SetupRoutes(),
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("closing the server", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("failed to shutdown the server: ", err)
	}
	fmt.Println("Closing the server. Bye")

}
