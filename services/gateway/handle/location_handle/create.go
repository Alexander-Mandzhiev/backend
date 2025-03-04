package location_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/locations"
	"context"
	"encoding/json"
	"net/http"
)

func (ls *LocationService) CreateLocation(w http.ResponseWriter, r *http.Request) {
	var req locations.CreateLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.RespondedError(w, r, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	resp, err := ls.client.Create(context.Background(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusCreated, resp)
}
