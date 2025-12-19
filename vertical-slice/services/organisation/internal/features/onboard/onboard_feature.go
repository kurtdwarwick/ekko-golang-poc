package onboard

import (
	"fmt"
	"log/slog"
	"organisation/internal/adapters"
	"organisation/internal/features/onboard/adapters/inbound/grpc"
	"organisation/internal/features/onboard/adapters/inbound/http"
	"os"

	organisationInMemoryAccess "organisation/internal/features/onboard/adapters/outbound/data/access/inmemory"
	organisationCommandHandlers "organisation/internal/features/onboard/core/commands/handlers"
	organisationPolicies "organisation/internal/features/onboard/core/policies"
	organisationRepositories "organisation/internal/features/onboard/core/repositories"

	"policies"

	"consumers"
)

type OnboardFeature struct {
	server consumers.Server
}

func NewOnboardFeature(server consumers.Server) *OnboardFeature {
	slog.Info("Creating policy handler")

	policyHandler := policies.NewPolicyHandler(
		organisationPolicies.LegalNameValidationPolicy{},
		organisationPolicies.TradingNameValidationPolicy{},
		organisationPolicies.WebsiteValidationPolicy{})

	slog.Info("Creating organisation DAO")

	organisationDao := organisationInMemoryAccess.NewInMemoryOrganisationDAO()

	slog.Info("Creating repository")

	repository := organisationRepositories.NewOrganisationRepository(
		organisationDao,
		*policyHandler,
	)

	slog.Info("Creating command handler")

	onboardOrganisationCommandHandler := organisationCommandHandlers.NewOnboardOrganisationCommandHandler(
		repository,
	)

	switch os.Args[1] {
	case "http":
		slog.Info("Creating HTTP consumer")

		httpServer := server.(*adapters.HttpServer)

		http.NewOnboardOrganisationHttpConsumer(
			*httpServer,
			onboardOrganisationCommandHandler,
		)
	case "grpc":
		slog.Info("Creating GRPC consumer")

		grpcServer := server.(*adapters.GrpcServer)

		grpc.NewOnboardOrganisationGrpcConsumer(
			*grpcServer,
			onboardOrganisationCommandHandler,
		)
	default:
		panic(fmt.Sprintf("invalid consumer: %s", os.Args[1]))
	}

	return &OnboardFeature{
		server: server,
	}
}

func (feature *OnboardFeature) Start() {
	feature.server.Start()
}

func (feature *OnboardFeature) Stop() {
	feature.server.Stop()
}
