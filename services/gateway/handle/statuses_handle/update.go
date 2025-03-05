package statuses_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/statuses"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *StatusesService) Update(w http.ResponseWriter, r *http.Request) {
	var req statuses.UpdateStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respond.RespondedError(w, r, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	if req.Id <= 0 {
		respond.RespondedError(w, r, http.StatusBadRequest, fmt.Errorf("ID must be provided in the request body"))
		return
	}

	res, err := s.client.Update(r.Context(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, res)
}
