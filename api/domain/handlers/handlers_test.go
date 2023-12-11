package handlers

import (
	"errors"
	"github.com/rpuggian/ports-service/api/api"
	"github.com/rpuggian/ports-service/api/domain/contracts"
	"github.com/rpuggian/ports-service/api/domain/contracts/mocks"
	"github.com/rpuggian/ports-service/api/domain/handlers/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_GetFileByID(t *testing.T) {
	type fields struct {
		portService contracts.PortService
	}

	tests := []struct {
		name         string
		fields       fields
		expectedCode int
	}{
		{
			name: "should return 200 when get file by id runs with success",
			fields: fields{
				portService: func() contracts.PortService {
					portService := &mocks.PortService{}
					portService.On("FindByID", mock.Anything, "1234").
						Return(&response.FindPortResponse{ID: "1234"}, nil)
					return portService
				}(),
			},
			expectedCode: http.StatusOK,
		},
		{
			name: "should return 500 when get file by id fails",
			fields: fields{
				portService: func() contracts.PortService {
					portService := &mocks.PortService{}
					portService.On("FindByID", mock.Anything, "1234").
						Return(nil, errors.New("find_error"))
					return portService
				}(),
			},
			expectedCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &portHandler{
				portService: tt.fields.portService,
			}

			router := api.NewMuxRouter(handler)
			req, _ := http.NewRequest(http.MethodGet, "/ports/1234", nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}
