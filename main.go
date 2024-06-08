package main

import (
	"image"
	"os"
	"strings"

	"github.com/fatih/color"

	"image/gif"
	"image/jpeg"
	"image/png"
)

func fatalErr(err error) {
	if err != nil {
		color.Red("fatal error:")
		color.Red(err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 3 {
		color.Red("Too many or too less arguments.")
		color.Blue("Example correct command: imagiconv image.jpg png")
	}
	fileName := os.Args[1]
	format := os.Args[2]

	color.Blue("Converting image...")
	file, err := os.Open(fileName)
	defer file.Close()
	fatalErr(err)
	image, _, err := image.Decode(file)
	fatalErr(err)

	newFileName := strings.Split(fileName, ".")[0] + "." + format
	switch format {
	case "png":
		imageFile, err := os.Create(newFileName)
		fatalErr(err)
		err = png.Encode(imageFile, image)
	case "gif":
		imageFile, err := os.Create(newFileName)
		fatalErr(err)
		opt := gif.Options{
			NumColors: 256,
		}
		err = gif.Encode(imageFile, image, &opt)
	case "jpg", "jpeg":
		imageFile, err := os.Create(newFileName)
		fatalErr(err)
		opt := jpeg.Options{
			Quality: 100,
		}
		err = jpeg.Encode(imageFile, image, &opt)
	default:
		color.Red("unknown format...")
		os.Exit(1)
	}
	color.Green("succesfully converted file! :D")
}
