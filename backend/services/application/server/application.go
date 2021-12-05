package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	ApplicationM "github.com/jalexanderII/solid-pancake/clients/application/models"
	"github.com/jalexanderII/solid-pancake/database"
	applicationpb "github.com/jalexanderII/solid-pancake/gen/application"
	"github.com/jalexanderII/solid-pancake/gen/common"
	ApplicationH "github.com/jalexanderII/solid-pancake/services/application/handlers"
	RealEstateM "github.com/jalexanderII/solid-pancake/services/realestate/models"
	UserM "github.com/jalexanderII/solid-pancake/services/users/models"
)

// ApplicationServiceServer is a gRPC server it implements the methods defined by the ApplicationServer interface
type ApplicationServiceServer struct {
	log hclog.Logger
	h   *ApplicationH.Handler
}

func toRequestModel(req *applicationpb.ApplicationReq) ApplicationM.ApplicantFormRequest {
	var (
		apartment RealEstateM.Apartment
		user      UserM.User
	)
	database.Database.Db.First(&apartment, req.ApartmentRef)
	database.Database.Db.First(&user, req.UserRef)
	return ApplicationM.ApplicantFormRequest{
		Name:           req.Name,
		UserRef:        int(req.UserRef),
		User:           user,
		SocialSecurity: req.SocialSecurity,
		DateOfBirth:    req.DateOfBirth,
		DriversLicense: req.DriversLicense,
		PreviousAddress: RealEstateM.Place{
			Address:      req.PreviousAddress.Address,
			Street:       req.PreviousAddress.Street,
			City:         req.PreviousAddress.City,
			State:        req.PreviousAddress.State,
			Zip:          req.PreviousAddress.Zip,
			Neighborhood: req.PreviousAddress.Neighborhood,
			Unit:         req.PreviousAddress.Unit,
			Lat:          req.PreviousAddress.Lat,
			Lng:          req.PreviousAddress.Lng,
		},
		Landlord:       req.Landlord,
		LandlordNumber: req.LandlordNumber,
		Employer:       req.Employer,
		Salary:         req.Salary,
		ApartmentRef:   int(req.ApartmentRef),
		Apartment:      apartment,
	}
}

// NewApplicationServer creates a new Application server
func NewApplicationServer(log hclog.Logger, h *ApplicationH.Handler) *ApplicationServiceServer {
	return &ApplicationServiceServer{log, h}
}

func (a *ApplicationServiceServer) Apply(_ context.Context, req *applicationpb.ApplicationReq) (*applicationpb.ApplicationRes, error) {
	a.log.Info("Handle Apply", "user_id", req.GetUserRef())
	return a.h.Apply(toRequestModel(req))
}

func (a *ApplicationServiceServer) ReadApplicationRequest(_ context.Context, req *common.ID) (*applicationpb.ApplicationReq, error) {
	a.log.Info("Handle GetApplicationById", "user_id", req.String())
	return a.h.GetApplication(req.GetId())
}

func (a *ApplicationServiceServer) DeleteApplicationRequest(_ context.Context, req *common.ID) (*applicationpb.ApplicationReq, error) {
	a.log.Info("Handle DeleteApplicationRequest", "user_id", req.String())
	return a.h.DeleteApplication(req.GetId())
}

func (a *ApplicationServiceServer) ListApplicationRequests(context.Context, *common.Empty) (*applicationpb.ListApplicationReqOut, error) {
	a.log.Info("Handle ListApplicationRequests")
	return a.h.GetApplications()
}

func (a *ApplicationServiceServer) ReadApplicationResponse(_ context.Context, req *common.ID) (*applicationpb.ApplicationRes, error) {
	a.log.Info("Handle ReadApplicationResponse", "user_id", req.String())
	return a.h.GetApplicationResponse(req.GetId())
}

func (a *ApplicationServiceServer) DeleteApplicationResponse(_ context.Context, req *common.ID) (*applicationpb.ApplicationRes, error) {
	a.log.Info("Handle DeleteApplicationResponse", "user_id", req.String())
	return a.h.DeleteApplicationResponse(req.GetId())
}
