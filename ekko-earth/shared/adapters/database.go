package adapters

type DatabaseConfiguration struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Schema   string
}

type Database interface {
	Connect() error
	Disconnect() error
}
