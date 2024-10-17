package idm

import (
	"context"
	"hestia/api/pb/idmanagement"
	"testing"
)

func TestIdentityManagementServer_Login(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *idmanagement.LoginRequest
	}
	tests := []struct {
		name    string
		s       *IdentityManagementServer
		args    args
		want    *idmanagement.LoginResponse
		wantErr bool
	}{
		{
			name: "No login",
			s:    &IdentityManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &idmanagement.LoginRequest{
					Email:    "daniel",
					Password: "password",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "No password",
			s:    &IdentityManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &idmanagement.LoginRequest{
					Email:    "daniel",
					Password: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid password",
			s:    &IdentityManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &idmanagement.LoginRequest{
					Email:    "daniel",
					Password: "password",
				},
			},
			want:    nil,
			wantErr: true,
		},
		//TODO
		{
			name: "Success",
			s:    &IdentityManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &idmanagement.LoginRequest{
					Email:    "valid@test.hestiatechnology.pt",
					Password: "5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8", //password
				},
			},
			want:    &idmanagement.LoginResponse{Token: "ceb9c1be-a7a3-4f45-bcdc-cd69b541be70", Name: "Test User", Email: "valid@test.hestiatechnology.pt"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Login(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentityManagementServer.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got == nil && tt.want == nil {
				return
			}

			// Set the token to a fixed value so we can compare the response
			if got != nil {
				got.Token = "ceb9c1be-a7a3-4f45-bcdc-cd69b541be70"
			}

			if got.Token != tt.want.Token || got.Name != tt.want.Name || got.Email != tt.want.Email {
				t.Errorf("IdentityManagementServer.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
