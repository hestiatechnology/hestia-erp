package user

import (
	"context"
	"hestia/api/utils/db"
	"time"
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
