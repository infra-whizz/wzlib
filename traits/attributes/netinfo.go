package wzlib_traits_attributes

import (
	"fmt"

	"github.com/elastic/go-sysinfo"
	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	wzlib_traits "github.com/infra-whizz/wzlib/traits"
	"github.com/shirou/gopsutil/net"
)

// NetInfo class
type NetInfo struct {
	KEY_PREFIX string
	container  *wzlib_traits.WzTraitsContainer
	wzlib_logger.WzLogger
}

// NewNetInfo is a constructor to SysInfo class
func NewNetInfo() *NetInfo {
	ni := new(NetInfo)
	ni.KEY_PREFIX = "net"
	return ni
}

// Load all sub-attributes
func (ni *NetInfo) Load(container *wzlib_traits.WzTraitsContainer) {
	ni.container = container
	ni.interfaces()
	ni.network()
}

func (ni *NetInfo) network() {
	host, err := sysinfo.Host()
	if err != nil {
		ni.GetLogger().Errorln("Error getting sysinfo for the host:", err.Error())
	}
	nfo := host.Info()
	ni.container.Set(fmt.Sprintf("%s.MAC", ni.KEY_PREFIX), nfo.MACs)
	ni.container.Set(fmt.Sprintf("%s.IP", ni.KEY_PREFIX), nfo.IPs)
}

func (ni *NetInfo) interfaces() {
	ifaces, err := net.Interfaces()
	if err != nil {
		ni.GetLogger().Errorln("Error getting network interface:", err.Error())
		return
	}

	for _, iface := range ifaces {
		addresses := make([]string, 0)
		for _, addr := range iface.Addrs {
			addresses = append(addresses, addr.Addr)
		}
		if len(addresses) > 0 && iface.HardwareAddr != "" {
			ni.container.Set(fmt.Sprintf("%s.%s.addr", ni.KEY_PREFIX, iface.Name), addresses)
			ni.container.Set(fmt.Sprintf("%s.%s.hw_addr", ni.KEY_PREFIX, iface.Name), iface.HardwareAddr)
			ni.container.Set(fmt.Sprintf("%s.%s.flags", ni.KEY_PREFIX, iface.Name), iface.Flags)
		}
	}
}
