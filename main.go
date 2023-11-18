package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fogleman/gg"
	"github.com/gllmjln/arch-gen/parser"
)

func main() {

	var inputFile string
	var outputFile string
	var outputImgHeight int
	var outputImgWidth int

	flag.StringVar(&inputFile, "i", "arch-gen.yml", "the yaml file containing the diagram definition")
	flag.StringVar(&outputFile, "o", "arch-gen.png", "the output png file")
	flag.IntVar(&outputImgHeight, "height", 1024, "the output png height")
	flag.IntVar(&outputImgWidth, "width", 1024, "the output png width")

	flag.Parse()

	p := parser.Parser{}
	err := p.ParseInput(inputFile)

	if err != nil {
		fmt.Printf("Could not parse file input: %v\n", err)
		os.Exit(1)
	}

	ctx := gg.NewContext(outputImgWidth, outputImgHeight)
	ctx.SetRGB(1, 1, 1)
	ctx.SetLineWidth(4)

	err = p.Root.Draw(ctx, 20, 20)

	if err != nil {
		fmt.Printf("Could not generate diagram %v\n", err)
		os.Exit(1)
	}

	ctx.SavePNG(outputFile)

	fmt.Println("Image generated")
	fmt.Println("Have a nice day!✨")
}
