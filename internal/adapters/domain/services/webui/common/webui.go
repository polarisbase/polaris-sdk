package common

import (
	"github.com/polarisbase/polaris-sdk/internal/adapters/domain/services/api"
)

type WebUI interface {
	Bootstrap(api *api.Service)
	AddPage(page Page)
}