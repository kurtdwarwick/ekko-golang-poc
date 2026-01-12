package grpc

import (
	"net"

	organisationCommands "github.com/ekko-earth/organisation/internal/features/onboard/core/commands"
	organisationCommandHandlers "github.com/ekko-earth/organisation/internal/features/onboard/core/commands/handlers"
	grpcAdapters "github.com/ekko-earth/shared/grpc/adapters"

	"context"

	"github.com/ekko-earth/organisation/internal/features/onboard/adapters/grpc/proto"
)

type OnboardOrganisationGrpcConsumer struct {
	proto.UnimplementedOnboardOrganisationServiceServer

	listener                          net.Listener
	onboardOrganisationCommandHandler *organisationCommandHandlers.OnboardOrganisationCommandHandler
}

type OnboardOrganisationGrpcServerConfiguration struct {
	Address string
	Port    string
	Network string
}

func NewOnboardOrganisationGrpcConsumer(
	server grpcAdapters.GrpcServer,
	onboardOrganisationCommandHandler *organisationCommandHandlers.OnboardOrganisationCommandHandler,
) *OnboardOrganisationGrpcConsumer {
	consumer := &OnboardOrganisationGrpcConsumer{
		listener:                          server.Listener,
		onboardOrganisationCommandHandler: onboardOrganisationCommandHandler,
	}

	proto.RegisterOnboardOrganisationServiceServer(server.Server, consumer)

	return consumer
}

func (consumer *OnboardOrganisationGrpcConsumer) OnboardOrganisation(
	ctx context.Context,
	request *proto.OnboardOrganisationRequest,
) (*proto.OnboardOrganisationResponse, error) {
	command := organisationCommands.OnboardOrganisationCommand{
		LegalName:   *request.LegalName,
		TradingName: *request.TradingName,
		Website:     request.Website,
	}

	organisationId, err := consumer.onboardOrganisationCommandHandler.Handle(command, ctx)
	organisationIdString := organisationId.String()

	return &proto.OnboardOrganisationResponse{
		Id: &organisationIdString,
	}, err
}
