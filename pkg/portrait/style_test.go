package portrait

import (
	"reflect"
	"testing"
)

func TestStyle_Get(t *testing.T) {
	type args struct {
		style string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name:    "basic style",
			args:    args{style: "basic"},
			want:    Style{}.basic(),
			wantErr: false,
		},
		{
			name:    "walk style",
			args:    args{style: "walk"},
			want:    Style{}.walk(),
			wantErr: false,
		},
		{
			name:    "invalid style",
			args:    args{style: "invalid"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Style{}
			got, err := s.Get(tt.args.style)
			if (err != nil) != tt.wantErr {
				t.Errorf("Style.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Style.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
