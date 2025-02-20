package handle

import (
	app "backend/protos/gen/go/apps"
	"context"
)

func (s *serverAPI) Create(ctx context.Context, req *app.CreateRequest) (*app.CreateResponse, error) {
	return s.service.Create(ctx, req)
}
