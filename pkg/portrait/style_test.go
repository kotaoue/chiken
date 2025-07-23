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
		s       Style
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name:    "basic style",
			s:       Style{},
			args:    args{style: BasicStyle},
			want:    Style{}.basic(),
			wantErr: false,
		},
		{
			name:    "walk style",
			s:       Style{},
			args:    args{style: WalkStyle},
			want:    Style{}.walk(),
			wantErr: false,
		},
		{
			name:    "wide style",
			s:       Style{},
			args:    args{style: WideStyle},
			want:    Style{}.wide(),
			wantErr: false,
		},
		{
			name:    "tiptoe style",
			s:       Style{},
			args:    args{style: TiptoeStyle},
			want:    Style{}.tiptoe(),
			wantErr: false,
		},
		{
			name:    "jump style",
			s:       Style{},
			args:    args{style: JumpStyle},
			want:    Style{}.jump(),
			wantErr: false,
		},
		{
			name:    "sleep style",
			s:       Style{},
			args:    args{style: SleepStyle},
			want:    Style{}.sleep(),
			wantErr: false,
		},
		{
			name:    "deepSleep style",
			s:       Style{},
			args:    args{style: DeepSleepStyle},
			want:    Style{}.deepSleep(),
			wantErr: false,
		},
		{
			name:    "wake style",
			s:       Style{},
			args:    args{style: WakeStyle},
			want:    Style{}.wake(),
			wantErr: false,
		},
		{
			name:    "invalid style",
			s:       Style{},
			args:    args{style: "invalid"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Get(tt.args.style)
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