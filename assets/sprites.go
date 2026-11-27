package assets

import (
	"embed"
	"image"
	"image/png"
	"strings"
)

//go:embed sprites/*.png
var files embed.FS

func LoadSprites() (map[string]image.Image, error) {
	path := "sprites"
	entries, err := files.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var m = make(map[string]image.Image)
	for _, e := range entries {
		filename := e.Name()
		img, err := loadImage(path + "/" + filename)
		if err != nil {
			return nil, err
		}
		spriteName := strings.TrimSuffix(filename, ".png")
		m[spriteName] = img
	}
	return m, nil
}

func loadImage(path string) (image.Image, error) {
	reader, err := files.Open(path)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	img, err := png.Decode(reader)
	return img, err
}
