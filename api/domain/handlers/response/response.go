package response

import proto "github.com/rpuggian/ports-service/proto"

type FindPortResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	Timezone    string    `json:"timezone"`
	Unlocs      []string  `json:"unlocs"`
	Code        string    `json:"code"`
}

func NewFindPortResponse(port *proto.Port) *FindPortResponse {
	return &FindPortResponse{
		ID:          port.Id,
		Name:        port.Name,
		City:        port.City,
		Province:    port.Province,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}
}

type UploadPortByFileResponse struct {
	Total int32 `json:"total"`
}

func NewUploadPortByFileResponse(response *proto.CreateResponse) *UploadPortByFileResponse {
	return &UploadPortByFileResponse{
		Total: response.Total,
	}
}
