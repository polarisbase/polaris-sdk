package store

import (
	"context"
	"github.com/polarisbase/polaris-sdk/v3/services/templates/template_1/internal/info/model"
)

type InfoStore interface {
	List(ctx context.Context, limit int, offset int) (infos []model.Info, err error, ok bool)
	CreateInfo(ctx context.Context, infoIn model.Info) (info model.Info, err error, ok bool)
}
