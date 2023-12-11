package response

import (
	proto "github.com/rpuggian/ports-service/proto"
	"reflect"
	"testing"
)

func TestNewUploadPortByFileResponse(t *testing.T) {
	type args struct {
		response *proto.CreateResponse
	}
	tests := []struct {
		name string
		args args
		want *UploadPortByFileResponse
	}{
		{
			name: "Should create response with total",
			args: args{
				response: &proto.CreateResponse{
					Total: 10,
				},
			},
			want: &UploadPortByFileResponse{
				Total: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUploadPortByFileResponse(tt.args.response); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUploadPortByFileResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFindPortResponse(t *testing.T) {
	type args struct {
		port *proto.Port
	}
	tests := []struct {
		name string
		args args
		want *FindPortResponse
	}{
		{
			name: "Should create response with port",
			args: args{
				port: &proto.Port{
					Id:          "id",
					Name:        "name",
					City:        "city",
					Province:    "province",
					Country:     "country",
					Alias:       []string{"alias"},
					Regions:     []string{"regions"},
					Coordinates: []float64{1.1, 2.2},
					Timezone:    "timezone",
					Unlocs:      []string{"unlocs"},
					Code:        "code",
				},
			},
			want: &FindPortResponse{
				ID:          "id",
				Name:        "name",
				City:        "city",
				Province:    "province",
				Country:     "country",
				Alias:       []string{"alias"},
				Regions:     []string{"regions"},
				Coordinates: []float64{1.1, 2.2},
				Timezone:    "timezone",
				Unlocs:      []string{"unlocs"},
				Code:        "code",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFindPortResponse(tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFindPortResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
