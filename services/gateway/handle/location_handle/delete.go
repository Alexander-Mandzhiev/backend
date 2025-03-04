package location_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/locations"
	"context"
	"net/http"
)

func (s *LocationService) Delete(w http.ResponseWriter, r *http.Request, id int32) {
	resp, err := s.client.Delete(context.Background(), &locations.DeleteLocationRequest{Id: id})
	if err != nil {
		respond.RespondedError(w, r, http.StatusNotFound, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, resp)
}
