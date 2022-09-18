package kata

import "testing"

func TestHumanReadableTime(t *testing.T) {
	type args struct {
		seconds int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				0,
			},
			"00:00:00",
		},
		{
			"2",
			args{
				59,
			},
			"00:00:59",
		},
		{
			"3",
			args{
				60,
			},
			"00:01:00",
		},
		{
			"4",
			args{
				3599,
			},
			"00:59:59",
		},
		{
			"5",
			args{
				3600,
			},
			"01:00:00",
		},
		{
			"6",
			args{
				86399,
			},
			"23:59:59",
		},
		{
			"7",
			args{
				86400,
			},
			"24:00:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumanReadableTime(tt.args.seconds); got != tt.want {
				t.Errorf("HumanReadableTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
