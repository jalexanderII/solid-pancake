package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	"github.com/jalexanderII/solid-pancake/gen/common"
)

// ApplicationServiceServer is a gRPC server it implements the methods defined by the ApplicationServer interface
type ApplicationServiceServer struct {
	log   hclog.Logger
}

// NewApplicationServer creates a new Application server
func NewApplicationServer(log hclog.Logger) *ApplicationServiceServer {
	return &ApplicationServiceServer{log}
}

func (a *ApplicationServiceServer) Apply(ctx context.Context, req *applicationpb.ApplicationReq) (*applicationpb.ApplicationRes, error){
	a.log.Info("Handle Apply", "user_id", req.GetUserRef())
	return nil, nil
}

func (a *ApplicationServiceServer) ReadApplicationRequest(ctx context.Context, req *common.ID) (*applicationpb.ApplicationReq, error){
	a.log.Info("Handle GetApplicationById", "user_id", req.String())
	return nil, nil
}

func (a *ApplicationServiceServer) UpdateApplicationRequest(ctx context.Context, req *applicationpb.ApplicationReq) (*applicationpb.ApplicationReq, error){
	a.log.Info("Handle UpdateApplicationRequest", "user_id", req.GetUserRef())
	return nil, nil
}

func (a *ApplicationServiceServer) DeleteApplicationRequest(ctx context.Context, req *common.ID) (*applicationpb.ApplicationReq, error){
	a.log.Info("Handle DeleteApplicationRequest", "user_id", req.String())
	return nil, nil
}

func (a *ApplicationServiceServer) ListApplicationRequests(ctx context.Context, _ *common.Empty) (*applicationpb.ListApplicationReqOut, error){
	a.log.Info("Handle ListApplicationRequests")
	return nil, nil
}

func (a *ApplicationServiceServer) ReadApplicationResponse(ctx context.Context, req *common.ID) (*applicationpb.ApplicationRes, error){
	a.log.Info("Handle ReadApplicationResponse", "user_id", req.String())
	return nil, nil
}

func (a *ApplicationServiceServer) DeleteApplicationResponse(ctx context.Context, req *common.ID) (*applicationpb.ApplicationRes, error){
	a.log.Info("Handle DeleteApplicationResponse", "user_id", req.String())
	return nil, nil
}

func (a *ApplicationServiceServer) ListApplicationResponse(ctx context.Context, _ *common.Empty) (*applicationpb.ListApplicationResOut, error){
	a.log.Info("Handle ListApplicationResponses")
	return nil, nil
}

