package methods

import (
	"context"
	"errors"

	"hestia/api/pb/company"
	"hestia/api/utils/db"
	"hestia/api/utils/herror"
	"hestia/api/utils/logger"

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
	return &company.Id{Id: uuid.NewString()}, nil
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
		logger.ErrorLogger.Println(err)
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseConnError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		logger.ErrorLogger.Println(err)
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
			logger.ErrorLogger.Println(err)
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
						logger.DebugLogger.Println(err)
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
		logger.ErrorLogger.Println(err)
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to commit transaction", herror.DatabaseTxError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	return &emptypb.Empty{}, nil
}
