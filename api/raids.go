package api

import (
	"github.com/maas/gomaasclient/entity"
)

type RAIDs interface {
	Get(systemID string) (raids []entity.Raid, err error)
	Create(systemID string, params *entity.RaidsParams) (raid *entity.Raid, err error)
}
