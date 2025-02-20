package handle

import (
	app "backend/protos/gen/go/apps"
	"context"
)

func (s *serverAPI) App(ctx context.Context, req *app.GetAppRequest) (*app.GetAppResponse, error) {
	return s.service.App(ctx, req)
}
