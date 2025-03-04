package location_handle

import (
	"backend/pkg/helpers"
	"backend/pkg/server/respond"
	"backend/protos/gen/go/locations"
	"fmt"
	"google.golang.org/grpc"
	"net/http"
)

type LocationService struct {
	client locations.LocationServiceClient
}

func New(conn *grpc.ClientConn) *LocationService {
	return &LocationService{client: locations.NewLocationServiceClient(conn)}
}

func (s *LocationService) Locations(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/location/"):]
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
