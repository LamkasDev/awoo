package driver

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
)

type AwooDriverTick func(internal *internal.AwooEmulatorInternal, driver *AwooDriver)

type AwooDriver struct {
	Id   uint16
	Name string
	Tick AwooDriverTick
	Data interface{}
}
