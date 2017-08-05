package system

type Display interface {
	Initialize()
	Render(machine *VirtualMachine)
	UpdateKeys(machine *VirtualMachine)
	Close()
}
