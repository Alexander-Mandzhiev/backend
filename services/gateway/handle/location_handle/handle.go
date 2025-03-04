package location_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/locations"
	"fmt"
	"google.golang.org/grpc"
	"net/http"
	"strconv"
)

type LocationService struct {
	client locations.LocationServiceClient
}

func New(conn *grpc.ClientConn) *LocationService {
	return &LocationService{client: locations.NewLocationServiceClient(conn)}
}

func (ls *LocationService) HandleLocation(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/v1/location/"):]
	if idStr == "" {
		switch r.Method {
		case http.MethodPost:
			ls.CreateLocation(w, r)
		case http.MethodGet:
			ls.ListLocations(w, r)
		default:
			w.Header().Set("Allow", "GET, POST")
			respond.RespondedError(w, r, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
		}
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		respond.RespondedError(w, r, http.StatusBadRequest, fmt.Errorf("invalid ID parameter"))
		return
	}

	switch r.Method {
	case http.MethodGet:
		ls.GetLocation(w, r, int32(id))
	case http.MethodPut:
		ls.UpdateLocation(w, r, int32(id))
	case http.MethodDelete:
		ls.DeleteLocation(w, r, int32(id))
	default:
		w.Header().Set("Allow", "GET, PUT, DELETE")
		respond.RespondedError(w, r, http.StatusMethodNotAllowed, fmt.Errorf("method not allowed"))
	}
}
