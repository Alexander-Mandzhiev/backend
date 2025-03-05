package statuses_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/statuses"
	"net/http"
)

func (s *StatusesService) Delete(w http.ResponseWriter, r *http.Request, id int32) {
	req := statuses.DeleteStatusRequest{Id: id}
	res, err := s.client.Delete(r.Context(), &req)
	if err != nil {
		respond.RespondedError(w, r, http.StatusInternalServerError, err)
		return
	}

	respond.Respond(w, r, http.StatusOK, res)
}
