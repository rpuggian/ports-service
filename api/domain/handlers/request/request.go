package request

import proto "github.com/rpuggian/ports-service/proto"

type CreatePortRequest struct {
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

func ParseUploadPortRequestToProtoPortArray(key string, portJson *CreatePortRequest) (ports *proto.Port) {
	return &proto.Port{
		Id:          key,
		Name:        portJson.Name,
		City:        portJson.City,
		Province:    portJson.Province,
		Country:     portJson.Country,
		Alias:       portJson.Alias,
		Regions:     portJson.Regions,
		Coordinates: portJson.Coordinates,
		Timezone:    portJson.Timezone,
		Unlocs:      portJson.Unlocs,
		Code:        portJson.Code,
	}

}
