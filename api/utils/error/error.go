package error

// Database Errors
const (
	DatabaseError        = "DB_ERROR"
	DatabaseConnError    = "DB_CONN_ERROR"
	DatabaseTxError      = "DB_TX_ERROR"
	DatabaseGenericError = "DB_GENERIC_ERROR"
)

// Authentication Errors
const (
	AuthMissingMetadataError = "MISSING_METADATA"
	AuthMissingTokenError    = "MISSING_TOKEN"
	AuthInvalidTokenError    = "INVALID_TOKEN"
)

// Authorization Errors
const (
	AuthMissingRoleError = "MISSING_ROLE"
	AuthInvalidRoleError = "INVALID_ROLE"
)
