package Builder

type HomeEditionPCBuilder struct {
	pc PersonalComputer
}

func (b *HomeEditionPCBuilder) SetRAM() PCBuilder {
	b.pc.ramCap = 4
	return b
}

func (b *HomeEditionPCBuilder) SetHDD() PCBuilder {
	b.pc.hddCap = 500
	return b
}

func (b *HomeEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.cpu = "i3"
	return b
}

func (b *HomeEditionPCBuilder) SetOS() PCBuilder {
	b.pc.os = "Windows 7 Home Edition"
	return b
}

func (b *HomeEditionPCBuilder) SetGPU() PCBuilder {
	b.pc.gpu = "Intel Graphic Card"
	return b
}

func (b *HomeEditionPCBuilder) GetPC() PersonalComputer {
	return b.pc
}
