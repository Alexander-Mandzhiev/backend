package helpers

import (
	"backend/pkg/server/respond"
	"fmt"
	"net/http"
	"strconv"
)

func ValidateAndParseID(idStr string, w http.ResponseWriter, r *http.Request) (int, bool) {
	if idStr == "" {
		respond.RespondedError(w, r, http.StatusBadRequest, fmt.Errorf("ID parameter is required"))
		return 0, false
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		respond.RespondedError(w, r, http.StatusBadRequest, fmt.Errorf("invalid ID parameter"))
		return 0, false
	}

	return id, true
}
