package daisyui

import (
	"os"
	"text/template"

	"github.com/bad-noodles/makeup/internal/core/color"
)

type Daisy struct {
}

var configTemplate = `
module.exports = {
	daisyui: {
		themes: [
			{
				mytheme: {
					"primary": "{{.Primary}}",
					"secondary": {{.Secondary}}",
					"accent": "{{.Accent}}",
					"neutral": "{{.Background}}",
					"base-100": "{{.Base}}",
					"info": "{{.Info}}",
					"success": "{{.Success}}",
					"warning": "{{.Warning}}",
					"error": "{{.Error}}",
				},
			},
		],
	},
	plugins: [
		require('daisyui'),
	],
}
`

type DaisyPalette struct {
	Primary    color.Color
	Secondary  color.Color
	Accent     color.Color
	Background color.Color
	Base       color.Color
	Info       color.Color
	Success    color.Color
	Warning    color.Color
	Error      color.Color
}

func (d Daisy) Build(colors color.Palette) {
	tmpl, err := template.New("tailwind.config.js").Parse(configTemplate)

	if err != nil {
		panic(err)
	}

	tmpl.Execute(os.Stdout, DaisyPalette{
		Primary:    colors.Primary,
		Secondary:  colors.Secondary,
		Accent:     colors.Primary.Complementary(),
		Background: colors.Background,
		Base:       colors.Background.Darker(0.3),
		Info:       colors.Info,
		Success:    colors.Success,
		Warning:    colors.Warning,
		Error:      colors.Error,
	})
}
