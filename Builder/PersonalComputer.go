package Builder

// Product to build
type PersonalComputer struct {
	ramCap int
	hddCap int
	cpu    string
	os     string
	gpu    string
}

func (m *PersonalComputer) RamCap() int {
	return m.ramCap
}

func (m *PersonalComputer) HDDCap() int {
	return m.hddCap
}

func (m *PersonalComputer) CPU() string {
	return m.cpu
}

func (m *PersonalComputer) OS() string {
	return m.os
}

func (m *PersonalComputer) GPU() string {
	return m.gpu
}
