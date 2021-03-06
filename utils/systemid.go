package wzlib_utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/infra-whizz/wzlib"

	wzlib_logger "github.com/infra-whizz/wzlib/logger"
)

var _MachineIDUtil *WzMachineIDUtil

type WzMachineIDUtilConsumer struct{}

// GetMAchineIdUtil returns machine ID instance
func (wzid WzMachineIDUtilConsumer) GetMachineIdUtil() *WzMachineIDUtil {
	if _MachineIDUtil == nil {
		panic("MachineID utility was not properly initialised yet")
	}

	return _MachineIDUtil
}

// WzMachineIDUtil object to keep/read/create machine-id
type WzMachineIDUtil struct {
	filePath  string
	machineid string
	wzlib_logger.WzLogger
}

// NewWzMachineIDUtil creates a new instance of an object
func newWzMachineIDUtil(filePath string) *WzMachineIDUtil {
	wmid := new(WzMachineIDUtil)
	wmid.filePath = filePath

	wmid.setupMachineId()

	return wmid
}

// SetupMachineIdUtil setting up the Utility
func (WzMachineIDUtil) SetupMachineIdUtil(filePath string) {
	if _MachineIDUtil == nil {
		_MachineIDUtil = newWzMachineIDUtil(filePath)
	}
}

// GetMachineId always returns machine-id
// If machine-id is not present, it will be copied from
// an existing one or generated, if no dbus found.
func (wmid *WzMachineIDUtil) GetMachineId() string {
	return _MachineIDUtil.machineid
}

// setupMachineId reads an existing machine ID or creates new one.
// Sequence as follows:
// 1. Read machine ID from filePath
// 2. If empty, copy from /etc/machine-id
// 3. If nothing on #2, create own one to filePath
func (wmid *WzMachineIDUtil) setupMachineId() {
	systemdMidFPath := "/etc/machine-id"
	if wmid.filePath == "" {
		wmid.filePath = systemdMidFPath
	}
	mid, err := ioutil.ReadFile(wmid.filePath)
	if err != nil {
		wmid.GetLogger().Debugf("File %s was not found", wmid.filePath)
		mid, err = ioutil.ReadFile(systemdMidFPath)
		if err != nil {
			wmid.GetLogger().Debugf("This system has no /etc/machine-id file, creating a replacement.")

			hasher := md5.New()
			_, err := io.WriteString(hasher, wzlib.MakeJid())
			if err != nil {
				panic(err)
			}
			mid = []byte(fmt.Sprintf("%x", hasher.Sum(nil)))
		}
		if wmid.filePath != systemdMidFPath {
			if err := ioutil.WriteFile(wmid.filePath, mid, 0644); err != nil {
				wmid.GetLogger().Errorf("Unable to duplicate machine id: %s", err.Error())
			}
		}
	}
	wmid.machineid = strings.TrimSpace(string(mid))
}
