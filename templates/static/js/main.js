console.log("Start js app")

let element = document.querySelector("h2")
element.addEventListener("click", function (e) {
    this.style.color = "red"
})