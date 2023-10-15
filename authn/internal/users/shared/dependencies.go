package shared

import (
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/lib/persistence"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/store"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/users/store/sqllite_ferret"
)

type Dependencies struct {
	persist   persistence.SqlLite
	UserStore store.UserStore
}

func NewDependencies(persist persistence.SqlLite) *Dependencies {

	d := &Dependencies{}

	d.persist = persist

	d.UserStore = sqllite_ferret.NewSqlLite(d.persist)

	return d

}
