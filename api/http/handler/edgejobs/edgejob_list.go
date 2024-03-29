package edgejobs

import (
	"net/http"

	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/response"
)

// @id EdgeJobList
// @summary Fetch EdgeJobs list
// @description **Access policy**: administrator
// @tags edge_jobs
// @security ApiKeyAuth
// @security jwt
// @produce json
// @success 200 {array} portainer.EdgeJob
// @failure 500
// @failure 400
// @failure 503 "Edge compute features are disabled"
// @router /edge_jobs [get]
// GET request on /api/edge_jobs
func (handler *Handler) edgeJobList(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	edgeJobs, err := handler.DataStore.EdgeJob().ReadAll()
	if err != nil {
		return httperror.InternalServerError("Unable to retrieve Edge jobs from the database", err)
	}

	return response.JSON(w, edgeJobs)
}
