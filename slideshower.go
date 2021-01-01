package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"github.com/ZacharyDuve/slideshower/imageloader"
)

func main() {
	var rootDirPath string
	flag.StringVar(&rootDirPath, "path", ".", "used for setting the path.")
	flag.Parse()
	app := app.New()
	app.Settings().SetTheme(theme.DarkTheme())
	w := app.NewWindow("SlideShower")
	w.Resize(fyne.NewSize(200, 200))
	iL, err := imageloader.NewFSImageLoader(rootDirPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	slideShow, err := iL.LoadImages()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	img := canvas.NewImageFromImage(image.NewRGBA(image.Rect(0, 0, 100, 100)))
	img.FillMode = canvas.ImageFillContain
	//container := fyne.NewContainerWithLayout(layout.NewCenterLayout(), img)
	//w.SetContent(container)
	//theme.Back
	w.SetContent(img)

	//w.SetFullScreen(true)

	go displayImage(img, slideShow)

	//image := canvas.NewImageFromFile("IMG_3331.jpg")

	w.ShowAndRun()
}

func displayImage(img *canvas.Image, slideShow *imageloader.ImageSlideShow) {
	for {
		curImageFile := slideShow.Next()
		curImage, err := curImageFile.GetImage()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// image := canvas.NewImageFromImage(curImage)
		img.Image = curImage
		img.Refresh()
		//image.FillMode = canvas.ImageFillContain
		//fyne.NewContainerWithLayout(layout.NewCenterLayout(), image)
		//w.SetContent(image)

		time.Sleep(time.Second * 5)
	}
}
