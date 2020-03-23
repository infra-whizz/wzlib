package wzlib_traits_attributes

import (
	"fmt"
	"sort"

	wzlib_traits "github.com/infra-whizz/wzlib/traits"
	"github.com/shirou/gopsutil/disk"
)

// DiskInfo class
type DiskInfo struct {
	KEY_PREFIX string
}

// NewDiskInfo is a constructor to SysInfo class
func NewDiskInfo() *DiskInfo {
	di := new(DiskInfo)
	di.KEY_PREFIX = "media"
	return di
}

// Load all sub-attributes
func (di *DiskInfo) Load(container *wzlib_traits.WzTraitsContainer) {
	deviceNames := di.physical(container)
	container.Set(fmt.Sprintf("%s.devices", di.KEY_PREFIX), deviceNames)
}

func (di *DiskInfo) physical(c *wzlib_traits.WzTraitsContainer) []string {
	names := make(map[string]interface{})
	devices, _ := disk.Partitions(false)
	for _, device := range devices {
		serial := disk.GetDiskSerialNumber(device.Device)
		if serial != "" {
			c.Set(fmt.Sprintf("%s.%s.serial", di.KEY_PREFIX, device.Device), serial)
			names[device.Device] = nil
		}
	}

	devNames := make([]string, 0)
	for name := range names {
		devNames = append(devNames, name)
	}
	sort.Strings(devNames)
	return devNames
}
