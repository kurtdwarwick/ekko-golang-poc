package organisation

import (
	adapters "github.com/ekko-earth/shared/adapters"
	messagingAdapters "github.com/ekko-earth/shared/messaging/adapters"

	mongoAdapters "github.com/ekko-earth/shared/mongodb/adapters"
	rabbitmqAdapters "github.com/ekko-earth/shared/rabbitmq/adapters"

	impactMongoAccess "github.com/ekko-earth/impact/internal/organisation/adapters/mongodb"
	impactRabbitmqAdapters "github.com/ekko-earth/impact/internal/organisation/adapters/rabbitmq"
	impactEventHandlers "github.com/ekko-earth/impact/internal/organisation/core/events/handlers"
	impactRepositories "github.com/ekko-earth/impact/internal/organisation/core/repositories"
)

type OrganisationDomain struct{}

func NewOrganisationDomain(
	database adapters.Database,
	inboundMessageBus messagingAdapters.MessageBus,
) *OrganisationDomain {
	//gormDatabase := database.(*gormAdapters.GormDatabase)
	mongoDatabase := database.(*mongoAdapters.MongoDatabase)
	//organisationDAO := impactGormAccess.NewGormOrganizationDAO(*gormDatabase)
	organisationDAO := impactMongoAccess.NewMongoDBOrganisationDAO(*mongoDatabase)

	repository := impactRepositories.NewOrganisationRepository(
		organisationDAO,
	)

	organisationOnboardedEventHandler := impactEventHandlers.NewOrganisationOnboardedEventHandler(
		repository,
	)

	inboundRabbitMQMessageBus := inboundMessageBus.(*rabbitmqAdapters.RabbitMQMessageBus)

	organisationOnboardedEventMessageTranslator := impactRabbitmqAdapters.OrganisationOnboardedEventMessageTranslator{}

	rabbitmqAdapters.NewRabbitMQMessageConsumer(
		*inboundRabbitMQMessageBus,
		&organisationOnboardedEventMessageTranslator,
		organisationOnboardedEventHandler,
		rabbitmqAdapters.RabbitMQMessageConsumerConfiguration{
			Queue:           "organisation.onboarded",
			AutoAcknowledge: true,
		},
	)

	return &OrganisationDomain{}
}

func (domain *OrganisationDomain) Start() {
}

func (domain *OrganisationDomain) Stop() {
}
