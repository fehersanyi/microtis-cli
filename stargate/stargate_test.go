package stargate

import (
	"strings"
	"testing"
)

func Test_getLocation(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"exact", "/Users/fehersandor/go/src/github.com/fehersanyi/microtis-cli/star_gate", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLocation()
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.TrimSpace(got) != tt.want {
				t.Errorf("getLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}
