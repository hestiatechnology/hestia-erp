package company

import (
	"context"
	"errors"
	"strconv"

	"hestia/api/pb/company"
	"hestia/api/utils/db"
	"hestia/api/utils/herror"
	"hestia/api/utils/storage"

	"github.com/hestiatechnology/autoridadetributaria/common"
	"github.com/rs/zerolog/log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CompanyManagementServer struct {
	company.UnimplementedCompanyManagementServer
}

func (s *CompanyManagementServer) CreateCompany(ctx context.Context, in *company.CreateCompanyRequest) (*company.Id, error) {
	//  Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Indicates whether the company is a sole trader or not
	// IsSoleTrader   bool      `protobuf:"varint,2,opt,name=isSoleTrader,proto3" json:"isSoleTrader,omitempty"`
	// CommercialName *string   `protobuf:"bytes,3,opt,name=commercialName,proto3,oneof" json:"commercialName,omitempty"`
	// VatId          int32     `protobuf:"varint,4,opt,name=vatId,proto3" json:"vatId,omitempty"`
	// Ssn            int32     `protobuf:"varint,5,opt,name=ssn,proto3" json:"ssn,omitempty"`
	// Location
	name := in.GetName()
	isSoleTrader := in.GetIsSoleTrader()
	commercialName := in.GetCommercialName()
	vatId := in.GetVatId()
	ssn := in.GetSsn()
	location := in.GetLocation()
	address := location.GetAddress()
	locality := location.GetLocality()
	postalCode := location.GetPostalCode()
	country := location.GetCountry()

	if name == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing name", []*errdetails.BadRequest_FieldViolation{{
			Field:       "name",
			Description: "Name is required",
		}}).Err()
	}

	if vatId == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing VAT ID", []*errdetails.BadRequest_FieldViolation{{
			Field:       "vatId",
			Description: "VAT ID is required",
		}}).Err()
	}

	// TODO: Implement the validation function on autoridadetributaria package

	if ssn == 0 {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing SSN", []*errdetails.BadRequest_FieldViolation{{
			Field:       "ssn",
			Description: "SSN is required",
		}}).Err()
	}
	// TODO: Implement the validation function on fiscal package

	if location == nil {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing location", []*errdetails.BadRequest_FieldViolation{{
			Field:       "location",
			Description: "Location is required",
		}}).Err()
	}

	if address == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing address", []*errdetails.BadRequest_FieldViolation{{
			Field:       "location.address",
			Description: "Address is required",
		}}).Err()
	}
	if locality == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing city", []*errdetails.BadRequest_FieldViolation{{
			Field:       "location.city",
			Description: "City is required",
		}}).Err()
	}
	if postalCode == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing postal code", []*errdetails.BadRequest_FieldViolation{{
			Field:       "location.postalCode",
			Description: "Postal code is required",
		}}).Err()
	}
	if country == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing country", []*errdetails.BadRequest_FieldViolation{{
			Field:       "location.country",
			Description: "Country is required",
		}}).Err()
	}

	countryFound := false
	for _, c := range common.CountryCodes {
		if country == c {
			countryFound = true
			break
		}
	}

	if !countryFound {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Invalid country code", []*errdetails.BadRequest_FieldViolation{{
			Field:       "location.country",
			Description: "Country code is invalid",
		}}).Err()
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Error().Err(err).Msg("Error getting database connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseConnError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to start transaction", herror.DatabaseTxError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}
	defer tx.Rollback(ctx)

	// Insert the company
	companyId := uuid.New()
	_, err = tx.Exec(ctx, "INSERT INTO companies.company (id, name, vat_id, ssn, street, locality, postal_code, country_id) VALUES ($1, $2, $3, $4, $5, $6, $7, (SELECT id FROM general.country WHERE code = $8))", companyId, name, vatId, strconv.Itoa(int(ssn)), address, locality, postalCode, country)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.ConstraintName == "pk_company_id" {
				return nil, herror.StatusWithInfo(codes.AlreadyExists, "Company already exists", herror.CompanyAlreadyExists, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
			} else if pgErr.ConstraintName == "uq_company_vat_id" {
				return nil, herror.StatusWithInfo(codes.AlreadyExists, "VAT ID already exists", herror.CompanyAlreadyExists, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
			} else if pgErr.ConstraintName == "uq_company_ssn" {
				return nil, herror.StatusWithInfo(codes.AlreadyExists, "SSN already exists", herror.CompanyAlreadyExists, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
			}
		}

		log.Error().Err(err).Msg("Error inserting company")
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	if commercialName != "" {
		_, err = tx.Exec(ctx, "UPDATE companies.company SET commercial_name = $1 WHERE id = $2", commercialName, companyId)
		if err != nil {
			log.Error().Err(err).Str("companyId", companyId.String()).Str("commercialName", commercialName).Msg("Error updating commercial name")
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		}
	}
	if isSoleTrader {
		_, err = tx.Exec(ctx, "UPDATE companies.company SET is_sole_trader = $1 WHERE id = $2", isSoleTrader, companyId)
		if err != nil {
			log.Error().Err(err).Str("companyId", companyId.String()).Bool("isSoleTrader", isSoleTrader).Msg("Error updating is_sole_trader")
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		}
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error committing transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to commit transaction", herror.DatabaseTxError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	go storage.CreateS3Bucket(companyId)

	return &company.Id{Id: companyId.String()}, nil
}

func (s *CompanyManagementServer) AddUserToCompany(ctx context.Context, in *company.AddUserToCompanyRequest) (*emptypb.Empty, error) {
	email := in.GetEmail()
	companyId := in.GetCompanyId()
	employeeId := in.GetEmployeeId()

	if email == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing email", []*errdetails.BadRequest_FieldViolation{{
			Field:       "email",
			Description: "Email is required",
		}}).Err()
	}
	if companyId == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing company ID", []*errdetails.BadRequest_FieldViolation{{
			Field:       "companyId",
			Description: "Company ID is required",
		}}).Err()
	}

	companyUuid, err := uuid.Parse(companyId)
	if err != nil {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Invalid company ID", []*errdetails.BadRequest_FieldViolation{{
			Field:       "companyId",
			Description: "Company Id is not a valid UUID",
		}}).Err()
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Error().Err(err).Msg("Error getting database connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseConnError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to start transaction", herror.DatabaseTxError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}
	defer tx.Rollback(ctx)

	// Get user id
	var userId uuid.UUID
	err = db.QueryRow(ctx, "SELECT id FROM users.users WHERE email = $1", email).Scan(&userId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Frontend should ask the user for an invite to the company
			return nil, herror.StatusWithInfo(codes.NotFound, "User not found", herror.InvalidUser, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		} else {
			log.Error().Err(err).Str("email", email).Msg("Error getting user id")
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		}

	}

	employeeUuid, err := uuid.Parse(employeeId)
	if err != nil && employeeId != "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Invalid employee ID", []*errdetails.BadRequest_FieldViolation{{
			Field:       "employeeId",
			Description: "Employee Id is not a valid UUID",
		}}).Err()
	}

	if err != nil {
		_, err = tx.Exec(ctx, "INSERT INTO users.user_company (user_id, company_id, employee_id) VALUES ($1, $2, $3)", userId, companyUuid, nil)
	} else {
		_, err = tx.Exec(ctx, "INSERT INTO users.user_company (user_id, company_id, employee_id) VALUES ($1, $2, $3)", userId, companyUuid, employeeUuid)
	}

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.ConstraintName == "pk_user_company" {
				if employeeId != "" {
					_, err = tx.Exec(ctx, "UPDATE users.user_company SET employee_id = $1 WHERE user_id = $2 AND company_id = $3", employeeUuid, userId, companyUuid)
					if err != nil {
						log.Error().Err(pgErr).Str("employeeId", employeeId).Str("userId", userId.String()).Str("companyId", companyId).Msg("Error updating employee id")
						return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
					}
					return &emptypb.Empty{}, nil
				}
				return nil, herror.StatusWithInfo(codes.AlreadyExists, "User already in the company", herror.UserAlreadyExists, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
			} else if pgErr.ConstraintName == "fk_user_company_company" {
				return nil, herror.StatusWithInfo(codes.NotFound, "Company not found", herror.InvalidCompany, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
			} else if pgErr.ConstraintName == "fk_user_company_users" {
				return nil, herror.StatusWithInfo(codes.NotFound, "User not found", herror.InvalidUser, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
			} else if pgErr.ConstraintName == "fk_user_company_users_employee_id" {
				return nil, herror.StatusWithInfo(codes.NotFound, "Employee not found", herror.InvalidEmployee, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
			}
		}

		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error committing transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to commit transaction", herror.DatabaseTxError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	return &emptypb.Empty{}, nil
}
