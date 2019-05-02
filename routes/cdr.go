package routers

import (
	"github.com/alx-t/cdr-service/controllers"
	"github.com/gorilla/mux"
)

func SetCDRRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", controllers.Root).Methods("GET")

	router.HandleFunc("/cdr", controllers.GetCDRJson).Methods("GET")
	// router.HandleFunc("/projects", controllers.CreateProject).Methods("POST")
	// router.HandleFunc("/projects/{id}", controllers.GetProjectById).Methods("GET")
	// router.HandleFunc("/projects/{id}", controllers.UpdateProject).Methods("PATCH")
	// router.HandleFunc("/projects/{id}", controllers.DeleteProject).Methods("DELETE")

	return router
}
