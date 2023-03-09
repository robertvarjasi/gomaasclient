package entity

type Raid struct {
	ID         int           `json:"id,omitempty"`
	Name       string        `json:"name,omitempty"`
	Size       int           `json:"size,omitempty"`
	UUID       string        `json:"uuid,omitempty"`
	Devices    []BlockDevice `json:"devices,omitempty"`
	Incomplete bool          `json:"__incomplete__,omitempty"`
}

type RaidParams struct {
	Name                  string   `url:"name,omitempty"`
	AddBlockDevices       []string `url:"add_block_devices,omitempty"`
	AddPartitions         []string `url:"add_partitions,omitempty"`
	AddSpareDevices       []string `url:"add_spare_devices,omitempty"`
	AddSparePartitions    []string `url:"add_spare_partitions,omitempty"`
	RemoveBlockDevices    []string `url:"remove_block_devices,omitempty"`
	RemovePartitions      []string `url:"remove_partitions,omitempty"`
	RemoveSpareDevices    []string `url:"remove_spare_devices,omitempty"`
	RemoveSparePartitions []string `url:"remove_spare_partitions,omitempty"`
	UUID                  string   `url:"uuid,omitempty"`
}

type RaidsParams struct {
	Name            string   `url:"name,omitempty"`
	BlockDevices    []string `url:"block_devices,omitempty"`
	Level           string   `url:"level,omitempty"`
	Partitions      []string `url:"partitions,omitempty"`
	SpareDevices    []string `url:"spare_devices,omitempty"`
	SparePartitions []string `url:"spare_partitions,omitempty"`
}
