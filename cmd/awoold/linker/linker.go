package linker

import "github.com/LamkasDev/awoo-emu/cmd/common/elf"

type AwooLinker struct {
	Contents map[string]elf.AwooElf
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

func LoadLinker(linker *AwooLinker, contents map[string]elf.AwooElf) {
	linker.Contents = contents
}
