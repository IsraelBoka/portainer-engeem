package endpointproxy

import (
	"github.com/gorilla/mux"
	portainer "github.com/portainer/portainer/api"
	"github.com/portainer/portainer/api/dataservices"
	"github.com/portainer/portainer/api/http/proxy"
	"github.com/portainer/portainer/api/http/security"
	httperror "github.com/portainer/portainer/pkg/libhttp/error"
)

// Handler is the HTTP handler used to proxy requests to external APIs.
type Handler struct {
	*mux.Router
	DataStore            dataservices.DataStore
	requestBouncer       security.BouncerService
	ProxyManager         *proxy.Manager
	ReverseTunnelService portainer.ReverseTunnelService
}

// NewHandler creates a handler to proxy requests to external APIs.
func NewHandler(bouncer security.BouncerService) *Handler {
	h := &Handler{
		Router:         mux.NewRouter(),
		requestBouncer: bouncer,
	}
	h.PathPrefix("/{id}/azure").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToAzureAPI)))
	h.PathPrefix("/{id}/docker").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToDockerAPI)))
	h.PathPrefix("/{id}/kubernetes").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToKubernetesAPI)))
	h.PathPrefix("/{id}/agent/docker").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToDockerAPI)))
	h.PathPrefix("/{id}/agent/kubernetes").Handler(
		bouncer.AuthenticatedAccess(httperror.LoggerHandler(h.proxyRequestsToKubernetesAPI)))
	return h
}
