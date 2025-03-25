const button = document.querySelector(".increment-button")

button.addEventListener("click", () => {
    let counterElement = document.getElementById("countings")
    let counter = parseInt(counterElement.textContent)
    counter += 1 
    counterElement.textContent = counter
})
