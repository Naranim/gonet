package kernel

type KernelProvider struct{}

func NewKernelProvider() *KernelProvider {
	return new(KernelProvider)
}

func NewKernel(name string, memSize int) *Kerneler {
	return kernel.NewKernel()
}
