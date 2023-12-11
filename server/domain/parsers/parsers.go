package parsers

import (
	proto "github.com/rpuggian/ports-service/proto"
	"github.com/rpuggian/ports-service/server/domain/entity"
)

func ParsePortEntityToProtoResponse(p *entity.Port) *proto.Port {
	return &proto.Port{
		Id:          p.ID,
		Name:        p.Name,
		City:        p.City,
		Province:    p.Province,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}

func ParsePortProtoRequestToEntity(p *proto.Port) *entity.Port {
	return &entity.Port{
		ID:          p.Id,
		Name:        p.Name,
		City:        p.City,
		Province:    p.Province,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}
