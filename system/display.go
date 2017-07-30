package system

type Display interface{
	Initialize(machine *VirtualMachine)
	Render(machine *VirtualMachine)
	Close()
}
