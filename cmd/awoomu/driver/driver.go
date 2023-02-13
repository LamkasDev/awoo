package driver

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
)

type AwooDriverAction func(internal *internal.AwooEmulatorInternal, driver *AwooDriver)

type AwooDriver struct {
	Id       uint16
	Name     string
	Tick     AwooDriverAction
	TickLong AwooDriverAction
	Clean    AwooDriverAction
	Data     interface{}
}
