package flags

import (
	"os"
	"strconv"

	"github.com/jwalton/gchalk"
	"github.com/jwalton/go-supportscolor"
)

func ResolveColor() {
	color, ok := os.LookupEnv("FORCE_COLOR")
	if ok {
		colorLevel, _ := strconv.Atoi(color)
		gchalk.SetLevel(supportscolor.ColorLevel(colorLevel))
	}
}
