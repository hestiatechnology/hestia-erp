package accounting

import (
	"context"
	"hestia/api/pb/accounting"
	"hestia/api/utils/db"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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
		log.Error().Err(err).Msg("Unable to connect to db")
		return nil, status.Error(codes.Internal, "Unable to connect to db")
	}

	rows, err := db.Query(ctx, "SELECT r.id, r.rate, t.code, r.country FROM accounting.vat_rate r LEFT JOIN accounting.vat_type t ON r.vat_type_id = t.id WHERE (r.start_date < NOW() AND (r.end_date IS NULL OR r.end_date > NOW()))")
	if err != nil {
		log.Error().Err(err).Msg("Unable to query db")
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
			log.Error().Err(err).Msg("Unable to scan row")
			return nil, status.Error(codes.Internal, "Something went wrong")
		}
		vatRate.Id = id.String()
		vatRates.VatRates = append(vatRates.VatRates, &vatRate)
	}

	return &vatRates, nil
}

func (s *TaxServer) GetVatExemptions(ctx context.Context, _ *emptypb.Empty) (*accounting.VatExemptions, error) {
	db, err := db.GetDBPoolConn()
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to db")
		return nil, status.Error(codes.Internal, "Unable to connect to db")
	}

	rows, err := db.Query(ctx, "SELECT id, code, description FROM accounting.vat_exemption WHERE (start_date < NOW() AND (end_date IS NULL OR end_date > NOW()))")
	if err != nil {
		log.Error().Err(err).Msg("Unable to query db")
		return nil, status.Error(codes.Internal, "Something went wrong")
	}

	defer rows.Close()

	var vatExemptions accounting.VatExemptions
	for rows.Next() {
		var vatExemption accounting.VatExemption
		// convert uuid to string
		var id uuid.UUID
		err = rows.Scan(&id, &vatExemption.Code, &vatExemption.Description)
		if err != nil {
			log.Error().Err(err).Msg("Unable to scan row")
			return nil, status.Error(codes.Internal, "Something went wrong")
		}
		vatExemption.Id = id.String()
		vatExemptions.VatExemptions = append(vatExemptions.VatExemptions, &vatExemption)
	}

	return &vatExemptions, nil
}
