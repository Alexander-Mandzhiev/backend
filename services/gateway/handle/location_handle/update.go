package location_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/locations"
	"context"
	"encoding/json"
	"net/http"
)

func (ls *LocationService) UpdateLocation(w http.ResponseWriter, r *http.Request, id int32) {
	var req locations.UpdateLocationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.RespondedError(w, r, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	req.Id = id

	resp, err := ls.client.Update(context.Background(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, resp)
}
