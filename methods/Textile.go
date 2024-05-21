package methods

import (
	"context"
	"hestia/api/pb/hestia/textile"
)

type TextileServer struct {
	textile.UnimplementedTextileServer
}

func (s *TextileServer) CreateTechnicalModel(ctx context.Context, in *textile.TechnicalModel) (*textile.TechnicalModel, error) {
	return &textile.TechnicalModel{}, nil
}
