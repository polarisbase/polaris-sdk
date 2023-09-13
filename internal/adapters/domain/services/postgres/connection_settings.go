package postgres

import "fmt"

type ConnectionSettings struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Ssl      bool
}

// NewConnectionSettings creates a new ConnectionSettings.
func NewConnectionSettings(host string, port int, username string, password string, database string, ssl bool) *ConnectionSettings {
	return &ConnectionSettings{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
		Ssl:      ssl,
	}
}

// PgxConnectionString returns a connection string for pgx.
func (c *ConnectionSettings) PgxConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.Username,
		c.Password, c.Host, c.Port, c.Database)
}
