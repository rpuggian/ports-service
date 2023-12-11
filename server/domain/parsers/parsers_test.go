package parsers

import (
	proto "github.com/rpuggian/ports-service/proto"
	"github.com/rpuggian/ports-service/server/domain/entity"
	"reflect"
	"testing"
)

func TestParsePortEntityToProtoResponse(t *testing.T) {
	type args struct {
		p *entity.Port
	}
	tests := []struct {
		name string
		args args
		want *proto.Port
	}{
		{
			name: "Should test if the entity parsed to proto is filled with the same values",
			args: args{
				p: &entity.Port{
					ID:          "ID",
					Name:        "Name",
					City:        "City",
					Province:    "Province",
					Country:     "Country",
					Alias:       []string{"Alias"},
					Regions:     []string{"Regions"},
					Coordinates: []float64{1.1, 2.2},
					Timezone:    "Timezone",
					Unlocs:      []string{"Unlocs"},
					Code:        "Code",
				},
			},
			want: &proto.Port{
				Id:          "ID",
				Name:        "Name",
				City:        "City",
				Province:    "Province",
				Country:     "Country",
				Alias:       []string{"Alias"},
				Regions:     []string{"Regions"},
				Coordinates: []float64{1.1, 2.2},
				Timezone:    "Timezone",
				Unlocs:      []string{"Unlocs"},
				Code:        "Code",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParsePortEntityToProtoResponse(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePortEntityToProtoResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParsePortProtoRequestToEntity(t *testing.T) {
	type args struct {
		p *proto.Port
	}
	tests := []struct {
		name string
		args args
		want *entity.Port
	}{
		{
			name: "Should test if the proto parsed to entity is filled with the same values",
			args: args{
				p: &proto.Port{
					Id:          "ID",
					Name:        "Name",
					City:        "City",
					Province:    "Province",
					Country:     "Country",
					Alias:       []string{"Alias"},
					Regions:     []string{"Regions"},
					Coordinates: []float64{1.1, 2.2},
					Timezone:    "Timezone",
					Unlocs:      []string{"Unlocs"},
					Code:        "Code",
				},
			},
			want: &entity.Port{
				ID:          "ID",
				Name:        "Name",
				City:        "City",
				Province:    "Province",
				Country:     "Country",
				Alias:       []string{"Alias"},
				Regions:     []string{"Regions"},
				Coordinates: []float64{1.1, 2.2},
				Timezone:    "Timezone",
				Unlocs:      []string{"Unlocs"},
				Code:        "Code",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParsePortProtoRequestToEntity(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePortProtoRequestToEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
