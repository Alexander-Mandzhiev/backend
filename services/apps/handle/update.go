package handle

import (
	app "backend/protos/gen/go/apps"
	"context"
)

func (s *serverAPI) Update(ctx context.Context, req *app.UpdateRequest) (*app.UpdateResponse, error) {
	return s.service.Update(ctx, req)
}
