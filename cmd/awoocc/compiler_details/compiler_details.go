package compiler_details

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

type CompileNodeValueDetails struct {
	Type     uint16
	Register arch.AwooRegisterIndex
	Address  CompileNodeValueDetailsAddress
}

type CompileNodeValueDetailsAddress struct {
	Register  arch.AwooRegisterIndex
	Immediate uint32
	Used      bool
}
