package emu

import (
	"os/user"
	"path"
	"testing"
)

func TestEmu(t *testing.T) {
	u, _ := user.Current()
	input := path.Join(u.HomeDir, "Documents", "awoo", "data", "obj", "input.awoobj")
	Load(input)
}
