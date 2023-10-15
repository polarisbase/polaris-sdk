package shared

import (
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/lib/persistence"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/sessions/store"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/sessions/store/sqllite_ferret"
)

type Dependencies struct {
	persist      persistence.SqlLite
	SessionStore store.SessionStore
}

func NewDependencies(persist persistence.SqlLite) *Dependencies {

	d := &Dependencies{}

	d.persist = persist

	d.SessionStore = sqllite_ferret.NewSqlLite(d.persist)

	return d

}
