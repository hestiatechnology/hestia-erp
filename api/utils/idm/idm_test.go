package idm

import (
	"context"
	"testing"
)

func TestRandomSalt(t *testing.T) {
	tests := []struct {
		name     string
		testFunc func(t *testing.T, got string)
	}{
		{
			name: "Check non-empty and length",
			testFunc: func(t *testing.T, got string) {
				if got == "" {
					t.Error("RandomSalt() returned an empty string, want non-empty")
				}
				if len(got) != 64 {
					t.Errorf("RandomSalt() returned a string of length %d, want 64", len(got))
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomSalt()
			tt.testFunc(t, got)
		})
	}
}

func TestPasswordHash(t *testing.T) {
	type args struct {
		password string
		salt     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Not a SHA-256 hash",
			args: args{
				password: "password",
				salt:     "salt",
			},
			want: "",
		},
		{
			name: "Empty salt",
			args: args{
				password: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
				salt:     "",
			},
			want: "",
		},
		{
			name: "Success",
			args: args{
				password: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8",
				salt:     "mySalt",
			},
			want: "4a1a4f8362882bba5f9a8d484846319e9c82a601e87b3a29d72f16ba703d88d4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PasswordHash(tt.args.password, tt.args.salt); got != tt.want {
				t.Errorf("PasswordHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSalt(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Failure",
			args: args{
				ctx:   context.Background(),
				email: "random@email.com",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Success",
			args: args{
				ctx:   context.Background(),
				email: "valid@test.hestiatechnology.pt",
			},
			want:    "4192230aa6421de5c43711119c97056bdf9462099c575dc1651ac9d7e533a5de",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSalt(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSalt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetSalt() = %v, want %v", got, tt.want)
			}
		})
	}
}
