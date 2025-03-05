package statuses_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/statuses"
	"net/http"
)

func (s *StatusesService) List(w http.ResponseWriter, r *http.Request) {
	req := statuses.ListStatusesRequest{}
	res, err := s.client.List(r.Context(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, res)
}
