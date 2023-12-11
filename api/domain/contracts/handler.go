package contracts

import (
	"net/http"
)

type PortHandler interface {
	GetFileByID(response http.ResponseWriter, request *http.Request)
	UploadPortFileHandler(response http.ResponseWriter, request *http.Request)
}
