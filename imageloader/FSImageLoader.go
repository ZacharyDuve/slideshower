package imageloader

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type fsImageLoader struct {
	rootPath string
}

func NewFSImageLoader(rootPath string) (iL ImageLoader, err error) {
	if rootPath == "" {
		err = errors.New("Need to provide root path")
	} else {
		iL = &fsImageLoader{rootPath: rootPath}
	}

	return iL, err
}

func (this *fsImageLoader) LoadImages() (slideShow *ImageSlideShow, err error) {
	//fInfo, err := os.Stat(this.rootPath)
	imageFiles := make([]*ImageFile, 0)
	var curImageID int

	err = filepath.Walk(this.rootPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && isDisplayablePicture(path) {
			imageFiles = append(imageFiles, NewImageFile(path, ImageFileID(fmt.Sprint(curImageID))))
			curImageID++
		}
		return nil
	})

	for _, curImage := range imageFiles {
		fmt.Println(curImage)
	}
	fmt.Println(len(imageFiles), "Images loaded")
	return NewImageSlideShow(imageFiles), err
}

func isDisplayablePicture(path string) bool {
	lowerPath := strings.ToLower(path)
	isImageFileExtension := strings.HasSuffix(lowerPath, "jpeg") || strings.HasSuffix(lowerPath, "jpg") || strings.HasSuffix(lowerPath, "png")
	isNotHiddenFile := !strings.Contains(path, "/.")

	return isImageFileExtension && isNotHiddenFile
}
