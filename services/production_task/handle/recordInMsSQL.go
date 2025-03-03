package handle

import (
	"backend/protos/gen/go/production_task"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) RecordInMsSQL(ctx context.Context, req *production_task.IDsRequest) (*production_task.EmptyResponse, error) {
	err := s.service.RecordInMsSQL(ctx, req.Ids)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to record in MsSQL: %v", err)
	}
	return &production_task.EmptyResponse{}, nil
}
