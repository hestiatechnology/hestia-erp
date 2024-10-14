package herror

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Database Errors
const (
	DatabaseError        = "DB_ERROR"
	DatabaseConnError    = "DB_CONN_ERROR"
	DatabaseTxError      = "DB_TX_ERROR"
	DatabaseRowScanError = "DB_ROW_SCAN_ERROR"
)

// Authentication Errors
const (
	AuthWrongCredentialsError = "WRONG_CREDENTIALS"
	AuthMissingMetadataError  = "MISSING_METADATA"
	AuthMissingTokenError     = "MISSING_TOKEN"
	AuthInvalidTokenError     = "INVALID_TOKEN"
)

// Authorization Errors
const (
	AuthMissingRoleError = "MISSING_ROLE"
	AuthInvalidRoleError = "INVALID_ROLE"
)

// User Errors
const (
	UserAlreadyExists = "USER_ALREADY_EXISTS"
	InvalidUser       = "INVALID_USER"
)

// Company Errors
const (
	CompanyAlreadyExists = "COMPANY_ALREADY_EXISTS"
	InvalidCompany       = "INVALID_COMPANY"
)

// Employee Errors
const (
	EmployeeAlreadyExists = "EMPLOYEE_ALREADY_EXISTS"
	InvalidEmployee       = "INVALID_EMPLOYEE"
)

func StatusWithInfo(code codes.Code, message string, reason string, domain string, metadata map[string]string) *status.Status {
	st := status.New(code, message)
	ds, err := st.WithDetails(&errdetails.ErrorInfo{
		Reason:   reason,
		Domain:   domain,
		Metadata: metadata,
	})
	if err != nil {
		return st
	}
	return ds
}

func StatusBadRequest(code codes.Code, message string, fieldViolations []*errdetails.BadRequest_FieldViolation) *status.Status {
	st := status.New(code, message)
	ds, err := st.WithDetails(&errdetails.BadRequest{
		FieldViolations: fieldViolations,
	})
	if err != nil {
		return st
	}
	return ds
}
