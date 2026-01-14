package onboard

import (
	"context"
	"log/slog"

	"github.com/ekko-earth/organisation/internal/features/onboard/adapters/http"
	"github.com/ekko-earth/shared/outbox"
	"github.com/ekko-earth/shared/policies"

	"github.com/ekko-earth/organisation/internal/features/onboard/core"

	adapters "github.com/ekko-earth/shared/adapters"
	messagingAdapters "github.com/ekko-earth/shared/messaging/adapters"

	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
	httpAdapters "github.com/ekko-earth/shared/http/adapters"

	organisationGormAccess "github.com/ekko-earth/organisation/internal/features/onboard/adapters/gorm"
	organisationPolicies "github.com/ekko-earth/organisation/internal/features/onboard/core/policies"
)

type OnboardFeature struct {
	outboundMessagePublisher messagingAdapters.MessagePublisher
	server                   adapters.Server
	database                 adapters.Database
}

func NewOnboardFeature(
	outboundMessagePublisher messagingAdapters.MessagePublisher,
	server adapters.Server,
	database adapters.Database,
	outboxRepository *outbox.OutboxRepository,
	unitOfWork adapters.UnitOfWork,
) *OnboardFeature {
	policyHandler := policies.NewPolicyHandler(
		organisationPolicies.LegalNameValidationPolicy{},
		organisationPolicies.TradingNameValidationPolicy{},
		organisationPolicies.WebsiteValidationPolicy{})

	slog.Info("Creating organisation DAO")

	gormDatabase := database.(*gormAdapters.GormDatabase)

	organisationDao := organisationGormAccess.NewGormOrganizationDAO(*gormDatabase)

	slog.Info("Creating repository")

	repository := core.NewOrganisationRepository(
		organisationDao,
		*policyHandler,
	)

	slog.Info("Creating command handler")

	onboardOrganisationCommandHandler := core.NewOnboardOrganisationCommandHandler(
		repository,
		unitOfWork,
		outboundMessagePublisher,
		outboxRepository,
	)

	slog.Info("Creating HTTP consumer")

	httpServer := server.(*httpAdapters.HttpServer)

	http.NewOnboardOrganisationHttpConsumer(
		httpServer,
		onboardOrganisationCommandHandler,
	)

	return &OnboardFeature{
		server:   server,
		database: database,
	}
}

func (feature *OnboardFeature) Start(ctx context.Context) error {
	return nil
}

func (feature *OnboardFeature) Stop(ctx context.Context) error {
	return nil
}
