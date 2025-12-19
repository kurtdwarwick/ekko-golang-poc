package grpc

import (
	"net"

	"organisation/internal/adapters"
	organisationCommands "organisation/internal/features/onboard/core/commands"
	organisationCommandHandlers "organisation/internal/features/onboard/core/commands/handlers"

	"context"

	"organisation/internal/features/onboard/adapters/inbound/grpc/proto"
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
	server adapters.GrpcServer,
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
	context context.Context,
	request *proto.OnboardOrganisationRequest,
) (*proto.OnboardOrganisationResponse, error) {
	command := organisationCommands.OnboardOrganisationCommand{
		LegalName:   *request.LegalName,
		TradingName: *request.TradingName,
		Website:     request.Website,
	}

	organisationId, err := consumer.onboardOrganisationCommandHandler.Handle(command)

	return &proto.OnboardOrganisationResponse{
		Id: organisationId,
	}, err
}
