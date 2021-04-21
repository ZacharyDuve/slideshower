package imageloader

import (
	"fmt"
	"image"
	"log"

	//Needed to decode jpegs
	_ "image/jpeg"
	//Needed to decode png
	_ "image/png"
	"os"
)

type ImageFileID string

type ImageFile struct {
	filePath string
	file     *os.File
	id       ImageFileID
}

func NewImageFile(path string, id ImageFileID) *ImageFile {
	return &ImageFile{filePath: path, id: id}
}

func (this *ImageFile) GetImage() (img image.Image, err error) {
	f, err := os.Open(this.filePath)

	defer f.Close()
	if err == nil {
		img, _, err = image.Decode(f)
		if err != nil {
			fmt.Println("Error decoding image", this.filePath)
		}
	}
	return img, err
}

func (this *ImageFile) ID() ImageFileID {
	return this.id
}

func (this *ImageFile) Read(p []byte) (n int, err error) {
	if this.file == nil {
		this.file, err = os.Open(this.filePath)
		log.Println("Opened file:", this.filePath)
	}

	if err != nil {
		return 0, err
	}

	return this.file.Read(p)
}

func (this *ImageFile) Close() error {
	if this.file != nil {
		err := this.file.Close()
		this.file = nil
		return err
	}

	return nil
}
