package Builder

type GameEditionPCBuilder struct {
	pc PersonalComputer
}

func (b *GameEditionPCBuilder) SetRAM() PCBuilder {
	b.pc.ramCap = 16
	return b
}

func (b *GameEditionPCBuilder) SetHDD() PCBuilder {
	b.pc.hddCap = 500
	return b
}

func (b *GameEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.cpu = "i7"
	return b
}

func (b *GameEditionPCBuilder) SetOS() PCBuilder {
	b.pc.os = "Windows 7 Ultimate"
	return b
}

func (b *GameEditionPCBuilder) SetGPU() PCBuilder {
	b.pc.gpu = "AMD Radeon X80"
	return b
}

func (b *GameEditionPCBuilder) GetPC() PersonalComputer {
	return b.pc
}
