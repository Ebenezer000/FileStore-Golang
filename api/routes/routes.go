package routes

import (
	"filestore-golang/api/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// File routes
	router.HandleFunc("/upload", controllers.UploadFile).Methods("POST")
	router.HandleFunc("/get_files", controllers.GetUploadedFiles).Methods("GET")
	router.HandleFunc("/download/{id}", controllers.DownloadFileByID).Methods("GET")

	return router
}
