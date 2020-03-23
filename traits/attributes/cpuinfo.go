package wzlib_traits_attributes

import (
	"fmt"

	wzlib_traits "github.com/infra-whizz/wzlib/traits"
	"github.com/shirou/gopsutil/cpu"
)

// CPUInfo class
type CPUInfo struct {
}

// NewCPUInfo is a constructor to SysInfo class
func NewCPUInfo() *CPUInfo {
	ci := new(CPUInfo)
	return ci
}

// Load all sub-attributes
func (ci *CPUInfo) Load(container *wzlib_traits.WzTraitsContainer) {
	ci.inspect(container)
}

func (ci *CPUInfo) inspect(c *wzlib_traits.WzTraitsContainer) {
	nfos, _ := cpu.Info()
	c.Set("cpu.count", len(nfos))
	for idx, nfo := range nfos {
		c.Set(fmt.Sprintf("cpu.%d.cores", idx), nfo.Cores)
		c.Set(fmt.Sprintf("cpu.%d.core_id", idx), nfo.CoreID)
		c.Set(fmt.Sprintf("cpu.%d.family", idx), nfo.Family)
		c.Set(fmt.Sprintf("cpu.%d.flags", idx), nfo.Flags)
		c.Set(fmt.Sprintf("cpu.%d.Mhz", idx), nfo.Mhz)
		c.Set(fmt.Sprintf("cpu.%d.microcode", idx), nfo.Microcode)
		c.Set(fmt.Sprintf("cpu.%d.model", idx), nfo.Model)
		c.Set(fmt.Sprintf("cpu.%d.model_name", idx), nfo.ModelName)
		c.Set(fmt.Sprintf("cpu.%d.id", idx), nfo.PhysicalID)
		c.Set(fmt.Sprintf("cpu.%d.physical_id", idx), nfo.PhysicalID)
		c.Set(fmt.Sprintf("cpu.%d.vendor_id", idx), nfo.VendorID)
		c.Set(fmt.Sprintf("cpu.%d.stepping", idx), nfo.Stepping)
		c.Set(fmt.Sprintf("cpu.%d.cache_size", idx), nfo.CacheSize)
		c.Set(fmt.Sprintf("cpu.%d", idx), nfo.CPU)
	}
}
