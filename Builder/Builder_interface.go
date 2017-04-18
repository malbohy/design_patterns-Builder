package Builder

// Each builder should implement this interface
type PCBuilder interface {
	SetRAM() PCBuilder
	SetHDD() PCBuilder
	SetCPU() PCBuilder
	SetOS() PCBuilder
	SetGPU() PCBuilder
	GetPC() PersonalComputer
}
