package pboptions

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/common"
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/postgres"
)

var PostgresServiceOptions PostgresServiceOption = PostgresServiceOption{}

type PostgresServiceOption struct{}

func (PostgresServiceOption) SetPostgresConnection(connection string) common.OptionServicePostgres {
	return postgres.NewOption("set-postgres-connection", func(obj interface{}) error {
		if service, ok := obj.(*postgres.Service); ok {
			// service.SetPostgresConnection(connection)
			_ = service
		}
		return nil
	})
}
