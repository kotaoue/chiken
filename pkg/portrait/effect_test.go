package portrait

import (
	"image/color"
	"reflect"
	"testing"
)

func TestEffect_Apply(t *testing.T) {
	type fields struct {
		style [][]int
		theme []color.Color
	}
	type args struct {
		effects string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]int
		want1   []color.Color
		wantErr bool
	}{
		{
			name: "mirror",
			fields: fields{
				style: [][]int{{1, 2}, {3, 4}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			},
			args:    args{effects: "mirror"},
			want:    [][]int{{2, 1}, {4, 3}},
			want1:   []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Effect{
				style: tt.fields.style,
				theme: tt.fields.theme,
			}
			got, got1, err := e.Apply(tt.args.effects)
			if (err != nil) != tt.wantErr {
				t.Errorf("Effect.Apply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Effect.Apply() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Effect.Apply() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
