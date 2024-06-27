package server

import (
	"context"
	rest_api "kardinal.kontrol/kardinal-manager/api/http_rest/server"
	rest_types "kardinal.kontrol/kardinal-manager/api/http_rest/types"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// List virtual services
// (GET /virtual-services)
func (Server) GetVirtualServices(ctx context.Context, object rest_api.GetVirtualServicesRequestObject) (rest_api.GetVirtualServicesResponseObject, error) {
	response := map[string]rest_types.VirtualService{
		"fake-virtual-service-01": {
			Name: "fake-virtual-service-01",
		},
		"fake-virtual-service-02": {
			Name: "fake-virtual-service-02",
		},
	}

	return rest_api.GetVirtualServices200JSONResponse(response), nil
}

// Create virtual service
// (POST /virtual-services)
func (Server) PostVirtualServices(ctx context.Context, object rest_api.PostVirtualServicesRequestObject) (rest_api.PostVirtualServicesResponseObject, error) {
	response := rest_types.VirtualService{
		Name: "fake-virtual-service-01",
	}

	return rest_api.PostVirtualServices200JSONResponse(response), nil
}

// Delete virtual service
// (DELETE /virtual-services)
func (Server) DeleteVirtualServices(ctx context.Context, object rest_api.DeleteVirtualServicesRequestObject) (rest_api.DeleteVirtualServicesResponseObject, error) {
	response := rest_types.VirtualService{
		Name: "fake-virtual-service-01",
	}

	return rest_api.DeleteVirtualServices200JSONResponse(response), nil
}
