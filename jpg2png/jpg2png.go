package jpg2png

import (
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

func jpg2png(reader io.Reader, writer io.Writer) error {
	image, error := jpeg.Decode(reader)
	if error != nil {
		return error
	}

	return png.Encode(writer, image)
}

func Convert(arg string) error {
	reader, error := os.Open(arg)
	if error != nil {
		return error
	}
	defer reader.Close()

	writer, error := os.Create(strings.Replace(arg, ".jpg", ".png", 1))
	if error != nil {
		return error
	}
	defer writer.Close()

	return jpg2png(reader, writer)
}
