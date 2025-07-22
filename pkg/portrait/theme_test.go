package portrait

import (
	"image/color"
	"reflect"
	"testing"
)

func TestTheme_Get(t *testing.T) {
	type args struct {
		theme string
	}
	tests := []struct {
		name    string
		args    args
		want    []color.Color
		wantErr bool
	}{
		{
			name:    "white theme",
			args:    args{theme: "white"},
			want:    Theme{}.basic(),
			wantErr: false,
		},
		{
			name:    "invalid theme",
			args:    args{theme: "invalid"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Theme{}
			got, err := tr.Get(tt.args.theme)
			if (err != nil) != tt.wantErr {
				t.Errorf("Theme.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Theme.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
