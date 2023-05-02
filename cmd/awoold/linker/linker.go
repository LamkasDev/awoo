package linker

import "github.com/LamkasDev/awoo-emu/cmd/common/elf"

type AwooLinker struct {
	Contents []elf.AwooElf
	Settings AwooLinkerSettings
}

type AwooLinkerSettings struct {
	Path     string
	Mappings AwooLinkerMappings
}

func SetupLinker(settings AwooLinkerSettings) AwooLinker {
	return AwooLinker{
		Settings: settings,
	}
}

func LoadLinker(linker *AwooLinker, contents []elf.AwooElf) {
	linker.Contents = contents
}
