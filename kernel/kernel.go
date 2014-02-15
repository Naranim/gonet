package kernel

import (
	"errors"
	"fmt"
	"github.com/Naranim/gonet/network"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	binaryName = "netkit-kernel"
)

const (
	ALREADY_RUNNING = "network running"
	NOT_RUNNING     = "network not running"
)

var (
	NETKIT_PATH,
	KERNEL_MODULES,
	KERNEL_PATH string
)

var (
	kernelStringFlags = []string{
		"name",
		"modules",
		"title",
		"umid",
		"ubd0",
		"def_route",
		"mem",
		"root",
		"uml_dir",
		"hosthome",
		"exec",
		"hostlab",
		"hostwd",
		"con0",
		"con1",
		"xterm",
	}

	kernelBoolFlags = []string{
		"quiet",
	}
)

func init() {
	NETKIT_PATH = os.Getenv("NETKIT_HOME")
	if NETKIT_PATH == "" {
		panic("Omg, bbq")
	}
	if _, err := os.Stat(NETKIT_PATH); err != nil {
		panic(err)
	}

	KERNEL_PATH = path.Join(NETKIT_PATH, binaryName)
	if _, err := os.Stat(KERNEL_PATH); err != nil {
		panic(err)
	}
}

type Kernel struct {
	running bool
	path    string
	bashrc  string
	strVal  map[string]string
	bVal    map[string]bool
}

func flagSupported(flagName string, flagList []string) bool {
	for _, name := range flagList {
		if name == flagName {
			return true
		}
	}

	return false
}

func (k *Kernel) SetFlag(key, val string) error {
	if !flagSupported(key, kernelStringFlags) {
		return errors.New(fmt.Sprintf("kernel %s: flag not supported", k.Name()))
	}
	k.strVal[key] = val
	return nil
}

func New(flags int) *Kernel {

}

func NewKernel() *Kernel {
	ret := new(Kernel)

	ret.path = KERNEL_PATH

	ret.strVal = make(map[string]string)
	ret.bVal = make(map[string]bool)

	return ret
}

func (k *Kernel) Name() string {
	return k.strVal["name"]
}

func (k *Kernel) SetName(name string) {
	k.SetFlag("name", name)
}

func (k *Kernel) getFlags() []string {
	tmp := make([]string, len(k.strVal)+len(k.bVal))

	i := 0
	for key, val := range k.strVal {
		tmp[i] = key + "=" + val
		i++
	}
	for key := range k.bVal {
		tmp[i+len(k.strVal)] = key
		i++
	}
	return tmp
}

func (k *Kernel) Run() error {
	command := exec.Command(k.path, k.getFlags()...)
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}

func (k *Kernel) Stop() error {
	if !k.IsRunning() {
		return errors.New(NOT_RUNNING)
	}
	return nil
}

func (k *Kernel) IsRunning() bool {
	return k.running
}

func (k *Kernel) AddConnection(hub network.Hubber, eth uint8) error {
	return nil
}

func (k *Kernel) ShortInfo() string {
	return fmt.Sprintf("kernel: %s", k.Name())
}

func (k *Kernel) String() string {
	tmp := make([]string, len(kernelStringFlags)+len(kernelBoolFlags))
	for i, name := range kernelStringFlags {
		tmp[i] = name + ": " + k.strVal[name]
	}
	for i, name := range kernelBoolFlags {
		tmp[i+len(kernelStringFlags)] = fmt.Sprintf("%v: %v\n", name, k.bVal[name])
	}
	return strings.Join(tmp, "\n")
}

func (k *Kernel) addStartCommand(command string) {
	k.bashrc += command + "\n"
}

func routingCommand(ip string, eth int) {

}

func (k *Kernel) AddRouting(ip string, eth int) error {
	if k.IsRunning() {
		return errors.New(ALREADY_RUNNING)
	}
	return nil
}
