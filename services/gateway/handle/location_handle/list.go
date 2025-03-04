package location_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/locations"
	"context"
	"net/http"
)

func (ls *LocationService) ListLocations(w http.ResponseWriter, r *http.Request) {
	resp, err := ls.client.List(context.Background(), &locations.ListLocationsRequest{})
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, resp)
}
