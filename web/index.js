
function main() {
    console.log("Main is running")
    let imageEl = document.getElementById("mainImg");
    getImage(imageEl);
    window.setInterval(function(){getImage(imageEl);}, 5000);
}

//document.onload = main;

function getImage(imageEl) {
    fetch("/photo/next").then(response => response.json()).then(json => {
        console.log(json)
        let newSrc = "/photo/" + json.imageid;
        console.log("Updating photo to" + newSrc);
        imageEl.src = newSrc;
    })
}

main();
console.log("me")