package endpoints

import (
	"fmt"
	"net/http"

	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/request"
)

// DEPRECATED
func (handler *Handler) endpointStatusInspect(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	endpointID, err := request.RetrieveNumericRouteVariableValue(r, "id")
	if err != nil {
		return httperror.BadRequest("Invalid environment identifier route variable", err)
	}

	url := fmt.Sprintf("/api/endpoints/%d/edge/status", endpointID)
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
	return nil
}
