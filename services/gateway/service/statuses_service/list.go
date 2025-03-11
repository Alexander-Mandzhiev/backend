package statuses_service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"log/slog"
)

func (s *Service) List(ctx context.Context, req *statuses.ListStatusesRequest) (*statuses.StatusListResponse, error) {
	op := "statuses.List"
	sl.Log.Info("Listing statuses", slog.String("op", op))
	resp, err := s.client.List(ctx, req)
	if err != nil {
		sl.Log.Error("List failed", sl.Err(err), slog.String("op", op))
		return nil, err
	}
	return resp, nil
}
