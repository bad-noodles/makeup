package parser

import (
	"fmt"

	"github.com/bad-noodles/makeup/internal/core/color"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

type InputColor struct {
	Main string `hcl:"main"`
	Text string `hcl:"text,optional"`
}

type InputColors struct {
	Primary    InputColor `hcl:"primary,block"`
	Secondary  InputColor `hcl:"secondary,block"`
	Background InputColor `hcl:"background,block"`
	Error      InputColor `hcl:"error,block"`
	Warning    InputColor `hcl:"warning,block"`
	Info       InputColor `hcl:"info,block"`
	Success    InputColor `hcl:"success,block"`
}

type InputFormat struct {
	Colors InputColors `hcl:"colors,block"`
}

func parseColor(name string, colorInput InputColor) (color.Color, error) {
	output, err := color.New(colorInput.Main)

	if err != nil {
		return output, fmt.Errorf("could not parse color \"%s\"", name)
	}

	if colorInput.Text != "" {
		err = output.SetText(colorInput.Text)
	}

	if err != nil {
		return output, fmt.Errorf("could not parse color \"%s\"", name)
	}

	return output, nil
}

func ParseInput(filename string, rawInput []byte) (color.Palette, error) {
	var input InputFormat
	var palette color.Palette

	err := hclsimple.Decode(filename, rawInput, nil, &input)

	if err != nil {
		return palette, err
	}

	palette.Primary, err = parseColor("Primary", input.Colors.Primary)

	if err != nil {
		return palette, err
	}

	palette.Secondary, err = parseColor("Secondary", input.Colors.Secondary)

	if err != nil {
		return palette, err
	}

	palette.Background, err = parseColor("Background", input.Colors.Background)

	if err != nil {
		return palette, err
	}

	palette.Error, err = parseColor("Error", input.Colors.Error)

	if err != nil {
		return palette, err
	}

	palette.Warning, err = parseColor("Warning", input.Colors.Warning)

	if err != nil {
		return palette, err
	}

	palette.Info, err = parseColor("Info", input.Colors.Info)

	if err != nil {
		return palette, err
	}

	palette.Success, err = parseColor("Success", input.Colors.Success)

	return palette, err
}
