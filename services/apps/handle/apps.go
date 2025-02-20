package handle

import (
	app "backend/protos/gen/go/apps"
	"context"
)

func (s *serverAPI) Apps(ctx context.Context, req *app.GetAppsRequest) (*app.GetAppsResponse, error) {
	return s.service.Apps(ctx)
}
