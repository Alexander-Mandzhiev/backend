package apps_handle

import (
	"backend/pkg/helpers"
	"backend/pkg/server/respond"
	app "backend/protos/gen/go/apps"
	"fmt"
	"google.golang.org/grpc"
	"net/http"
)

type AppsService struct {
	client app.AppProviderServiceClient
}

func New(conn *grpc.ClientConn) *AppsService {
	return &AppsService{client: app.NewAppProviderServiceClient(conn)}
}

func (s *AppsService) Apps(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/apps/"):]
	switch r.Method {
	case http.MethodPost:
		s.Create(w, r)
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
