package file

import (
	"context"
	"hestia/api/pb/file"
	"hestia/api/utils/herror"
	"hestia/api/utils/storage"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
)

type FileServer struct {
	file.UnimplementedFileServer
}

func (s *FileServer) UploadFile(ctx context.Context, in *file.FileUpload) (*file.PresignedURL, error) {
	filename := in.GetFileName()

	url, err := storage.UploadFile(uuid.UUID{}, filename)
	if err != nil {
		log.Error().Err(err).Msg("failed to upload file")
		return nil, herror.StatusWithInfo(codes.Internal, "failed to generate presigned url", herror.PresignedURL, file.File_ServiceDesc.ServiceName, nil).Err()
	}

	return &file.PresignedURL{Url: url}, nil
}
