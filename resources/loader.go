package resources

import (
	"bufio"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"os"
)

func LoadImage(filepath string) (*ebiten.Image, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			panic(closeErr)
		}
	}()

	resource, _, err := image.Decode(bufio.NewReader(file))
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(resource), nil
}
