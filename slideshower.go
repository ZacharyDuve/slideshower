package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ZacharyDuve/slideshower/imageloader"
	"github.com/gorilla/mux"
)

type nextResponse struct {
	ImageID imageloader.ImageFileID `json:"imageid"`
}

func main() {

	router := mux.NewRouter()

	var rootDirPath string
	flag.StringVar(&rootDirPath, "path", ".", "used for setting the path.")
	flag.Parse()

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

	router.HandleFunc("/photo/next", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("/photo/next was called")
		nextImageFileID := slideShow.Next()
		//defer nextImageFile.Close()
		//fmt.Println("Should have next photo")
		//photoData, err := ioutil.ReadAll(nextImageFile)

		fmt.Println("Photodata has been pulled", nextImageFileID, " ;")
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(err.Error()))
		} else {
			nextResponse := &nextResponse{ImageID: nextImageFileID}
			responseDataJson, _ := json.Marshal(nextResponse)
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(responseDataJson))
		}
	})

	router.HandleFunc("/photo/{id}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		imageFileID := imageloader.ImageFileID(vars["id"])
		imageFile := slideShow.GetImageByID(imageFileID)

		if imageFile == nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Unable to find image with id " + vars["id"]))
		} else {
			defer imageFile.Close()
			photoData, err := ioutil.ReadAll(imageFile)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				rw.Write([]byte(err.Error()))
			} else {
				rw.WriteHeader(http.StatusOK)
				rw.Write(photoData)
			}
		}
	})
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":8080", router)
}
