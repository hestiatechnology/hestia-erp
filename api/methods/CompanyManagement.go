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
			// Theres no user with that email
			// TODO: if employeeId is provided, check if there's a user with that employeeId in the company
			// if there isnt, create a new user with the email and associate it with the company
			// TODO: Send an email to the user to invite him to the platform
			logger.WarningLogger.Println("Implement email sending to invite user to the platform")
			return nil, herror.StatusWithInfo(codes.NotFound, "User not found", herror.InvalidUser, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		} else {
			logger.ErrorLogger.Println(err)
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		}

	}

	// Check if the user is already in the company
	var count int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM users.user_company WHERE user_id = $1 AND company_id = $2", userId, companyUuid).Scan(&count)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	if count > 0 {
		return nil, herror.StatusWithInfo(codes.AlreadyExists, "User already in the company", herror.UserAlreadyExists, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// if employeeId is provided, check if the there's already a user with that employeeId in the company
	if employeeId != "" {
		var count int
		err = db.QueryRow(ctx, "SELECT COUNT(*) FROM users.user_company WHERE employee_id = $1 AND company_id = $2", in.GetEmployeeId(), in.GetCompanyId()).Scan(&count)
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		}

		if count > 0 {
			return nil, herror.StatusWithInfo(codes.AlreadyExists, "Employee already in the company", herror.EmployeeAlreadyExists, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		}

		// Associate user with the company
		_, err = tx.Exec(ctx, "INSERT INTO users.user_company (user_id, company_id, employee_id) VALUES ($1, $2, $3)", userId, in.GetCompanyId(), in.GetEmployeeId())

		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				//// Check if the error also sho
				//if pgErr.Code == "23505" {
				//	return nil, status.Error(codes.AlreadyExists, "User already in the company")
				//}
				logger.ErrorLogger.Println(pgErr.ColumnName, pgErr.ConstraintName, pgErr.Error())
			}
			logger.ErrorLogger.Println(err)
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		}
	} else {
		// Associate user with the company
		_, err = tx.Exec(ctx, "INSERT INTO users.user_company (user_id, company_id) VALUES ($1, $2)", userId, in.GetCompanyId())
		if err != nil {
			logger.ErrorLogger.Println(err)
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
		}
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		logger.ErrorLogger.Println(err)
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to commit transaction", herror.DatabaseTxError, company.CompanyManagement_ServiceDesc.ServiceName, nil).Err()
	}

	return &emptypb.Empty{}, nil
}
