package main

import "testing"

func TestPowerConsumption(t *testing.T) {
	tests := []struct {
		name    string
		ones    []int
		zeros   []int
		want    int
		wantErr bool
	}{
		{"sample 1",
			[]int{7, 5, 8, 7, 5},
			[]int{5, 7, 4, 5, 7},
			198,
			false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := powerConsumption(tt.ones, tt.zeros)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr = %v\n", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("wanted = %d, got = %d", tt.want, got)
			}
		})
	}
}
