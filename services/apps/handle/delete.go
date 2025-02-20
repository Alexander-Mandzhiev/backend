package handle

import (
	app "backend/protos/gen/go/apps"
	"context"
)

func (s *serverAPI) Delete(ctx context.Context, req *app.DeleteRequest) (*app.DeleteResponse, error) {
	return s.service.Delete(ctx, req)
}
