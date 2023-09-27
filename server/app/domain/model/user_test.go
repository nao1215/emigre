package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	type args struct {
		id        string
		name      string
		biography string
		email     string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				id:        "014KG56DC01GG4TEB01ZEX7WFJ",
				name:      "test",
				biography: "test",
				email:     "test@example.com",
			},
			want: &User{
				ID:        "014KG56DC01GG4TEB01ZEX7WFJ",
				Name:      "test",
				Biography: "test",
				Email:     "test@example.com",
			},
			wantErr: false,
		},
		{
			name: "invalid name: too short",
			args: args{
				id:        "014KG56DC01GG4TEB01ZEX7WFJ",
				name:      "",
				biography: "test",
				email:     "test@example.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid name: too long",
			args: args{
				id:        "014KG56DC01GG4TEB01ZEX7WFJ",
				name:      "012345678901234567890",
				biography: "test",
				email:     "test@example.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid biography: too short",
			args: args{
				id:        "014KG56DC01GG4TEB01ZEX7WFJ",
				name:      "test",
				biography: "",
				email:     "test@example.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid biography: too long",
			args: args{
				id:        "014KG56DC01GG4TEB01ZEX7WFJ",
				name:      "test",
				biography: "Harry Potter is a young wizard who attends Hogwarts School of Witchcraft and Wizardry. He discovers that he is the Boy Who Lived and embarks on a journey to defeat the dark wizard Lord Voldemort. Along the way, Harry learns about friendship, bravery, and the power of love. With his loyal friends Ron and Hermione by his side, they face challenges, unravel mysteries, and strive to bring peace to the magical world.",
				email:     "test@example.com",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := NewUser(tt.args.id, tt.args.name, tt.args.biography, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("value is mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
