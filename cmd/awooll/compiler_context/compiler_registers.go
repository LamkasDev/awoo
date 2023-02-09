package compiler_context

import (
	"fmt"

	"github.com/LamkasDev/awoo-emu/cmd/awooll/awerrors"
	"github.com/jwalton/gchalk"
)

type AwooCompilerRegisters struct {
	Entries  map[string]AwooCompilerContextRegisterEntry
	Position uint16
}

type AwooCompilerContextRegisterEntry struct {
	Name     string
	Register uint16
	Type     uint16
	Data     interface{}
}

func PushCompilerRegister(context *AwooCompilerContext, entry AwooCompilerContextRegisterEntry) (uint16, error) {
	entry.Register = context.Registers.Position
	context.Registers.Entries[entry.Name] = entry

	return entry.Register, nil
}

func PopCompilerRegister(context *AwooCompilerContext, name string) error {
	context.Registers.Position--
	delete(context.Registers.Entries, name)

	return nil
}

func GetCompilerRegister(context *AwooCompilerContext, name string) (AwooCompilerContextRegisterEntry, error) {
	entry, ok := context.Registers.Entries[name]
	if !ok {
		return AwooCompilerContextRegisterEntry{}, fmt.Errorf("%w: %s", awerrors.ErrorUnknownVariable, gchalk.Red(name))
	}

	return entry, nil
}
