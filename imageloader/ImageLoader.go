package imageloader

type ImageLoader interface {
	LoadImages() (slideShow *ImageSlideShow, err error)
}
