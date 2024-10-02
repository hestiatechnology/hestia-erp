package methods

import (
	"context"
	"hestia/api/pb/accounting"
	"hestia/api/utils/db"
	"hestia/api/utils/logger"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TaxServer struct {
	accounting.UnimplementedTaxServer
}

func (s *TaxServer) GetVatRates(ctx context.Context, _ *emptypb.Empty) (*accounting.VatRates, error) {
	db, err := db.GetDBPoolConn()
	if err != nil {
		logger.ErrorLogger.Println("Unable to connect to db: ", err)
		return nil, status.Error(codes.Internal, "Unable to connect to db")
	}

	rows, err := db.Query(ctx, "SELECT r.id, r.rate, t.code, r.country FROM accounting.vat_rate r LEFT JOIN accounting.vat_type t ON r.vat_type_id = t.id")
	if err != nil {
		logger.ErrorLogger.Println("Unable to query db: ", err)
		return nil, status.Error(codes.Internal, "Something went wrong")
	}

	defer rows.Close()

	var vatRates accounting.VatRates
	for rows.Next() {
		var vatRate accounting.VatRate
		// convert uuid to string
		var id uuid.UUID
		err = rows.Scan(&id, &vatRate.Rate, &vatRate.Code, &vatRate.Country)
		if err != nil {
			logger.ErrorLogger.Println("Unable to scan row: ", err)
			return nil, status.Error(codes.Internal, "Something went wrong")
		}
		vatRate.Id = id.String()
		vatRates.VatRates = append(vatRates.VatRates, &vatRate)
	}

	return &vatRates, nil
}
