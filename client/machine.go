package client

import (
	"encoding/json"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
	"github.com/maas/gomaasclient/entity"
)

// Contains functionality for manipulating the Machine entity.
type Machine struct {
	ApiClient ApiClient
}

func (m *Machine) client(systemID string) ApiClient {
	return m.ApiClient.GetSubObject("machines").GetSubObject(systemID)
}

// Get machine details.
func (m *Machine) Get(systemID string) (ma *entity.Machine, err error) {
	ma = new(entity.Machine)
	err = m.client(systemID).Get("", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Update machine.
func (m *Machine) Update(systemID string, machineParams *entity.MachineParams, powerParams map[string]string) (ma *entity.Machine, err error) {
	qsp, err := query.Values(machineParams)
	if err != nil {
		return
	}
	for k, v := range powerParams {
		qsp.Add(k, v)
	}
	ma = new(entity.Machine)
	err = m.client(systemID).Put(qsp, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Delete machine.
func (m *Machine) Delete(systemID string) error {
	return m.client(systemID).Delete()
}

// Commission machine.
func (m *Machine) Commission(systemID string, params *entity.MachineCommissionParams) (ma *entity.Machine, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	ma = new(entity.Machine)
	err = m.client(systemID).Post("commission", qsp, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Deploy machine.
func (m *Machine) Deploy(systemID string, params *entity.MachineDeployParams) (ma *entity.Machine, err error) {
	qsp, err := query.Values(params)
	if err != nil {
		return
	}
	ma = new(entity.Machine)
	err = m.client(systemID).Post("deploy", qsp, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Lock machine.
func (m *Machine) Lock(systemID string, comment string) (ma *entity.Machine, err error) {
	qsp := make(url.Values)
	if comment != "" {
		qsp.Set("comment", comment)
	}
	ma = new(entity.Machine)
	err = m.client(systemID).Post("lock", qsp, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

// Clear default gateways.
func (m *Machine) ClearDefaultGateways(systemID string) (ma *entity.Machine, err error) {
	ma = new(entity.Machine)
	err = m.client(systemID).Post("clear_default_gateways", url.Values{}, func(data []byte) error {
		return json.Unmarshal(data, ma)
	})
	return
}

func (m *Machine) GetPowerParameters(systemID string) (params map[string]string, err error) {
	params_ := map[string]interface{}{}
	err = m.client(systemID).Get("power_parameters", url.Values{}, func(data []byte) error {
		err_ := json.Unmarshal(data, &params_)
		params = make(map[string]string, len(params_))
		for k,v := range params_ {
			if k == "workaround_flags" {
				wfs := make([]string, len(v.([]interface{})))
				for k_,v_ := range v.([]interface{}) {
					wfs[k_] = v.(string)
				}
				params[k] = strings.Join(wfs, ",")
			} else {
				params[k] = v.(string)
			}
		}
		return err_
	})
	return
}
