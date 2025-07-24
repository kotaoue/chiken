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
			name:    "brown theme",
			args:    args{theme: "brown"},
			want:    Theme{}.brown(),
			wantErr: false,
		},
		{
			name:    "black theme",
			args:    args{theme: "black"},
			want:    Theme{}.black(),
			wantErr: false,
		},
		{
			name:    "brownBlack theme",
			args:    args{theme: "brownBlack"},
			want:    Theme{}.brownBlack(),
			wantErr: false,
		},
		{
			name:    "panda theme",
			args:    args{theme: "panda"},
			want:    Theme{}.panda(),
			wantErr: false,
		},
		{
			name:    "yellow theme",
			args:    args{theme: "yellow"},
			want:    Theme{}.yellow(),
			wantErr: false,
		},
		{
			name:    "green theme",
			args:    args{theme: "green"},
			want:    Theme{}.green(),
			wantErr: false,
		},
		{
			name:    "mossGreen theme",
			args:    args{theme: "mossGreen"},
			want:    Theme{}.mossGreen(),
			wantErr: false,
		},
		{
			name:    "lightBlue theme",
			args:    args{theme: "lightBlue"},
			want:    Theme{}.lightBlue(),
			wantErr: false,
		},
		{
			name:    "blue theme",
			args:    args{theme: "blue"},
			want:    Theme{}.blue(),
			wantErr: false,
		},
		{
			name:    "bluePurple theme",
			args:    args{theme: "bluePurple"},
			want:    Theme{}.bluePurple(),
			wantErr: false,
		},
		{
			name:    "purple theme",
			args:    args{theme: "purple"},
			want:    Theme{}.purple(),
			wantErr: false,
		},
		{
			name:    "pinkPurple theme",
			args:    args{theme: "pinkPurple"},
			want:    Theme{}.pinkPurple(),
			wantErr: false,
		},
		{
			name:    "pink theme",
			args:    args{theme: "pink"},
			want:    Theme{}.pink(),
			wantErr: false,
		},
		{
			name:    "red theme",
			args:    args{theme: "red"},
			want:    Theme{}.red(),
			wantErr: false,
		},
		{
			name:    "orange theme",
			args:    args{theme: "orange"},
			want:    Theme{}.orange(),
			wantErr: false,
		},
		{
			name:    "gray theme",
			args:    args{theme: "gray"},
			want:    Theme{}.gray(),
			wantErr: false,
		},
		{
			name:    "player2 theme",
			args:    args{theme: "player2"},
			want:    Theme{}.player2(),
			wantErr: false,
		},
		{
			name:    "player3 theme",
			args:    args{theme: "player3"},
			want:    Theme{}.player3(),
			wantErr: false,
		},
		{
			name:    "player4 theme",
			args:    args{theme: "player4"},
			want:    Theme{}.player4(),
			wantErr: false,
		},
		{
			name:    "player5 theme",
			args:    args{theme: "player5"},
			want:    Theme{}.player5(),
			wantErr: false,
		},
		{
			name:    "vivid theme",
			args:    args{theme: "vivid"},
			want:    Theme{}.vivid(),
			wantErr: false,
		},
		{
			name:    "random theme",
			args:    args{theme: "random"},
			wantErr: false,
		},
		{
			name:    "invalid theme",
			args:    args{theme: "invalid"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "party theme",
			args:    args{theme: "party180"},
			wantErr: false,
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

			if tt.name == "random theme" || tt.name == "party theme" {
				if len(got) != len(Theme{}.basic()) {
					t.Errorf("Theme.Get() len = %v, want %v", len(got), len(Theme{}.basic()))
				}
				for _, c := range got {
					if _, ok := c.(color.RGBA); !ok {
						t.Errorf("Theme.Get() color is not RGBA")
					}
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Theme.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTheme_Adjust(t *testing.T) {
	type args struct {
		theme string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "party adjust",
			args:    args{theme: "party2"},
			want:    "party0-party180",
			wantErr: false,
		},
		{
			name:    "invalid party adjust",
			args:    args{theme: "partyinvalid"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "no adjust",
			args:    args{theme: "white"},
			want:    "white",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Theme{}
			got, err := tr.Adjust(tt.args.theme)
			if (err != nil) != tt.wantErr {
				t.Errorf("Theme.Adjust() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Theme.Adjust() = %v, want %v", got, tt.want)
			}
		})
	}
}
