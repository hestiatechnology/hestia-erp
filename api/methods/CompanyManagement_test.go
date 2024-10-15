package methods

import (
	"context"
	"hestia/api/pb/company"
	"reflect"
	"testing"

	"google.golang.org/protobuf/types/known/emptypb"
)

func TestCompanyManagementServer_AddUserToCompany(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *company.AddUserToCompanyRequest
	}
	tests := []struct {
		name    string
		s       *CompanyManagementServer
		args    args
		want    *emptypb.Empty
		wantErr bool
	}{
		{
			name: "Fail: Wrong company ID",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.AddUserToCompanyRequest{
					Email:     "valid@test.hestiatechnology.pt",
					CompanyId: "6c8f45be-808c-4162-9f4c-b9783b55d21f",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fail: Missing company ID",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.AddUserToCompanyRequest{
					Email:     "valid@test.hestiatechnology.pt",
					CompanyId: "",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Fail: Missing email",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.AddUserToCompanyRequest{
					Email:     "",
					CompanyId: "6c8f45be-808c-4162-9f4c-b9783b55d21f",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.AddUserToCompanyRequest{
					Email:     "valid@test.hestiatechnology.pt",
					CompanyId: "6c8f45be-808c-4162-9f4c-b9783b55d28f",
				},
			},
			want:    &emptypb.Empty{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.AddUserToCompany(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyManagementServer.AddUserToCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompanyManagementServer.AddUserToCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}
