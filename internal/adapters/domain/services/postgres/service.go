package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services"
)

type Service struct {
	*services.ServiceBase[common.OptionServicePostgres]
	connectionSettings *ConnectionSettings
	pgxc               *pgx.Conn
}

// Connect connects to the database.
func (s *Service) Connect() error {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), s.connectionSettings.PgxConnectionString())
	if err != nil {
		if shouldContinue, err := s.OnInternalError(err); !shouldContinue {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}

		return err
	}

	// Test the connection.
	if err := conn.Ping(context.Background()); err != nil {
		if shouldContinue, err := s.OnInternalError(err); !shouldContinue {
			panic(err)
		}
		return err
	}

	// Set the connection.
	s.pgxc = conn

	return nil
}

// ApplyLocalDefaults applies the local defaults to the service.
func (s *Service) ApplyLocalDefaults() error {

	// Set default connection settings.
	s.connectionSettings = &ConnectionSettings{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "postgres",
		Database: "postgres",
		Ssl:      false,
	}

	return nil
}

// New creates a new Service.
func New(applicationName string, options ...common.Option) *Service {
	// Create the service.
	s := &Service{}

	// Apply local the defaults.
	if err := s.ApplyLocalDefaults(); err != nil {
		panic(err)
	}

	// Set the base service
	s.ServiceBase = services.NewBase[common.OptionServicePostgres](common.DB_POSTGRES_SERVICE, s)

	// Apply options
	s.ServiceBase.ApplyOptions(options...)

	// Set Startup Entrypoint
	s.SetServiceStartupEntrypoint(func() error {
		return s.Connect()
	})

	return s
}
