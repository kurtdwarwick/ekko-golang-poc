package adapters

import (
	"fmt"
	"log/slog"

	"github.com/ekko-earth/shared/adapters"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GormDatabase struct {
	configuration adapters.DatabaseConfiguration

	Database *gorm.DB
}

func NewGormDatabase(configuration adapters.DatabaseConfiguration) *GormDatabase {
	config := &gorm.Config{}

	if configuration.Schema != "" {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix: configuration.Schema + ".",
		}
	}

	host := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s  sslmode=disable search_path=%s",
		configuration.Host,
		configuration.Port,
		configuration.Username,
		configuration.Password,
		configuration.Database,
		configuration.Schema,
	)

	slog.Info("Connecting to PostgreSQL", "host", host)

	database, err := gorm.Open(postgres.Open(host), config)

	if err != nil {
		panic(err)
	}

	return &GormDatabase{configuration: configuration, Database: database}
}

func (database *GormDatabase) Connect() error {
	return nil
}

func (database *GormDatabase) Disconnect() error {
	return nil
}
