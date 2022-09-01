package img

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
)

func GenerateFavicon() (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)

	m := image.NewRGBA(image.Rect(0, 0, 150, 150))
	clr := color.RGBA{184, 97, 97, 1}

	draw.Draw(m, m.Bounds(), &image.Uniform{C: clr}, image.ZP, draw.Src)
	var img image.Image = m
	if err := jpeg.Encode(buffer, img, nil); err != nil {
		return nil, err
	}
	return buffer, nil
}
