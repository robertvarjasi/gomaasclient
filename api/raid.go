package api

import (
	"github.com/maas/gomaasclient/entity"
)

type RAID interface {
	Get(systemID string, id int) (raid *entity.Raid, err error)
	Update(systemID string, id int, params *entity.RaidParams) (raid *entity.Raid, err error)
	Delete(systemID string, id int) (err error)
}
