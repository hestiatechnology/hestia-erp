package methods

import (
	"context"
	"errors"
	"log"
	"time"

	"hestia/api/pb"
	"hestia/api/utils/db"
	"hestia/api/utils/idm"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type IdentityManagementServer struct {
	pb.UnimplementedIdentityManagementServer
}

func (s *IdentityManagementServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	if in.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing email")
	}
	if in.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing password")
	} else if len(in.GetPassword()) != 64 {
		return nil, status.Error(codes.InvalidArgument, "Invalid password")
	}

	//Check if user exists in the database
	db, err := db.GetDbPoolConn()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Get salt from the database
	salt, err := idm.GetSalt(ctx, in.GetEmail())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Wrong email or password")
		} else {
			log.Println(err)
			return nil, status.Error(codes.Internal, "Database error")
		}
	}

	// Hash the password
	hashedPassword := idm.PasswordHash(in.GetPassword(), salt)

	var userId uuid.UUID
	var name string
	err = db.QueryRow(ctx, "SELECT id, name FROM users.users WHERE email = $1 AND password = $2", in.GetEmail(), hashedPassword).Scan(&userId, &name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "Wrong email or password")
		} else {
			log.Println(err)
			return nil, status.Error(codes.Internal, "Database error")
		}
	}

	// Get companies of the user
	rows, err := db.Query(ctx, "SELECT id, name FROM companies.company WHERE id  = (SELECT company_id FROM users.user_company WHERE user_id = $1)", userId)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}
	defer rows.Close()

	companies := []*pb.CompanyList{}
	for rows.Next() {
		var companyId uuid.UUID
		var companyName string
		err = rows.Scan(&companyId, &companyName)
		if err != nil {
			log.Println(err)
			return nil, status.Error(codes.Internal, "Database error")
		}
		companies = append(companies, &pb.CompanyList{Id: companyId.String(), Name: companyName})
	}

	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}
	defer tx.Rollback(ctx)

	// Insert token into the database
	token := uuid.New()
	_, err = tx.Exec(ctx, "INSERT INTO users.users_session (id, user_id, expiry_date) VALUES ($1, $2, $3)", token, userId, time.Now().Add(time.Hour*72))
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	return &pb.LoginResponse{Token: token.String(), Name: name, Email: in.GetEmail(), Companies: companies}, nil
}

func (s *IdentityManagementServer) Register(ctx context.Context, in *pb.RegisterRequest) (*emptypb.Empty, error) {
	if in.GetEmail() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing email")
	}
	if in.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing password")
	}
	if in.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing name")
	}
	if in.GetTimezone() == "" {
		return nil, status.Error(codes.InvalidArgument, "Missing timezone")
	} else {
		_, err := time.LoadLocation(in.GetTimezone())
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Timezone "+in.GetTimezone()+" is invalid")
		}
	}

	db, err := db.GetDbPoolConn()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}
	defer tx.Rollback(ctx)

	// Check if user already exists in the database
	var count int
	err = tx.QueryRow(ctx, "SELECT COUNT(*) FROM users.users WHERE email = $1", in.GetEmail()).Scan(&count)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}
	if count > 0 {
		return nil, status.Error(codes.AlreadyExists, "User already exists")
	}

	// Insert user into the database
	_, err = tx.Exec(ctx, "INSERT INTO users.users (email, password, name, timezone) VALUES ($1, $2, $3, $4)", in.GetEmail(), in.GetPassword(), in.GetName(), in.GetTimezone())
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}
	return &emptypb.Empty{}, nil
}

func (s *IdentityManagementServer) Alive(ctx context.Context, in *pb.TokenRequest) (*emptypb.Empty, error) {

	token := in.GetToken()
	if token == "" {
		return nil, status.Error(codes.Unauthenticated, "Missing token")
	}

	db, err := db.GetDbPoolConn()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Check if token exists in the database
	var expiry_date time.Time
	err = db.QueryRow(ctx, "SELECT expiry_date FROM users.users_session WHERE id = $1", token).Scan(&expiry_date)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// For security reasons, we don't want to give the user any information about the token
			return nil, status.Error(codes.Unauthenticated, "Token expired")
		} else {
			// Handle other errors
			log.Println(err)
			return nil, status.Error(codes.Internal, "Database error")
		}
	}

	// Check if token is expired
	if expiry_date.Before(time.Now()) {
		return nil, status.Error(codes.Unauthenticated, "Token expired")
	}

	return &emptypb.Empty{}, nil
}

func (s *IdentityManagementServer) Logout(ctx context.Context, in *pb.TokenRequest) (*emptypb.Empty, error) {
	token := in.GetToken()
	if token == "" {
		return nil, status.Error(codes.Unauthenticated, "Missing token")
	}

	db, err := db.GetDbPoolConn()
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Start a transaction
	tx, err := db.Begin(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}
	defer tx.Rollback(ctx)

	// Delete token from the database
	_, err = tx.Exec(ctx, "DELETE FROM users.users_session WHERE id = $1", token)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	// Commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		log.Println(err)
		return nil, status.Error(codes.Internal, "Database error")
	}

	return &emptypb.Empty{}, nil
}
