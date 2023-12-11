package request

import (
	proto "github.com/rpuggian/ports-service/proto"
	"reflect"
	"testing"
)

func TestParseUploadPortRequestToProtoPortArray(t *testing.T) {
	type args struct {
		portsJson CreatePortRequest
	}
	tests := []struct {
		name      string
		args      args
		wantPorts *proto.Port
	}{
		{
			name: "Should parse JSON request to proto filling all fields",
			args: args{
				portsJson: CreatePortRequest{
					Name:        "P1",
					City:        "C1",
					Country:     "CO1",
					Alias:       []string{"A1", "A2"},
					Regions:     []string{"R1", "R2"},
					Coordinates: []float64{1.1, 2.2},
					Province:    "PR1",
					Timezone:    "TZ1",
					Unlocs:      []string{"U1", "U2"},
					Code:        "C1",
				},
			},
			wantPorts: &proto.Port{
				Id:          "P1",
				Name:        "P1",
				City:        "C1",
				Country:     "CO1",
				Alias:       []string{"A1", "A2"},
				Regions:     []string{"R1", "R2"},
				Coordinates: []float64{1.1, 2.2},
				Province:    "PR1",
				Timezone:    "TZ1",
				Unlocs:      []string{"U1", "U2"},
				Code:        "C1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPorts := ParseUploadPortRequestToProtoPortArray(tt.args.portsJson.Name, &tt.args.portsJson); !reflect.DeepEqual(gotPorts, tt.wantPorts) {
				t.Errorf("ParseUploadPortRequestToProtoPortArray() = %v, want %v", gotPorts, tt.wantPorts)
			}
		})
	}
}
