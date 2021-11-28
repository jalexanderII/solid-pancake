package server

import (
	"context"

	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	"github.com/jalexanderII/solid-pancake/gen/common"
	"github.com/jalexanderII/solid-pancake/services/application/handlers"
)

// ApplicationServiceServer is a gRPC server it implements the methods defined by the ApplicationServer interface
type ApplicationServiceServer struct {
	h *handlers.Handler
}

// NewApplicationServer creates a new Application server
func NewApplicationServer(h *handlers.Handler) *ApplicationServiceServer {
	return &ApplicationServiceServer{h}
}

func (a *ApplicationServiceServer) Apply(context.Context, *applicationpb.ApplicationReq) (*applicationpb.ApplicationRes, error){
	return nil, nil
}

func (a *ApplicationServiceServer) ReadApplicationRequest(context.Context, *common.ID) (*applicationpb.ApplicationReq, error){
	return nil, nil
}

func (a *ApplicationServiceServer) UpdateApplicationRequest(context.Context, *applicationpb.ApplicationReq) (*applicationpb.ApplicationReq, error){
	return nil, nil
}

func (a *ApplicationServiceServer) DeleteApplicationRequest(context.Context, *common.ID) (*applicationpb.ApplicationReq, error){
	return nil, nil
}

func (a *ApplicationServiceServer) ListApplicationRequests(context.Context, *common.Empty) (*applicationpb.ListApplicationReqOut, error){
	return nil, nil
}

func (a *ApplicationServiceServer) ReadApplicationResponse(context.Context, *common.ID) (*applicationpb.ApplicationRes, error){
	return nil, nil
}

func (a *ApplicationServiceServer) DeleteApplicationResponse(context.Context, *common.ID) (*applicationpb.ApplicationRes, error){
	return nil, nil
}

func (a *ApplicationServiceServer) ListApplicationResponse(context.Context, *common.Empty) (*applicationpb.ListApplicationResOut, error){
	return nil, nil
}

