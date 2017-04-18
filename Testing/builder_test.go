package Testing

import (
	"testing"

	"../Builder"
)

func TestManufacturer(t *testing.T) {
	manufacturer := Builder.Manufacturer{}

	homePCBuilder := &Builder.HomeEditionPCBuilder{}
	manufacturer.SetBuilder(homePCBuilder)
	manufacturer.ConstructPC()
	homePC := (homePCBuilder).GetPC()

	if homePC.CPU() != "i3" {
		t.Errorf("Home edition PC cpu shoul be 'i3' but found %s", homePC.CPU())
	}

	if homePC.GPU() != "Intel Graphic Card" {
		t.Errorf("Home edition PC gpu shoul be 'Intel Graphic Card' but found %s", homePC.GPU())
	}

	gamePCBuilder := &Builder.GameEditionPCBuilder{}
	manufacturer.SetBuilder(gamePCBuilder)
	manufacturer.ConstructPC()
	gamePC := gamePCBuilder.GetPC()

	if gamePC.CPU() != "i7" {
		t.Errorf("Game edition PC cpu shoul be 'i7' but found %s", gamePC.CPU())
	}

	if gamePC.GPU() != "AMD Radeon X80" {
		t.Errorf("Game edition PC gpu shoul be 'Intel Graphic Card' but found %s", gamePC.GPU())
	}
}
