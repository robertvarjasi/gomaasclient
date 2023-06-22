package client

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

type RAID struct {
	ApiClient ApiClient
}

func (b *RAID) client(systemID string, id int) ApiClient {
	return b.ApiClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("raid").GetSubObject(fmt.Sprintf("%v", id))
}

func (b *RAID) Get(systemID string, id int) (raid *entity.Raid, err error) {
	raid = new(entity.Raid)
	err = b.client(systemID, id).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, raid)
	})
	return
}

func (b *RAID) Update(systemID string, id int, params *entity.RaidParams) (raid *entity.Raid, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	raid = new(entity.Raid)
	err = b.client(systemID, id).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, raid)
	})
	return
}

func (b *RAID) Delete(systemID string, id int) error {
	return b.client(systemID, id).Delete()
}
