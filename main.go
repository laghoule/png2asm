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

	fmt.Printf("Version: %s, Git Commit: %s\n", version, gitCommit)

	reader, err := os.Open(*srcName)
	if err != nil {
		fmt.Printf("Error: Could not open image (%v)\n", err)
		return
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: Could not decode PNG (%v)\n", err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	outFile, err := os.Create(*dstName)
	if err != nil {
		fmt.Printf("Error: Could not create .inc file (%v)\n", err)
		return
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
				panic("Image is not in indexed mode. Please use a PNG image with a 256-color palette.")
			}

			if x == width-1 {
				fmt.Fprintf(outFile, "%d", index)
			} else {
				fmt.Fprintf(outFile, "%d,", index)
			}
		}
		fmt.Fprintln(outFile)
	}

	fmt.Printf("Conversion complete! File generated: sprite.inc (%dx%d)\n", height, height)
}
