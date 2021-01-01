package imageloader

import (
	"errors"
	"math/rand"
)

type ImageSlideShow struct {
	images   []*ImageFile
	curIndex int
}

func NewImageSlideShow(imgs []*ImageFile) *ImageSlideShow {
	iSS := &ImageSlideShow{}
	iSS.curIndex = 0
	iSS.images = imgs

	if iSS.images == nil {
		iSS.images = make([]*ImageFile, 0)
	}

	iSS.Shuffle()

	return iSS
}

func (this *ImageSlideShow) Shuffle() {
	imagesLen := len(this.images)
	if imagesLen != 0 {
		//Do some shuffle magic
		for i := 0; i < imagesLen; i++ {
			i1 := rand.Intn(imagesLen)
			i2 := rand.Intn(imagesLen)

			tmpImg := this.images[i1]
			this.images[i1] = this.images[i2]
			this.images[i2] = tmpImg
		}
	}
}

func (this *ImageSlideShow) Next() *ImageFile {
	var retImage *ImageFile

	if len(this.images) != 0 {
		if this.curIndex == len(this.images) {
			this.Shuffle()
			this.curIndex = 0
		}
		retImage = this.images[this.curIndex]
		this.curIndex++
	} else {
		panic(errors.New("Cannot show when there are no images in slideshow"))
	}

	return retImage
}
