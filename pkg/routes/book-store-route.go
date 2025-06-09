package routes

import (
	"github.com/gorilla/mux"
	"github.com/kushal88053/Go_PROJECT_2/pkg/controllers"
)

var RegirsterRoutes = func(router *mux.Router) {

	// router.HandleFunc("/api/v1/health", controllers.HealthCheck).Methods("GET")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBooksById).Methods("GET")
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

}
