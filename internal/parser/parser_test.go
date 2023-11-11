package parser

import "testing"

func TestParseInputOnlyRequiredColors(t *testing.T) {
	input := `
colors {
	primary {
		main = "#FFFFFF"
	}
	secondary {
		main = "#000"
	}
	background {
		main = "#ccc"
	}
	text {
		main = "#3f3FcC"
	}
	error {
		main = "rgb(0,0,0)"
	}
	warning {
		main = "rgba(0,0,0,1)"
	}
	info {
		main = "rgba(0,0,0,1)"
	}
	success {
		main = "rgba(0,0,0,1)"
	}
}
`

	output, err := ParseInput("test.hcl", []byte(input))

	if err != nil {
		t.Fatal(err)
	}

	if output.Primary.Main == nil {
		t.Fatal("Primary color is not set")
	}

	if output.Primary.Dark == nil {
		t.Fatal("Primary dark color is not set")
	}

	if output.Primary.Light == nil {
		t.Fatal("Primary light color is not set")
	}

	if output.Primary.Text == nil {
		t.Fatal("Primary text color is not set")
	}

	if output.Secondary.Main == nil {
		t.Fatal("Secondary color is not set")
	}

	if output.Secondary.Dark == nil {
		t.Fatal("Secondary dark color is not set")
	}

	if output.Secondary.Light == nil {
		t.Fatal("Secondary light color is not set")
	}

	if output.Secondary.Text == nil {
		t.Fatal("Secondary text color is not set")
	}

	if output.Background.Main == nil {
		t.Fatal("Background color is not set")
	}

	if output.Background.Dark == nil {
		t.Fatal("Background dark color is not set")
	}

	if output.Background.Light == nil {
		t.Fatal("Background light color is not set")
	}

	if output.Background.Text == nil {
		t.Fatal("Background text color is not set")
	}

	if output.Error.Main == nil {
		t.Fatal("Error color is not set")
	}

	if output.Error.Dark == nil {
		t.Fatal("Error dark color is not set")
	}

	if output.Error.Light == nil {
		t.Fatal("Error light color is not set")
	}

	if output.Error.Text == nil {
		t.Fatal("Error text color is not set")
	}

	if output.Warning.Main == nil {
		t.Fatal("Warning color is not set")
	}

	if output.Warning.Dark == nil {
		t.Fatal("Warning dark color is not set")
	}

	if output.Warning.Light == nil {
		t.Fatal("Warning light color is not set")
	}

	if output.Warning.Text == nil {
		t.Fatal("Warning text color is not set")
	}

	if output.Info.Main == nil {
		t.Fatal("Info color is not set")
	}

	if output.Info.Dark == nil {
		t.Fatal("Info dark color is not set")
	}

	if output.Info.Light == nil {
		t.Fatal("Info light color is not set")
	}

	if output.Info.Text == nil {
		t.Fatal("Info text color is not set")
	}

	if output.Success.Main == nil {
		t.Fatal("Success color is not set")
	}

	if output.Success.Dark == nil {
		t.Fatal("Success dark color is not set")
	}

	if output.Success.Light == nil {
		t.Fatal("Success light color is not set")
	}

	if output.Success.Text == nil {
		t.Fatal("Success text color is not set")
	}
}

func TestParseInputAllColors(t *testing.T) {
	input := `
colors {
	primary {
		main = "#FFFFFF"
	}
	secondary {
		main = "#000"
	}
	background {
		main = "#ccc"
	}
	text {
		main = "#3f3FcC"
	}
	error {
		main = "rgb(0,0,0)"
	}
	warning {
		main = "rgba(0,0,0,1)"
	}
	info {
		main = "rgba(0,0,0,1)"
	}
	success {
		main = "rgba(0,0,0,1)"
	}
}
`

	output, err := ParseInput("test.hcl", []byte(input))

	if err != nil {
		t.Fatal(err)
	}

	if output.Primary.Main == nil {
		t.Fatal("Primary color is not set")
	}

	if output.Primary.Dark == nil {
		t.Fatal("Primary dark color is not set")
	}

	if output.Primary.Light == nil {
		t.Fatal("Primary light color is not set")
	}

	if output.Primary.Text == nil {
		t.Fatal("Primary text color is not set")
	}

	if output.Secondary.Main == nil {
		t.Fatal("Secondary color is not set")
	}

	if output.Secondary.Dark == nil {
		t.Fatal("Secondary dark color is not set")
	}

	if output.Secondary.Light == nil {
		t.Fatal("Secondary light color is not set")
	}

	if output.Secondary.Text == nil {
		t.Fatal("Secondary text color is not set")
	}

	if output.Background.Main == nil {
		t.Fatal("Background color is not set")
	}

	if output.Background.Dark == nil {
		t.Fatal("Background dark color is not set")
	}

	if output.Background.Light == nil {
		t.Fatal("Background light color is not set")
	}

	if output.Background.Text == nil {
		t.Fatal("Background text color is not set")
	}

	if output.Error.Main == nil {
		t.Fatal("Error color is not set")
	}

	if output.Error.Dark == nil {
		t.Fatal("Error dark color is not set")
	}

	if output.Error.Light == nil {
		t.Fatal("Error light color is not set")
	}

	if output.Error.Text == nil {
		t.Fatal("Error text color is not set")
	}

	if output.Warning.Main == nil {
		t.Fatal("Warning color is not set")
	}

	if output.Warning.Dark == nil {
		t.Fatal("Warning dark color is not set")
	}

	if output.Warning.Light == nil {
		t.Fatal("Warning light color is not set")
	}

	if output.Warning.Text == nil {
		t.Fatal("Warning text color is not set")
	}

	if output.Info.Main == nil {
		t.Fatal("Info color is not set")
	}

	if output.Info.Dark == nil {
		t.Fatal("Info dark color is not set")
	}

	if output.Info.Light == nil {
		t.Fatal("Info light color is not set")
	}

	if output.Info.Text == nil {
		t.Fatal("Info text color is not set")
	}

	if output.Success.Main == nil {
		t.Fatal("Success color is not set")
	}

	if output.Success.Dark == nil {
		t.Fatal("Success dark color is not set")
	}

	if output.Success.Light == nil {
		t.Fatal("Success light color is not set")
	}

	if output.Success.Text == nil {
		t.Fatal("Success text color is not set")
	}
}
