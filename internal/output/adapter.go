package output

import (
	"github.com/bad-noodles/makeup/internal/core/color"
)

type Adapter interface {
	Build(color.Palette)
}
