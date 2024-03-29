package teammemberships

import (
	"net/http"

	"github.com/portainer/portainer/api/dataservices"
	"github.com/portainer/portainer/api/http/security"
	httperror "github.com/portainer/portainer/pkg/libhttp/error"

	"github.com/gorilla/mux"
)

// Handler is the HTTP handler used to handle team membership operations.
type Handler struct {
	*mux.Router
	DataStore dataservices.DataStore
}

// NewHandler creates a handler to manage team membership operations.
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}

	h.Use(bouncer.TeamLeaderAccess)

	h.Handle("/team_memberships", httperror.LoggerHandler(h.teamMembershipCreate)).Methods(http.MethodPost)
	h.Handle("/team_memberships", httperror.LoggerHandler(h.teamMembershipList)).Methods(http.MethodGet)
	h.Handle("/team_memberships/{id}", httperror.LoggerHandler(h.teamMembershipUpdate)).Methods(http.MethodPut)
	h.Handle("/team_memberships/{id}", httperror.LoggerHandler(h.teamMembershipDelete)).Methods(http.MethodDelete)

	return h
}
