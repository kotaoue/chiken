package cutil

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
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
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
