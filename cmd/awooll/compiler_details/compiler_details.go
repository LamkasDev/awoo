package compiler_details

type CompileNodeValueDetails struct {
	Type     uint16
	Register uint8
	Address  CompileNodeValueDetailsAddress
}

type CompileNodeValueDetailsAddress struct {
	Register  uint8
	Immediate uint32
	Used      bool
}
