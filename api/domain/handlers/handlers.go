package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rpuggian/ports-service/api/domain/contracts"
	"log"
	"net/http"
)

type portHandler struct {
	portService contracts.PortService
}

func NewPortHandler(portService contracts.PortService) contracts.PortHandler {
	return &portHandler{
		portService: portService,
	}
}

func (s *portHandler) GetFileByID(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	portID := mux.Vars(request)["id"]

	port, err := s.portService.FindByID(ctx, portID)
	if err != nil {
		log.Println("error when get port by id: ", err)
		http.Error(response, "error when get port by id", http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(port)
	if err != nil {
		log.Println("error when marshal response: ", err)
		http.Error(response, "error when marshal response", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	_, err = response.Write(bytes)
	if err != nil {
		log.Println("error when write response: ", err)
		http.Error(response, "error when write response", http.StatusInternalServerError)
		return
	}
}

func (s *portHandler) UploadPortFileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//1MB max file size
	if err := r.ParseMultipartForm(1024 * 1024); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createResponse, err := s.portService.UploadPortFile(ctx, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(createResponse)
	if err != nil {
		log.Println("error when marshal response: ", err)
		http.Error(w, "error when marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)
	if err != nil {
		log.Println("error when write response: ", err)
		http.Error(w, "error when write response", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
