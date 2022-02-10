package tools

import (
	"fmt"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/webp"
)

func WEBPtoPNG(fileloc string) {
	file, err := os.Open(fileloc)

	filename := file.Name()

	filedir := strings.TrimSuffix(fileloc, filename)
	filename = strings.TrimSuffix(filename, ".webp")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	image, err := webp.Decode(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	pngfile, err := os.Create(filedir + filename + ".png")

	if err != nil {
		fmt.Println(err)
		return
	}

	err = png.Encode(pngfile, image)

	if err != nil {
		fmt.Println(err)
		return
	}
}
