package roles

import (
	"net/http"

	httperror "github.com/portainer/portainer/pkg/libhttp/error"
	"github.com/portainer/portainer/pkg/libhttp/response"
)

// @id RoleList
// @summary List roles
// @description List all roles available for use
// @description **Access policy**: administrator
// @tags roles
// @security ApiKeyAuth
// @security jwt
// @produce json
// @success 200 {array} portainer.Role "Success"
// @failure 500 "Server error"
// @router /roles [get]
func (handler *Handler) roleList(w http.ResponseWriter, r *http.Request) *httperror.HandlerError {
	roles, err := handler.DataStore.Role().ReadAll()
	if err != nil {
		return httperror.InternalServerError("Unable to retrieve authorization sets from the database", err)
	}

	return response.JSON(w, roles)
}
