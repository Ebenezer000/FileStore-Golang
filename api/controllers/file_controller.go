package controllers

import (
	"encoding/json"
	"filestore-golang/services"
	"net/http"

	"github.com/gorilla/mux"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Call file service to handle upload
	fileID, err := services.UploadFileService(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"file_id": fileID})
}

func GetUploadedFiles(w http.ResponseWriter, r *http.Request) {
	files, err := services.GetFilesService()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(files)
}

func DownloadFileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["id"]

	err := services.DownloadFileService(fileID, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
