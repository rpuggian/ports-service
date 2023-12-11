package handlers

import (
	"errors"
	"github.com/rpuggian/ports-service/api/api"
	"github.com/rpuggian/ports-service/api/domain/contracts"
	"github.com/rpuggian/ports-service/api/domain/contracts/mocks"
	"github.com/rpuggian/ports-service/api/domain/handlers/request"
	"github.com/rpuggian/ports-service/server/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestServer_decodeFile(t *testing.T) {
	type args struct {
		fileStream io.ReadCloser
	}
	tests := []struct {
		name    string
		args    args
		want    request.UploadPortByFileRequest
		wantErr bool
	}{
		{
			name: "should decode file with success",
			args: args{
				fileStream: io.NopCloser(strings.NewReader(`
					{
					  "A": {
						"name": "nameA",
						"city": "cityA",
						"country": "countryA",
						"alias": [],
						"regions": [],
						"coordinates": [
						  1.1,
						  2.2
						],
						"province": "provinceA",
						"timezone": "timezoneA",
						"unlocs": [
						  "A"
						],
						"code": "1"
					  },
					  "B": {
						"name": "nameB",
						"coordinates": [
						  2.2,
						  3.3
						],
						"city": "cityB",
						"province": "provinceB",
						"country": "countryB",
						"alias": [],
						"regions": [],
						"timezone": "timezoneB",
						"unlocs": [
						  "B"
						],
						"code": "2"
					  }
					}
					`)),
			},
			want: request.UploadPortByFileRequest{
				"A": request.CreatePortRequest{
					Name:        "nameA",
					City:        "cityA",
					Province:    "provinceA",
					Country:     "countryA",
					Alias:       []string{},
					Regions:     []string{},
					Coordinates: []float64{1.1, 2.2},
					Timezone:    "timezoneA",
					Unlocs:      []string{"A"},
					Code:        "1",
				},
				"B": request.CreatePortRequest{
					Name:        "nameB",
					City:        "cityB",
					Province:    "provinceB",
					Country:     "countryB",
					Alias:       []string{},
					Regions:     []string{},
					Coordinates: []float64{2.2, 3.3},
					Timezone:    "timezoneB",
					Unlocs:      []string{"B"},
					Code:        "2",
				},
			},
			wantErr: false,
		},
		{
			name: "should return error when body is a invalid json",
			args: args{
				fileStream: io.NopCloser(strings.NewReader(`invalid_json_body`)),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &portHandler{}
			got, err := h.sendFiles(tt.args.fileStream)
			if (err != nil) != tt.wantErr {
				t.Errorf("sendFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sendFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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
						Return(&entity.Port{ID: "1234"}, nil)
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
			req, _ := http.NewRequest(http.MethodGet, "/port/1234", nil)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}

func TestServer_UploadPortFileHandler(t *testing.T) {
	type fields struct {
		portService contracts.PortService
	}
	tests := []struct {
		name         string
		fields       fields
		body         string
		expectedCode int
	}{
		{
			name: "should return 201 when decode and stream create ports runs with success",
			fields: fields{
				portService: func() contracts.PortService {
					portService := &mocks.PortService{}
					portService.On("StreamCreate", mock.Anything, request.UploadPortByFileRequest{
						"MOCKID": request.CreatePortRequest{
							Name:        "MOCKNAME",
							City:        "MOCKCITY",
							Country:     "MOCKCOUNTRY",
							Alias:       []string{},
							Regions:     []string{},
							Coordinates: []float64{1.1, 2.2},
							Province:    "MOCKPROVINCE",
							Timezone:    "MOCKTIMEZONE",
							Unlocs:      []string{"MOCKID"},
							Code:        "1",
						},
					}).Return(nil)
					return portService
				}(),
			},
			body: `	{
					  "MOCKID": {
						"name": "MOCKNAME",
						"city": "MOCKCITY",
						"country": "MOCKCOUNTRY",
						"alias": [],
						"regions": [],
						"coordinates": [
						  1.1,
						  2.2
						],
						"province": "MOCKPROVINCE",
						"timezone": "MOCKTIMEZONE",
						"unlocs": [
						  "MOCKID"
						],
						"code": "1"
					  }
					}`,
			expectedCode: http.StatusCreated,
		},
		{
			name: "should return 500 when decode fails",
			fields: fields{
				portService: func() contracts.PortService {
					return nil
				}(),
			},
			body:         `invalid_json`,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name: "should return 500 when StreamCreate function fails",
			fields: fields{
				portService: func() contracts.PortService {
					portService := &mocks.PortService{}
					portService.On("StreamCreate", mock.Anything, request.UploadPortByFileRequest{
						"MOCKID": request.CreatePortRequest{Name: "MOCKNAME"},
					}).Return(errors.New("stream_create_error"))
					return portService
				}(),
			},
			body: `{
					  "MOCKID": {
						"name": "MOCKNAME",
					  }
                  }`,
			expectedCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := &portHandler{
				portService: tt.fields.portService,
			}

			router := api.NewMuxRouter(handler)
			req, _ := http.NewRequest(
				http.MethodPost,
				"/port",
				io.NopCloser(strings.NewReader(tt.body)),
			)
			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}
