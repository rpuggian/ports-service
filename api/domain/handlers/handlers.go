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

func (s *portHandler) UploadPortFileHandler(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	createResponse, err := s.portService.UploadPortFile(ctx, request.Body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(createResponse)
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
	response.WriteHeader(http.StatusCreated)
}
