package statuses_handle

import (
	"backend/pkg/helpers"
	"backend/pkg/server/respond"
	"backend/protos/gen/go/statuses"
	"fmt"
	"google.golang.org/grpc"
	"net/http"
)

type StatusesService struct {
	client statuses.StatusServiceClient
}

func New(conn *grpc.ClientConn) *StatusesService {
	return &StatusesService{client: statuses.NewStatusServiceClient(conn)}
}

func (s *StatusesService) Statuses(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/statuses/"):]
	switch r.Method {
	case http.MethodGet:
		if idStr == "" {
			s.List(w, r)
		} else {
			id, isValid := helpers.ValidateAndParseID(idStr, w, r)
			if !isValid {
				return
			}
			s.Get(w, r, int32(id))
		}
	case http.MethodPost:
		s.Create(w, r)
	case http.MethodPut:
		s.Update(w, r)
	case http.MethodDelete:
		id, isValid := helpers.ValidateAndParseID(idStr, w, r)
		if !isValid {
			return
		}
		s.Delete(w, r, int32(id))
	default:
		w.Header().Set("Allow", "GET, POST, PUT, DELETE")
		respond.RespondedError(w, r, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
	}
}
