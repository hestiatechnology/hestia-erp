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

func TestCompanyManagementServer_CreateCompany(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *company.CreateCompanyRequest
	}
	tests := []struct {
		name    string
		s       *CompanyManagementServer
		args    args
		want    *company.Id
		wantErr bool
	}{
		{
			name: "Missing company name",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in:  &company.CreateCompanyRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Missing VAT ID",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.CreateCompanyRequest{
					Name: "Test Company",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Missing SSN",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.CreateCompanyRequest{
					Name:  "Test Company",
					VatId: "123456789",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Missing location",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.CreateCompanyRequest{
					Name:  "Test Company",
					VatId: "123456789",
					Ssn:   12345678911,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Missing address",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.CreateCompanyRequest{
					Name:  "Test Company",
					VatId: "123456789",
					Ssn:   12345678911,
					Location: &company.Location{
						Locality:   "Test Locality",
						PostalCode: "1234-567",
						Country:    "PT",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Missing locality",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.CreateCompanyRequest{
					Name:  "Test Company",
					VatId: "123456789",
					Ssn:   123456789,
					Location: &company.Location{
						Address:    "Test Address",
						PostalCode: "1234-567",
						Country:    "PT",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Missing Postal Code",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.CreateCompanyRequest{
					Name:  "Test Company",
					VatId: "123456789",
					Ssn:   123456789,
					Location: &company.Location{
						Address:  "Test Address",
						Locality: "Test Locality",
						Country:  "PT",
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Missing Country",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.CreateCompanyRequest{
					Name:  "Test Company",
					VatId: "123456789",
					Ssn:   123456789,
					Location: &company.Location{
						Address:    "Test Address",
						Locality:   "Test Locality",
						PostalCode: "1234-567",
					},
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
				in: &company.CreateCompanyRequest{
					Name:  "Test Company",
					VatId: "123456778",
					Ssn:   123456789,
					Location: &company.Location{
						Address:    "Test Address",
						Locality:   "Test Locality",
						PostalCode: "1234-567",
						Country:    "PT",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Company already added via VAT ID or SSN",
			s:    &CompanyManagementServer{},
			args: args{
				ctx: context.Background(),
				in: &company.CreateCompanyRequest{
					Name:  "Test Company",
					VatId: "123456778",
					Ssn:   123456789,
					Location: &company.Location{
						Address:    "Test Address",
						Locality:   "Test Locality",
						PostalCode: "1234-567",
						Country:    "PT",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateCompany(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("CompanyManagementServer.CreateCompany() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if doesnt want error, check if a uuid was returned
			if !tt.wantErr && got.GetId() == "" {
				t.Errorf("CompanyManagementServer.CreateCompany() = %v, want %v", got, tt.want)
			}
		})
	}
}
