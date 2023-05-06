package cpu

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

const AwooCrsMstatus = arch.AwooRegister(0x400) // TODO: this should be 0x300
const AwooCrsMstatusHalt = 1
