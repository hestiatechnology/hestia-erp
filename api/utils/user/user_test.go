package user

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestVerifyAuthToken(t *testing.T) {
	type args struct {
		ctx   context.Context
		token uuid.UUID
	}
	tests := []struct {
		name    string
		args    args
		valid   bool
		expired bool
	}{
		{
			name: "Error (no rows)",
			args: args{
				ctx:   context.Background(),
				token: uuid.MustParse("8a43e239-83d1-4d25-86fb-b00e20569246"), // Random UUID
			},
			valid:   false,
			expired: false,
		},
		{
			name: "Valid token",
			args: args{
				ctx:   context.Background(),
				token: uuid.MustParse("fd0d33c9-05b2-4fae-9b19-9e41414cee27"),
			},
			valid:   true,
			expired: false,
		},
		{
			name: "Expired token",
			args: args{
				ctx:   context.Background(),
				token: uuid.MustParse("5d82b60a-7bc9-4fa8-be79-d7bb70435dc5"),
			},
			valid:   false,
			expired: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, expired := VerifyAuthToken(tt.args.ctx, tt.args.token)
			if valid != tt.valid {
				t.Errorf("VerifyAuthToken() valid = %v, want %v", valid, tt.valid)
			}
			if expired != tt.expired {
				t.Errorf("VerifyAuthToken() expired = %v, want %v", expired, tt.expired)
			}
		})
	}
}
