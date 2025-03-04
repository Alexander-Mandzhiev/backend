package apps_handle

import (
	"backend/pkg/server/respond"
	app "backend/protos/gen/go/apps"
	"context"
	"net/http"
)

func (s *AppsService) Get(w http.ResponseWriter, r *http.Request, id int32) {
	resp, err := s.client.App(context.Background(), &app.GetAppRequest{Id: id})
	if err != nil {
		respond.RespondedError(w, r, http.StatusNotFound, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, resp)
}
