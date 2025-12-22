package organisation

import (
	adapters "github.com/ekko-earth/shared/adapters"
	messagingAdapters "github.com/ekko-earth/shared/messaging/adapters"

	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
	rabbitmqAdapters "github.com/ekko-earth/shared/rabbitmq/adapters"

	impactGormAccess "github.com/ekko-earth/impact/internal/organisation/adapters/gorm"
	impactRabbitmqAdapters "github.com/ekko-earth/impact/internal/organisation/adapters/rabbitmq"
	impactEventHandlers "github.com/ekko-earth/impact/internal/organisation/core/events/handlers"
	impactRepositories "github.com/ekko-earth/impact/internal/organisation/core/repositories"
)

type OrganisationDomain struct{}

func NewOrganisationDomain(
	database adapters.Database,
	inboundMessageBus messagingAdapters.MessageBus,
) *OrganisationDomain {
	gormDatabase := database.(*gormAdapters.GormDatabase)

	organisationDAO := impactGormAccess.NewGormOrganizationDAO(*gormDatabase)

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
