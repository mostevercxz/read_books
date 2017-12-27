package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputFile := flag.String("i", "", "The input image file")
	outputFormat := flag.String("fmt", "jpg", "The output image format")
	if len(os.Args) <= 3 {
		flag.Usage()
		os.Exit(0)
	} else {
		flag.Parse()
	}

	// Check input file valid
	f, err := os.Open(*inputFile)
	defer f.Close()
	if err != nil {
		fmt.Println(os.Stderr, "Failed to open input file", *inputFile)
		os.Exit(1)
	}

	// Check if extension is the same
	fileExtension := filepath.Ext(*inputFile)
	*outputFormat = "." + *outputFormat
	if fileExtension == *outputFormat {
		fmt.Println("Nothing to do, convert complete")
		os.Exit(0)
	}

	// Open output file
	filePrefix := strings.TrimSuffix(*inputFile, fileExtension)
	outputFileName := filePrefix + *outputFormat
	outf, err := os.Create(outputFileName)
	defer outf.Close()

	// Start encode and decode
	img, kind, err := image.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "image decode error:%s,kind:%s", err, kind)
		os.Exit(1)
	}
	// Check output format valid
	switch *outputFormat {
	case "jpg":
		jpeg.Encode(outf, img, &jpeg.Options{Quality: 100})
	case "png":
		png.Encode(outf, img)
	case "gif":
		gif.Encode(outf, img, &gif.Options{NumColors: 256})
	default:
		fmt.Println("Invalid output format")
		os.Exit(1)
	}

	fmt.Printf("Successfully convert from %s to %s\n", *inputFile, outputFileName)
	os.Exit(0)
}
