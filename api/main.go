package main

import (
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"hestia/api/interceptor"
	maccounting "hestia/api/methods/accounting"
	mcompany "hestia/api/methods/company"
	"hestia/api/methods/idm"
	mtextile "hestia/api/methods/textile"
	"hestia/api/pb/accounting"
	"hestia/api/pb/company"
	"hestia/api/pb/idmanagement"
	"hestia/api/pb/textile"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Caller().Logger()
	// Check if env is set to dev or args --dev is passed
	if strings.ToLower(os.Getenv("ENV")) == "dev" || len(os.Args) > 1 && os.Args[1] == "--dev" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	PORT := 9000
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}
	log.Info().Int("port", PORT).Msg("Server listening")
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(interceptor.AuthInterceptor))

	if strings.ToLower(os.Getenv("ENV")) == "dev" || len(os.Args) > 1 && os.Args[1] == "--dev" {
		log.Info().Msg("Running in development mode")
		log.Info().Msg("Registering reflection service")
		reflection.Register(s)
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		log.Info().Msg("Shutting down gracefully...")
		s.GracefulStop()
		log.Info().Msg("Server stopped")
	}()

	// Service registration
	idmanagement.RegisterIdentityManagementServer(s, &idm.IdentityManagementServer{})
	textile.RegisterTextileServer(s, &mtextile.TextileServer{})
	company.RegisterCompanyManagementServer(s, &mcompany.CompanyManagementServer{})
	accounting.RegisterTaxServer(s, &maccounting.TaxServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("failed to serve")
	}
}
