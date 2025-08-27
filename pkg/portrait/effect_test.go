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
		{
			name: "negative",
			fields: fields{
				style: [][]int{{1, 2}, {3, 4}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}, color.RGBA{R: 255, G: 255, B: 255, A: 255}},
			},
			args:    args{effects: "negative"},
			want:    [][]int{{1, 2}, {3, 4}},
			want1:   []color.Color{color.RGBA{R: 255, G: 255, B: 255, A: 255}, color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			wantErr: false,
		},
		{
			name: "grayscale",
			fields: fields{
				style: [][]int{{1, 2}, {3, 4}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}, color.RGBA{R: 255, G: 0, B: 0, A: 255}},
			},
			args:    args{effects: "grayscale"},
			want:    [][]int{{1, 2}, {3, 4}},
			want1:   []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}, color.RGBA{R: 76, G: 76, B: 76, A: 255}},
			wantErr: false,
		},
		{
			name: "rotateClockwise",
			fields: fields{
				style: [][]int{{1, 2}, {3, 4}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			},
			args:    args{effects: "rotateClockwise"},
			want:    [][]int{{3, 1}, {4, 2}},
			want1:   []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			wantErr: false,
		},
		{
			name: "rotateCounterClockwise",
			fields: fields{
				style: [][]int{{1, 2}, {3, 4}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			},
			args:    args{effects: "rotateCounterClockwise"},
			want:    [][]int{{2, 4}, {1, 3}},
			want1:   []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			wantErr: false,
		},
		{
			name: "rightLoop",
			fields: fields{
				style: [][]int{{1, 2, 3}, {4, 5, 6}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			},
			args:    args{effects: "rightLoop1"},
			want:    [][]int{{3, 1, 2}, {6, 4, 5}},
			want1:   []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			wantErr: false,
		},
		{
			name: "leftLoop",
			fields: fields{
				style: [][]int{{1, 2, 3}, {4, 5, 6}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			},
			args:    args{effects: "leftLoop1"},
			want:    [][]int{{2, 3, 1}, {5, 6, 4}},
			want1:   []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			wantErr: false,
		},
		{
			name: "upLoop",
			fields: fields{
				style: [][]int{{1, 2}, {3, 4}, {5, 6}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			},
			args:    args{effects: "upLoop1"},
			want:    [][]int{{3, 4}, {5, 6}, {1, 2}},
			want1:   []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			wantErr: false,
		},
		{
			name: "downLoop",
			fields: fields{
				style: [][]int{{1, 2}, {3, 4}, {5, 6}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			},
			args:    args{effects: "downLoop1"},
			want:    [][]int{{5, 6}, {1, 2}, {3, 4}},
			want1:   []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			wantErr: false,
		},
		{
			name: "invalid effect",
			fields: fields{
				style: [][]int{{1, 2}, {3, 4}},
				theme: []color.Color{color.RGBA{R: 0, G: 0, B: 0, A: 255}},
			},
			args:    args{effects: "invalid"},
			want:    nil,
			want1:   nil,
			wantErr: true,
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

func TestEffect_Adjust(t *testing.T) {
	type fields struct {
		style [][]int
		theme []color.Color
	}
	type args struct {
		effects string
		size    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "rightLoop",
			fields: fields{
				style: [][]int{},
				theme: []color.Color{},
			},
			args:    args{effects: "rightLoop2", size: 4},
			want:    "rightLoop0-rightLoop2",
			wantErr: false,
		},
		{
			name: "leftLoop",
			fields: fields{
				style: [][]int{},
				theme: []color.Color{},
			},
			args:    args{effects: "leftLoop2", size: 4},
			want:    "leftLoop0-leftLoop2",
			wantErr: false,
		},
		{
			name: "upLoop",
			fields: fields{
				style: [][]int{},
				theme: []color.Color{},
			},
			args:    args{effects: "upLoop2", size: 4},
			want:    "upLoop0-upLoop2",
			wantErr: false,
		},
		{
			name: "downLoop",
			fields: fields{
				style: [][]int{},
				theme: []color.Color{},
			},
			args:    args{effects: "downLoop2", size: 4},
			want:    "downLoop0-downLoop2",
			wantErr: false,
		},
		{
			name: "no loop",
			fields: fields{
				style: [][]int{},
				theme: []color.Color{},
			},
			args:    args{effects: "mirror", size: 4},
			want:    "mirror",
			wantErr: false,
		},
		{
			name: "invalid loop",
			fields: fields{
				style: [][]int{},
				theme: []color.Color{},
			},
			args:    args{effects: "rightLoopinvalid", size: 4},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Effect{
				style: tt.fields.style,
				theme: tt.fields.theme,
			}
			got, err := e.Adjust(tt.args.effects, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Effect.Adjust() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Effect.Adjust() = %v, want %v", got, tt.want)
			}
		})
	}
}
