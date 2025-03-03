package handle

import (
	"backend/pkg/server/respond"
	"net/http"
)

func (h *Handler) test(w http.ResponseWriter, r *http.Request) {
	respond.Respond(w, r, http.StatusOK, "ok")
}
