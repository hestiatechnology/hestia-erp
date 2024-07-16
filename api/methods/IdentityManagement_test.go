package methods

import (
	"context"
	"hestia/api/pb/idmanagement"
	"reflect"
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
		{
			name: "Success",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Login(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("IdentityManagementServer.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IdentityManagementServer.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
