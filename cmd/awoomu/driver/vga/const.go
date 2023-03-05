package vga

import "github.com/LamkasDev/awoo-emu/cmd/common/arch"

const AwooDriverIdVga = 0x000
const AwooDriverVgaVector = arch.AwooRegister(0xb8000)
const AwooDriverVgaSize = arch.AwooRegister(32768)
const AwooDriverVgaCharacterSize = 2

const AwooDriverVgaFps = 90
const AwooDriverVgaFrameWidth = 80
const AwooDriverVgaFrameHeight = 25
const AwooDriverVgaFrameSize = AwooDriverVgaFrameWidth * AwooDriverVgaFrameHeight

var AwooDriverVgaFrame [AwooDriverVgaFrameSize]uint32
