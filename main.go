package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"os"
)

var (
	version   = "unknown"
	gitCommit = "unknown"
)

func main() {
	srcName := flag.String("src", "sprite.png", "Path to the PNG image to convert")
	dstName := flag.String("dst", "sprite.inc", "Path to the output .inc file")
	flag.Parse()

	fmt.Printf("png2asm version: %s, git commit: %s\n", version, gitCommit)

	reader, err := os.Open(*srcName)
	if err != nil {
		exitWithError(fmt.Sprintf("Could not open image (%v)", err))
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		exitWithError(fmt.Sprintf("Could not decode PNG (%v)", err))
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	outFile, err := os.Create(*dstName)
	if err != nil {
		exitWithError(fmt.Sprintf("Could not create .inc file (%v)", err))
	}
	defer outFile.Close()

	fmt.Fprintln(outFile, "; Automatically generated sprite")
	fmt.Fprintln(outFile, "my_sprite LABEL BYTE")

	for y := range height {
		fmt.Fprintf(outFile, "    DB ")
		for x := range width {
			var index uint8
			if pimg, ok := img.(*image.Paletted); ok {
				index = pimg.ColorIndexAt(x, y)
			} else {
				exitWithError("Image is not in indexed mode. Please use a PNG image with a 256-color palette.")
			}

			if x == width-1 {
				fmt.Fprintf(outFile, "%d", index)
			} else {
				fmt.Fprintf(outFile, "%d,", index)
			}
		}
		fmt.Fprintln(outFile)
	}

	fmt.Printf("Conversion complete! File generated: sprite.inc (%dx%d)\n", width, height)
}

func exitWithError(msg string) {
	fmt.Printf("Error: %s\n", msg)
	os.Exit(1)
}
