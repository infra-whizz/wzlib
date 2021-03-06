package wzlib_utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/google/uuid"
	wzlib_logger "github.com/infra-whizz/wzlib/logger"
)

// WzContainerParam object
type WzContainerParam struct {
	Root    string
	Command string
	Args    []string
}

// WzContainer struct
type WzContainer struct {
	conf *WzContainerParam
	self string
	jid  string

	wzlib_logger.WzLogger
}

// NewWzContainer class
func NewWzContainer(conf *WzContainerParam) *WzContainer {
	c := new(WzContainer)
	c.conf = conf
	c.self = "/proc/self/exe"
	c.jid = strings.Split(uuid.New().String(), "-")[0]

	return c
}

// Run a command in its own container
func (c *WzContainer) Run() (string, string, error) {
	stdout, stderr, err := c.start()
	if err != nil {
		return "", "", err
	}
	return strings.TrimSpace(stdout), strings.TrimSpace(stderr), nil
}

// ParsePkArgs parses packed ":" separated args
// TODO: probably move to base64 instead
func (c *WzContainer) ParsePkArgs(pa string) *WzContainerParam {

	return nil
}

// Start command in a container
func (c *WzContainer) start() (string, string, error) {
	c.GetLogger().Debugf("Starting nanocontainer at PID %d", os.Getpid())
	var ob, eb bytes.Buffer

	packedArgs := fmt.Sprintf(":c-%s:%s:%s:%s", c.jid, c.conf.Root, c.conf.Command, strings.Join(c.conf.Args, ","))
	cmd := exec.Command(c.self, "local", packedArgs)
	cmd.Stdout = io.MultiWriter(&ob)
	cmd.Stderr = io.MultiWriter(&eb)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	if err := cmd.Run(); err != nil {
		c.GetLogger().Info(ob.String())
		c.GetLogger().Error(eb.String())
		c.GetLogger().Errorf("Error running container: %s", err.Error())
		return "", "", err
	}

	return ob.String(), eb.String(), nil
}

func (c *WzContainer) Container() (string, string, error) {
	fmt.Printf("Running in new UTS namespace %v as %d\n", os.Args[2:], os.Getpid())

	if err := c.cGroups(); err != nil {
		return "", "", err
	}
	c.GetLogger().Debugln("Cgroups are ready")

	if err := syscall.Chroot(c.conf.Root); err != nil {
		return "", "", err
	}
	c.GetLogger().Debugf("Chrooted to '%s'", c.conf.Root)

	if err := syscall.Chdir("/"); err != nil {
		return "", "", err
	}
	c.GetLogger().Debugf("At the root of the inner filesystem")

	proc := true
	if err := syscall.Mount("proc", "proc", "proc", 0, ""); err != nil {
		proc = false
		c.GetLogger().Warnf("Unable to mound proc: %s", err.Error())
	}

	var ob, eb bytes.Buffer
	cmd := exec.Command(c.conf.Command, c.conf.Args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = io.MultiWriter(&ob)
	cmd.Stderr = io.MultiWriter(&eb)

	if err := cmd.Run(); err != nil {
		c.GetLogger().Errorf("Error running inner container process: %s", err.Error())
		return "", "", err
	}

	if proc {
		if err := syscall.Unmount("/proc", 0); err != nil {
			return "", "", nil
		}
	}

	c.GetLogger().Infof(">> STDOUT:\n%s\n", ob.String())
	c.GetLogger().Infof(">> STDERR:\n%s\n", eb.String())

	return ob.String(), eb.String(), nil
}

func (c *WzContainer) cGroups() error {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	cname := "waka-container"

	cgroupsPath := filepath.Join(pids, cname)
	if _, err := os.Stat(cgroupsPath); os.IsNotExist(err) {
		if err := os.Mkdir(cgroupsPath, 0755); err != nil {
			return err
		}
	} else {
		c.GetLogger().Debugf("Path for cgroups already exists: %s", cgroupsPath)
	}

	// fill in pids etc
	if err := ioutil.WriteFile(filepath.Join(pids, cname+"/pids.max"), []byte("10"), 0700); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(pids, cname+"/notify_on_release"), []byte("1"), 0700); err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(pids, cname+"/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700)
}
