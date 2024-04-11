package router

import (
	"github.com/Rawat-Senapi/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.Createovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/api/deleteAllMovies", controller.GetAllMovies).Methods("DELETE")

	return router

}
