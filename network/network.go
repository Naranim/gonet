package network

import (
	"errors"
)

const (
	ALREADY_RUNNING = "network running"
	NOT_RUNNING     = "network not running"
)

const (
	VERBOSE = 1 << iota
	SATIC_ROUTING
	DYNAMIC_ROUTING
)

type Kerneler interface {
	New(flags int) error
	Run() error
	Stop() error
	IsRunning() bool
	AddConnection(hub Hubber, eth uint8) error
	AddRouting(ip string, eth int) error
	ShortInfo() string
}

type Hubber interface {
	Start() error
	Stop() error
}

type KernelProvider interface {
	NewKernel(name string, memSize int) *Kerneler
}

type HubProvider interface {
	NewHub() *Hubber
}

type Network struct {
	running        bool
	kernels        map[string]*Kerneler
	kernelProvider *KernelProvider
	hubProvider    *HubProvider
}

func New(kernelProvider *KernelProvider, hubProvider *HubProvider) *Network {
	ret := new(Network)
	ret.running = false
	ret.kernels = make(map[string]*Kerneler)
	ret.hubProvider = hubProvider
	ret.kernelProvider = kernelProvider
	return ret
}

func (n *Network) AddKernel(name string, k *Kerneler) error {
	if n.running {
		return errors.New(ALREADY_RUNNING)
	}
	n.kernels[name] = k
	return nil
}

func (n *Network) Run() error {
	if n.running {
		return errors.New(ALREADY_RUNNING)
	}

	for _, ker := range n.kernels {
		err := (*ker).Run()
		if err != nil {
			n.Stop()
			return err
		}
	}

	return nil
}

func (n *Network) Stop() error {
	if !n.running {
		return errors.New(NOT_RUNNING)
	}

	problems := ""

	for _, ker := range n.kernels {
		err := (*ker).Stop()
		if err != nil {
			if problems == "" {
				problems = err.Error()
			} else {
				problems += "; " + err.Error()
			}
		}
	}

	if problems != "" {
		return errors.New(problems)
	}

	return nil
}
