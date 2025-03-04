package location_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/locations"
	"context"
	"net/http"
)

func (s *LocationService) Get(w http.ResponseWriter, r *http.Request, id int32) {
	resp, err := s.client.Location(context.Background(), &locations.GetLocationRequest{Id: id})
	if err != nil {
		respond.RespondedError(w, r, http.StatusNotFound, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, resp)
}
