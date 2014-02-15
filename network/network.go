package network

import (
	"container/list"
	"errors"
)

const (
	ALREADY_RUNNING = "network running"
	NOT_RUNNING     = "network not running"
)

type Kerneler interface {
	Run() error
	Stop() error
	IsRunning() bool
	AddConnection(hub Hubber, eth uint8) error
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
	ret.kernels = list.New()
	ret.hubProvider = hubProvider
	ret.kernelProvider = kernelProvider
	return ret
}

func (n *Network) AddKernel(k *Kerneler) error {
	if n.running {
		return errors.New(ALREADY_RUNNING)
	}
	n.kernels.PushBack(k)
	return nil
}

func (n *Network) Run() error {
	if n.running {
		return errors.New(ALREADY_RUNNING)
	}

	for e := n.kernels.Front(); e != nil; e = e.Next() {
		var i interface{}
		i = e
		currKernel := i.(Kerneler)
		err := currKernel.Run()
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

	for e := n.kernels.Front(); e != nil; e = e.Next() {
		var i interface{}
		i = e
		currKernel := i.(Kerneler)
		err := currKernel.Stop()
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
