package statuses_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/statuses"
	"net/http"
)

func (s *StatusesService) Get(w http.ResponseWriter, r *http.Request, id int32) {
	req := statuses.GetStatusRequest{Id: id}
	res, err := s.client.Status(r.Context(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusNotFound, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, res)
}
