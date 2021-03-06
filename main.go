package main

import (
	"log"
	"net/http"

	"github.com/alx-t/cdr-service/routes"
	"github.com/codegangsta/negroni"
)

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	router := routers.InitRoutes()

	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}

	// defer common.Db.Close()
	// common.ExecuteCmd("ls -al")

	log.Println("Listening...")
	server.ListenAndServe()
}
