package apple_login

import (
	"reflect"
	"testing"
)

func TestGetJWKList(t *testing.T) {
	tests := []struct {
		name    string
		want    []*JWK
		wantErr bool
	}{
		{
			name: "noraml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getJWKList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetJWKList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJWKList() = %v, want %v", got, tt.want)
			}
		})
	}
}
