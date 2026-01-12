module github.com/ekko-earth/shared/outbox

go 1.25.4

require github.com/google/uuid v1.6.0

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/ekko-earth/shared/adapters v0.0.0
	github.com/ekko-earth/shared/gorm v0.0.0
	github.com/ekko-earth/shared/messaging v0.0.0
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gorm.io/datatypes v1.2.7
	gorm.io/driver/mysql v1.5.6 // indirect
	gorm.io/driver/postgres v1.6.0 // indirect
	gorm.io/gorm v1.31.1 // indirect
)

require golang.org/x/sys v0.39.0 // indirect

replace github.com/ekko-earth/shared/adapters => ../adapters

replace github.com/ekko-earth/shared/messaging => ../messaging

replace github.com/ekko-earth/shared/gorm => ../gorm
