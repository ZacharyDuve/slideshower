package imageloader

import (
	"errors"
	"math/rand"
)

type ImageSlideShow struct {
	images            map[ImageFileID]*ImageFile
	slideShowOrder    []ImageFileID
	curSlideShowIndex int
}

func NewImageSlideShow(images []*ImageFile) *ImageSlideShow {
	iSS := &ImageSlideShow{}
	iSS.curSlideShowIndex = 0

	numImagesIn := len(images)
	iSS.images = make(map[ImageFileID]*ImageFile, numImagesIn)
	iSS.slideShowOrder = make([]ImageFileID, numImagesIn)
	for i, curImageIn := range images {
		if curImageIn != nil {
			curImageID := curImageIn.ID()
			iSS.images[curImageID] = curImageIn
			iSS.slideShowOrder[i] = curImageID
		}
	}

	iSS.Shuffle()

	return iSS
}

func (this *ImageSlideShow) Shuffle() {
	imagesLen := len(this.slideShowOrder)
	if imagesLen != 0 {
		//Do some shuffle magic
		for i := 0; i < imagesLen; i++ {
			i1 := rand.Intn(imagesLen)
			i2 := rand.Intn(imagesLen)

			tmpImg := this.slideShowOrder[i1]
			this.slideShowOrder[i1] = this.slideShowOrder[i2]
			this.slideShowOrder[i2] = tmpImg
		}
	}
}

func (this *ImageSlideShow) Next() ImageFileID {
	var retImageID ImageFileID

	if len(this.images) != 0 {
		if this.curSlideShowIndex == len(this.images) {
			this.Shuffle()
			this.curSlideShowIndex = 0
		}
		retImageID = this.slideShowOrder[this.curSlideShowIndex]
		this.curSlideShowIndex++
	} else {
		panic(errors.New("Cannot show when there are no images in slideshow"))
	}

	return retImageID
}

func (this *ImageSlideShow) GetImageByID(id ImageFileID) *ImageFile {
	return this.images[id]
}
