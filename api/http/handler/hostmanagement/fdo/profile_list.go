package fdo

import (
	"net/http"

	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/response"
)

// @id fdoProfileList
// @summary retrieves all FDO profiles
// @description retrieves all FDO profiles
// @description **Access policy**: administrator
// @tags intel
// @security jwt
// @produce json
// @success 200 "Success"
// @failure 400 "Invalid request"
// @failure 403 "Permission denied to access settings"
// @failure 500 "Server error"
// @failure 500 "Bad gateway"
// @router /fdo/profiles [get]
func (handler *Handler) fdoProfileList(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {

	profiles, err := handler.DataStore.FDOProfile().ReadAll()
	if err != nil {
		return httperror.InternalServerError(err.Error(), err)
	}

	return response.JSON(w, profiles)
}
