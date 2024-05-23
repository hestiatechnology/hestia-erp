package user

import (
	"context"
	"hestia/api/utils/db"
	"time"

	"github.com/google/uuid"
)

func VerifyAuthToken(ctx context.Context, token string) bool {
	db, err := db.GetDbPoolConn()
	if err != nil {
		return false
	}

	var expiry_date time.Time
	err = db.QueryRow(ctx, "SELECT expiry_date FROM users.users_session WHERE id = $1", token).Scan(&expiry_date)
	if err != nil {
		return false
	}

	if expiry_date.Before(time.Now()) {
		return false
	}

	return true
}

func IsEmployeeIdUsed(ctx context.Context, employeeId uuid.UUID) bool {
	db, err := db.GetDbPoolConn()
	if err != nil {
		return false
	}

	var count int
	err = db.QueryRow(ctx, "SELECT COUNT(*) FROM users.user_company WHERE employee_id = $1", employeeId).Scan(&count)
	if err != nil {
		return false
	}

	if count > 0 {
		return true
	}

	return false
}
