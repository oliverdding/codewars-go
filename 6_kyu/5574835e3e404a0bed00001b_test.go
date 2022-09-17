package kata

import "testing"

func TestGetParticipants(t *testing.T) {
	type args struct {
		h int
	}
	tests := []struct {
		name  string
		args  args
		wantN int
	}{
		{
			"1",
			args{
				1,
			},
			2,
		},
		{
			"6",
			args{
				6,
			},
			4,
		},
		{
			"7",
			args{
				7,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotN := GetParticipants(tt.args.h); gotN != tt.wantN {
				t.Errorf("GetParticipants() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
