package Builder

//Manufacturer object which aware of build process for builder type
type Manufacturer struct {
	b PCBuilder
}

func (m *Manufacturer) SetBuilder(builder PCBuilder) {
	m.b = builder
}

func (m *Manufacturer) ConstructPC() {
	m.b.SetCPU().SetHDD().SetRAM().SetGPU().SetOS()
}
