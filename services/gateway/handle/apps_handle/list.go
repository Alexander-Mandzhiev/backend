package apps_handle

import (
	"backend/pkg/server/respond"
	app "backend/protos/gen/go/apps"
	"context"
	"net/http"
)

func (s *AppsService) List(w http.ResponseWriter, r *http.Request) {
	resp, err := s.client.Apps(context.Background(), &app.GetAppsRequest{})
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, resp)
}
