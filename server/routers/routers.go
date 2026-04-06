package routers

import (
	"github.com/gorilla/mux"
	"github.com/harkirath1511/mongo-api/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies/all", controllers.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movies/create", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movies/update/{id}", controllers.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movies/delete/{id}", controllers.DeleteOne).Methods("DELETE")
	router.HandleFunc("/api/movies/deleteAll", controllers.DeleteAll).Methods("DELETE")

	return router
}
