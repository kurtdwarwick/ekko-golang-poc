package query

import (
	"context"

	"github.com/ekko-earth/organisation/internal/features/query/adapters/http"
	"github.com/ekko-earth/organisation/internal/features/query/core/repositories"
	"github.com/ekko-earth/shared/adapters"

	organisationGormAccess "github.com/ekko-earth/organisation/internal/features/query/adapters/gorm"
	queries "github.com/ekko-earth/organisation/internal/features/query/core/queries/handlers"

	gormAdapters "github.com/ekko-earth/shared/gorm/adapters"
	httpAdapters "github.com/ekko-earth/shared/http/adapters"
)

type QueryFeature struct {
	server   adapters.Server
	database adapters.Database
}

func NewQueryFeature(
	server adapters.Server,
	database adapters.Database,
) *QueryFeature {
	gormDatabase := database.(*gormAdapters.GormDatabase)

	organisationDao := organisationGormAccess.NewGormOrganizationDAO(*gormDatabase)

	repository := repositories.NewOrganisationRepository(
		organisationDao,
	)

	getOrganisationByIdQueryHandler := queries.NewGetOrganisationByIdQueryHandler(
		repository,
	)

	httpServer := server.(*httpAdapters.HttpServer)

	// getOrganisationByIdHttpDtoMessageTranslator := http.GetOrganisationByIdHttpDtoMessageTranslator{}

	// httpAdapters.NewHttpConsumer(
	// 	*httpServer,
	// 	&getOrganisationByIdHttpDtoMessageTranslator,
	// 	getOrganisationByIdQueryHandler,
	// 	httpAdapters.HttpConsumerConfiguration{
	// 		Route:   "/organisations/{id}",
	// 		Methods: []string{"GET"},
	// 	},
	// )

	http.NewGetOrganisationByIdHttpConsumer(
		httpServer,
		*getOrganisationByIdQueryHandler,
	)

	return &QueryFeature{
		server:   server,
		database: database,
	}
}

func (feature *QueryFeature) Start(ctx context.Context) error {
	return nil
}

func (feature *QueryFeature) Stop(ctx context.Context) error {
	return nil
}
