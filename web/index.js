
const delaySeconds = 5;

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}


function main() {
    //let windowHeight = window.innerHeight;
    //let windowWidth = window.innerWidth;
    console.log("Main is running")
    //Get the first image ready
    let curImage = document.getElementById("img0");
    //image0Elm.hidden = true;
    preloadNextImage(curImage);
    //Load in the second image
    let nextImage = document.getElementById("img1");
    nextImage.hidden = true;
    //But don't show it yet
    preloadNextImage(nextImage);

    runLoop(curImage, nextImage);
}

async function runLoop(curImage, nextImage) {
    while(true) {
        console.log("Swapping images");
        //Show current image for this time
        await sleep(delaySeconds * 1000)
        //Swap the two images
        curImage.hidden = true;
        let tempImage = curImage;
        nextImage.hidden = false;
        curImage = nextImage;
        nextImage = tempImage;
        //Start loading in the next image
        preloadNextImage(nextImage);
    }
}

function preloadNextImage(imageEl) {
    fetch("/photo/next").then(response => response.json()).then(json => {
        console.log(json)
        let newSrc = "/photo/" + json.imageid;
        console.log("Updating photo to" + newSrc);
        imageEl.src = newSrc;
    })
}

main();