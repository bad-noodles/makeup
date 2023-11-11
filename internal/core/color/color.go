package color

import (
	"fmt"
	"image/color"

	"github.com/go-playground/colors"
	"github.com/muesli/gamut"
)

type Color struct {
	color color.Color
	text  *Color
}

func New(input string) (Color, error) {
	color, err := colors.Parse(input)

	return Color{color: color}, err
}

func new(input color.Color) Color {
	return Color{color: input}
}

func (c *Color) SetText(input string) error {
	text, err := New(input)

	if err != nil {
		return err
	}

	c.text = &text

	return nil
}

func (c Color) Hex() string {
	return gamut.ToHex(c.color)
}

func to255(v uint32) uint32 {
	return (v * 255) / 0xffff
}

func (c Color) String() string {
	r, g, b, a := c.color.RGBA()

	if a == 0xffff {
		return c.Hex()
	}

	return fmt.Sprintf(
		"rgba(%d,%d,%d,%.1f)",
		to255(r),
		to255(g),
		to255(b),
		float64(a)/float64(0xffff))
}

func (c Color) Darker(percent float64) Color {
	return new(gamut.Darker(c.color, percent))
}

func (c Color) Lighter(percent float64) Color {
	return new(gamut.Lighter(c.color, percent))
}

func (c Color) Contrast() Color {
	return new(gamut.Contrast(c.color))
}

func (c Color) Text() Color {
	if c.text != nil {
		return *c.text
	}

	return c.Contrast()
}

func (c Color) Complementary() Color {
	return new(gamut.Complementary(c.color))
}

func (c Color) HueOffset(amount int) Color {
	return new(gamut.HueOffset(c.color, amount))
}
