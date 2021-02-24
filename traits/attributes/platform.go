package wzlib_traits_attributes

import (
	"fmt"
	"strings"

	"github.com/elastic/go-sysinfo"
	"github.com/elastic/go-sysinfo/types"
	wzlib_logger "github.com/infra-whizz/wzlib/logger"
	wzlib_traits "github.com/infra-whizz/wzlib/traits"
	wzlib_utils "github.com/infra-whizz/wzlib/utils"
	"golang.org/x/sys/unix"
)

// SysInfo class
type SysInfo struct {
	host types.Host
	wzlib_logger.WzLogger
}

// NewSysInfo is a constructor to SysInfo class
func NewSysInfo() *SysInfo {
	si := new(SysInfo)

	var err error
	si.host, err = sysinfo.Host()
	if err != nil {
		si.GetLogger().Error(fmt.Sprintf("SysInfo error: %s", err.Error()))
	}
	return si
}

// Load all sub-attributes
func (si *SysInfo) Load(container *wzlib_traits.WzTraitsContainer) {
	if si.host != nil {
		si.memory(container)
		si.info(container)
		si.osInfo(container)
		si.syscallInfo(container)
	}
}

func (si *SysInfo) memory(c *wzlib_traits.WzTraitsContainer) {
	meminfo, err := si.host.Memory()
	if err != nil {
		si.GetLogger().Errorf("Unable to detect memory data: %s", err.Error())
	} else {
		c.Set("memory.total", meminfo.Total)
		c.Set("memory.vtotal", meminfo.VirtualTotal)
	}
}

func (si *SysInfo) info(c *wzlib_traits.WzTraitsContainer) {
	nfo := si.host.Info()
	c.Set("arch", nfo.Architecture)
	c.Set("container", nfo.Containerized)
	c.Set("hostname", nfo.Hostname)
	c.Set("kernel_version", nfo.KernelVersion)
	c.Set("uid", nfo.UniqueID)
}

func (si *SysInfo) osInfo(c *wzlib_traits.WzTraitsContainer) {
	nfo := si.host.Info().OS
	if nfo != nil {
		c.Set("os.build", nfo.Build)
		c.Set("os.codename", nfo.Codename)
		c.Set("os.family", nfo.Family)
		c.Set("os.ver", nfo.Version)
		c.Set("os.version", nfo.Version) // alias to "ver"
		c.Set("os.ver_major", nfo.Major)
		c.Set("os.ver_minor", nfo.Minor)
		c.Set("os.ver_patch", nfo.Patch)
		c.Set("os.name", nfo.Name)
		c.Set("os.platform", nfo.Platform)
	}
}

// This won't work on MacOS, because Uname is not a part of a system.
func (si *SysInfo) syscallInfo(c *wzlib_traits.WzTraitsContainer) {
	var uts unix.Utsname
	if err := unix.Uname(&uts); err == nil {
		c.Set("os.sysname", strings.ToLower(wzlib_utils.Ba65Str(uts.Sysname)))
	}
}
