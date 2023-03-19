package compiler_details

import (
	"github.com/LamkasDev/awoo-emu/cmd/common/cpu"
	"github.com/LamkasDev/awoo-emu/cmd/common/types"
)

type CompileNodeValueDetails struct {
	Type     types.AwooTypeId
	Register cpu.AwooRegisterId
	Address  CompileNodeValueDetailsAddress
}

type CompileNodeValueDetailsAddress struct {
	Register  cpu.AwooRegisterId
	Immediate uint32
	Used      bool
}
