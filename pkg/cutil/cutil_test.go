package cutil

import (
	"image/color"
	"reflect"
	"testing"
)

func TestHexToColor(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    *color.RGBA
		wantErr bool
	}{
		{
			name:    "valid hex",
			s:       "#ffffff",
			want:    &color.RGBA{R: 255, G: 255, B: 255, A: 255},
			wantErr: false,
		},
		{
			name:    "invalid hex",
			s:       "invalid",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToColor(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexToColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
