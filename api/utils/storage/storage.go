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
		Credentials: credentials.NewStaticCredentials("--------", "-------------", ""),
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

	// Enable SSE-B2 default encryption
	_, err = svc.PutBucketEncryption(&s3.PutBucketEncryptionInput{
		Bucket: aws.String(bucketName),
		ServerSideEncryptionConfiguration: &s3.ServerSideEncryptionConfiguration{
			Rules: []*s3.ServerSideEncryptionRule{
				{
					ApplyServerSideEncryptionByDefault: &s3.ServerSideEncryptionByDefault{
						SSEAlgorithm: aws.String("AES256"),
					},
				},
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Str("name", bucketName).Msg("failed to enable SSE-B2 default encryption")
		return fmt.Errorf("failed to enable SSE-B2 default encryption: %v", err)
	}

	log.Info().Str("name", bucketName).Msg("SSE-B2 default encryption enabled")
	return nil
}

// Returns a pre-signed URL for the given object key
// Filename can be a path
func UploadFile(companyId uuid.UUID, filename string) (string, error) {
	bucketName := "hst-" + companyId.String()
	if strings.ToLower(os.Getenv("ENV")) == "dev" {
		bucketName = "hst-dev-" + bucketName
	}

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("------", "--------------------", ""),
		// Backblaze endpoint
		Endpoint: aws.String("https://s3.eu-central-003.backblazeb2.com"),
		Region:   aws.String("eu-central-003")},
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to create session")
		return "", fmt.Errorf("failed to create session: %v", err)
	}

	svc := s3.New(sess)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	})

	url, err := req.Presign(15 * 60)
	if err != nil {
		log.Error().Err(err).Str("key", filename).Msg("failed to presign URL")
		return "", fmt.Errorf("failed to presign URL: %v", err)
	}

	return url, nil
}
