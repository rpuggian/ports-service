package request

import (
	proto "github.com/rpuggian/ports-service/proto"
	"reflect"
	"testing"
)

func TestParseUploadPortRequestToProtoPortArray(t *testing.T) {
	type args struct {
		portsJson UploadPortByFileRequest
	}
	tests := []struct {
		name      string
		args      args
		wantPorts []*proto.Port
	}{
		{
			name: "Should parse JSON request to proto filling all fields",
			args: args{
				portsJson: UploadPortByFileRequest{
					"P1": CreatePortRequest{
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
					"P2": CreatePortRequest{
						Name:        "P2",
						City:        "C2",
						Country:     "CO2",
						Alias:       []string{"A3", "A4"},
						Regions:     []string{"R3", "R4"},
						Coordinates: []float64{3.3, 4.4},
						Province:    "PR2",
						Timezone:    "TZ2",
						Unlocs:      []string{"U3", "U4"},
						Code:        "C2",
					},
				},
			},
			wantPorts: []*proto.Port{
				{
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
				{
					Id:          "P2",
					Name:        "P2",
					City:        "C2",
					Country:     "CO2",
					Alias:       []string{"A3", "A4"},
					Regions:     []string{"R3", "R4"},
					Coordinates: []float64{3.3, 4.4},
					Province:    "PR2",
					Timezone:    "TZ2",
					Unlocs:      []string{"U3", "U4"},
					Code:        "C2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPorts := ParseUploadPortRequestToProtoPortArray(tt.args.portsJson); !reflect.DeepEqual(gotPorts, tt.wantPorts) {
				t.Errorf("ParseUploadPortRequestToProtoPortArray() = %v, want %v", gotPorts, tt.wantPorts)
			}
		})
	}
}
