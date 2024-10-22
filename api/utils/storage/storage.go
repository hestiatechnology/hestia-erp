package storage

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func CreateS3Bucket(companyId uuid.UUID) error {
	bucketName := "hst-" + companyId.String()
	if strings.ToLower(os.Getenv("ENV")) == "dev" {
		bucketName = "hst-dev-" + bucketName
	}

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("003dfe653b1e4290000000001", "K003if50xGYcjc6SnsPqKoNYs2HE+I8", ""),
		// Backblaze endpoint
		Endpoint: aws.String("https://s3.eu-central-003.backblazeb2.com"),
		Region:   aws.String("eu-central-003")},
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to create session")
		return fmt.Errorf("failed to create session: %v", err)
	}

	svc := s3.New(sess)

	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Error().Err(err).Str("name", bucketName).Msg("failed to create bucket")
		return fmt.Errorf("failed to create bucket: %v", err)
	}

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Error().Err(err).Str("name", bucketName).Msg("error occurred while waiting for bucket to be created")
		return fmt.Errorf("error occurred while waiting for bucket to be created: %v", err)
	}

	log.Info().Str("name", bucketName).Msg("bucket created")

	// Apply CORS rules
	corsConfiguration := &s3.CORSConfiguration{
		CORSRules: []*s3.CORSRule{
			{
				AllowedHeaders: []*string{
					aws.String("*"),
				},
				AllowedMethods: []*string{
					aws.String("GET"),
					aws.String("POST"),
					aws.String("PUT"),
					aws.String("DELETE"),
				},
				AllowedOrigins: []*string{
					aws.String("*"),
				},
				MaxAgeSeconds: aws.Int64(3000),
			},
		},
	}

	_, err = svc.PutBucketCors(&s3.PutBucketCorsInput{
		Bucket:            aws.String(bucketName),
		CORSConfiguration: corsConfiguration,
	})
	if err != nil {
		log.Error().Err(err).Str("name", bucketName).Msg("failed to apply CORS rules")
		return fmt.Errorf("failed to apply CORS rules: %v", err)
	}

	log.Info().Str("name", bucketName).Msg("CORS rules applied")
	return nil
}
