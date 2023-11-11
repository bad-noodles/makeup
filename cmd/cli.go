package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bad-noodles/makeup/internal/output/daisyui"
	"github.com/bad-noodles/makeup/internal/parser"
)

func main() {
	var output daisyui.Daisy

	flag.Parse()
	themePath := flag.Arg(0)

	if themePath == "" {
		fmt.Println("Please provide a theme")
		flag.Usage()
		os.Exit(1)
	}

	themeData, err := os.ReadFile(themePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	theme, err := parser.ParseInput(themePath, themeData)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output.Build(theme)
}
