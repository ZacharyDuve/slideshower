package imageloader

import (
	"image"
	//Needed to decode jpegs
	_ "image/jpeg"
	//Needed to decode png
	_ "image/png"
	"os"
)

type ImageFile struct {
	filePath string
}

func NewImageFile(path string) *ImageFile {
	return &ImageFile{filePath: path}
}

func (this *ImageFile) GetImage() (img image.Image, err error) {
	f, err := os.Open(this.filePath)

	defer f.Close()
	if err == nil {
		img, _, err = image.Decode(f)
	}
	return img, err
}
