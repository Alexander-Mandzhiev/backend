package service

import (
	sl "backend/pkg/logger"
	"backend/protos/gen/go/statuses"
	"context"
	"errors"
	"log/slog"
)

func (s *Service) Update(ctx context.Context, request *statuses.UpdateStatusRequest) (*statuses.StatusResponse, error) {
	op := "service.Update"
	id := int(request.Id)
	sl.Log.Debug("Updating status", slog.Int("id", id), slog.String("op", op))

	existingStatus, err := s.statusesProvider.Status(ctx, id)
	if err != nil {
		if errors.Is(err, ErrStatusNotFound) {
			sl.Log.Warn("Status not found for update", slog.Int("id", id), slog.String("op", op))
		} else {
			sl.Log.Error("Failed to fetch status for update", sl.Err(err), slog.Int("id", id), slog.String("op", op))
		}
		return nil, err
	}

	existingStatus.Name = request.Name
	existingStatus.Description = request.Description

	if err = s.statusesProvider.Update(ctx, existingStatus); err != nil {
		sl.Log.Error("Failed to update status", sl.Err(err), slog.Int("id", id), slog.String("op", op))
		return nil, err
	}

	sl.Log.Info("Status updated successfully", slog.Int("id", int(existingStatus.Id)), slog.String("op", op))
	return existingStatus, nil
}
