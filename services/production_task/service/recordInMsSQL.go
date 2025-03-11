package service

import (
	sl "backend/pkg/logger"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) RecordInMsSQL(ctx context.Context, ids []int64) error {
	if err := s.productionTaskProvider.RecordedInMsSQL(ctx, ids); err != nil {
		sl.Log.Error("Failed recording in MsSQL", slog.String("error", err.Error()))
		return status.Errorf(codes.Internal, "failed recording in MsSQL: %v", err)
	}
	return nil
}
