package client

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"net/url"

	"github.com/maas/gomaasclient/entity"
)

type RAIDs struct {
	ApiClient ApiClient
}

func (r *RAIDs) client(systemID string) ApiClient {
	return r.ApiClient.GetSubObject("nodes").GetSubObject(systemID).GetSubObject("raids")
}

func (r *RAIDs) Get(systemID string) (raids []entity.Raid, err error) {
	err = r.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, &raids)
	})
	return
}

func (r *RAIDs) Create(systemID string, params *entity.RaidsParams) (raid *entity.Raid, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	raid = new(entity.Raid)
	err = r.client(systemID).Post("", qsp, func(data []byte) error {
		return json.Unmarshal(data, raid)
	})
	return
}
