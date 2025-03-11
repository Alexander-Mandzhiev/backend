package service

import (
	sl "backend/pkg/logger"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (s *Service) RecordOutMsSQL(ctx context.Context, ids []int64) error {
	if err := s.productionTaskProvider.RecordedOutMsSQL(ctx, ids); err != nil {
		sl.Log.Error("Failed recording out MsSQL", slog.String("error", err.Error()))
		return status.Errorf(codes.Internal, "failed recording out MsSQL: %v", err)
	}
	return nil
}
