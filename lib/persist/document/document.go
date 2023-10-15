package document

import "github.com/polarisbase/polaris-sdk/v3/lib/persist"

type Store interface {
	persist.Persist
}
