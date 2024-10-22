package idm

import (
	"context"
	"errors"
	"time"

	"hestia/api/pb/company"
	"hestia/api/pb/idmanagement"
	"hestia/api/utils/db"
	"hestia/api/utils/herror"
	"hestia/api/utils/idm"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IdentityManagementServer struct {
	idmanagement.UnimplementedIdentityManagementServer
}

func (s *IdentityManagementServer) Login(ctx context.Context, in *idmanagement.LoginRequest) (*idmanagement.LoginResponse, error) {
	email := in.GetEmail()
	password := in.GetPassword()

	if email == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing email", []*errdetails.BadRequest_FieldViolation{{
			Field:       "email",
			Description: "Email is required",
		}}).Err()
	}
	if password == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing password", []*errdetails.BadRequest_FieldViolation{{
			Field:       "password",
			Description: "Password is required",
		}}).Err()
	} else if len(password) != 64 {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Invalid password", []*errdetails.BadRequest_FieldViolation{{
			Field:       "password",
			Description: "Password must be 64 characters long (SHA-256 hash)",
		}}).Err()
	}

	//Check if user exists in the database
	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Error().Err(err).Msg("Error getting database connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseConnError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Get salt from the database
	salt, err := idm.GetSalt(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, herror.StatusWithInfo(codes.Unauthenticated, "Wrong email or password", herror.AuthWrongCredentialsError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
		} else {
			log.Error().Err(err).Str("email", email).Msg("Error getting salt from database")
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
		}
	}

	// Hash the password
	hashedPassword := idm.PasswordHash(password, salt)

	var userId uuid.UUID
	var name string
	err = db.QueryRow(ctx, "SELECT id, name FROM users.users WHERE email = $1 AND password = $2", email, hashedPassword).Scan(&userId, &name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Debug().Str("email", email).Str("hashedPassword", hashedPassword).Msg("User not found")
			return nil, herror.StatusWithInfo(codes.Unauthenticated, "Wrong email or password", herror.AuthWrongCredentialsError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
		} else {
			log.Error().Err(err).Str("email", email).Str("hashedPassword", hashedPassword).Msg("Error getting user from database")
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
		}
	}

	// Get companies of the user
	rows, err := db.Query(ctx, "SELECT id, name FROM companies.company WHERE id  = (SELECT company_id FROM users.user_company WHERE user_id = $1)", userId)
	if err != nil {
		log.Error().Err(err).Msg("Error getting companies from database")
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}
	defer rows.Close()

	companies := []*company.Company{}
	for rows.Next() {
		var companyId uuid.UUID
		var companyName string
		err = rows.Scan(&companyId, &companyName)
		if err != nil {
			return nil, herror.StatusWithInfo(codes.Internal, "Error scanning DB rows", herror.DatabaseRowScanError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
		}
		companies = append(companies, &company.Company{Id: companyId.String(), Name: companyName})
	}

	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error starting transaction", herror.DatabaseTxError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}
	defer tx.Rollback(ctx)

	// Insert token into the database
	token := uuid.New()
	_, err = tx.Exec(ctx, "INSERT INTO users.users_session (id, user_id, expiry_date) VALUES ($1, $2, $3)", token, userId, time.Now().Add(time.Hour*72))
	if err != nil {
		log.Error().Err(err).Msg("Error creating user session")
		return nil, herror.StatusWithInfo(codes.Internal, "Error while creating user session", herror.DatabaseError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error committing transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to commit database transaction", herror.DatabaseTxError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	return &idmanagement.LoginResponse{Token: token.String(), Name: name, Email: email, Companies: companies}, nil
}

func (s *IdentityManagementServer) Register(ctx context.Context, in *idmanagement.RegisterRequest) (*emptypb.Empty, error) {
	email := in.GetEmail()
	password := in.GetPassword()
	name := in.GetName()
	timezone := in.GetTimezone()

	if email == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing email", []*errdetails.BadRequest_FieldViolation{{
			Field:       "email",
			Description: "Email is required",
		}}).Err()
	}
	if password == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing password", []*errdetails.BadRequest_FieldViolation{{
			Field:       "password",
			Description: "Password is required",
		}}).Err()
	}
	if name == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing name", []*errdetails.BadRequest_FieldViolation{{
			Field:       "name",
			Description: "Name is required",
		}}).Err()
	}
	if timezone == "" {
		timezone = "Europe/Lisbon"
	} else {
		_, err := time.LoadLocation(timezone)
		if err != nil {
			return nil, herror.StatusBadRequest(codes.InvalidArgument, "Timezone "+timezone+" is invalid", []*errdetails.BadRequest_FieldViolation{{
				Field:       "timezone",
				Description: "Timezone must be a IANA Timezone",
			}}).Err()
		}
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Error().Err(err).Msg("Error getting database connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseConnError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error starting transaction", herror.DatabaseTxError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}
	defer tx.Rollback(ctx)

	// Check if user already exists in the database
	var count int
	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM users.users WHERE email = $1", email).Scan(&count)
	if err != nil {
		log.Error().Err(err).Str("email", email).Msg("Error checking if user exists")
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}
	if count > 0 {
		return nil, herror.StatusWithInfo(codes.AlreadyExists, "User already exists", herror.UserAlreadyExists, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Insert user into the database
	salt := idm.RandomSalt()
	hashedPassword := idm.PasswordHash(password, salt)

	_, err = tx.Exec(ctx, "INSERT INTO users.users (name, email, password, salt, timezone) VALUES ($1, $2, $3, $4, $5)", name, email, hashedPassword, salt, timezone)
	if err != nil {
		log.Error().Err(err).Msg("Error creating user")
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to create user", herror.DatabaseError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error committing transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to commit database transaction", herror.DatabaseTxError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}
	return &emptypb.Empty{}, nil
}

func (s *IdentityManagementServer) Alive(ctx context.Context, in *idmanagement.TokenRequest) (*emptypb.Empty, error) {

	token := in.GetToken()
	if token == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing token", []*errdetails.BadRequest_FieldViolation{{
			Field:       "token",
			Description: "Token is required",
		}}).Err()
	}

	// Convert token to UUID
	_, err := uuid.Parse(token)
	if err != nil {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Invalid token", []*errdetails.BadRequest_FieldViolation{{
			Field:       "token",
			Description: "Token is not a valid UUID",
		}}).Err()
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Error().Err(err).Msg("Error getting database connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseConnError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Check if token exists in the database
	var expiry_date time.Time
	err = db.QueryRow(ctx, "SELECT expiry_date FROM users.users_session WHERE id = $1", token).Scan(&expiry_date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// For security reasons, we don't want to give the user any information about the token
			return nil, herror.StatusWithInfo(codes.Unauthenticated, "Invalid token", herror.AuthInvalidTokenError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
		} else {
			// Handle other errors
			log.Error().Err(err).Str("token", token).Msg("Error getting token from database")
			return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
		}
	}

	// Check if token is expired
	if expiry_date.Before(time.Now()) {
		return nil, herror.StatusWithInfo(codes.Unauthenticated, "Token expired", herror.AuthInvalidTokenError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	return &emptypb.Empty{}, nil
}

func (s *IdentityManagementServer) Logout(ctx context.Context, in *idmanagement.TokenRequest) (*emptypb.Empty, error) {
	token := in.GetToken()
	if token == "" {
		return nil, herror.StatusBadRequest(codes.InvalidArgument, "Missing token", []*errdetails.BadRequest_FieldViolation{{
			Field:       "token",
			Description: "Token is required",
		}}).Err()
	}

	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Error().Err(err).Msg("Error getting database connection")
		return nil, herror.StatusWithInfo(codes.Internal, "Database error", herror.DatabaseConnError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error starting transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Error starting transaction", herror.DatabaseTxError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}
	defer tx.Rollback(ctx)

	// Delete token from the database
	_, err = tx.Exec(ctx, "DELETE FROM users.users_session WHERE id = $1", token)
	if err != nil {
		log.Error().Err(err).Str("token", token).Msg("Error deleting token")
		return nil, herror.StatusWithInfo(codes.Internal, "Error deleting token", herror.DatabaseError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error committing transaction")
		return nil, herror.StatusWithInfo(codes.Internal, "Unable to commit database transaction", herror.DatabaseTxError, idmanagement.IdentityManagement_ServiceDesc.ServiceName, nil).Err()
	}

	return &emptypb.Empty{}, nil
}
