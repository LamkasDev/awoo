package driver

import (
	"github.com/LamkasDev/awoo-emu/cmd/awoomu/internal"
)

type AwooDriverAction func(internal *internal.AwooEmulatorInternal, driver *AwooDriver)
type AwooDriverActionError func(internal *internal.AwooEmulatorInternal, driver *AwooDriver) error

type AwooDriver struct {
	Id       uint16
	Name     string
	Tick     AwooDriverAction
	TickLong AwooDriverAction
	Clean    AwooDriverActionError
	Data     interface{}
}
