window.onload = function() {
    console.log("Window loaded!")

    function parseCmd() {
        // Get app info.
        var idx = parseInt(this.getAttribute("data-index"), 10)
        var type = parseInt(this.getAttribute("data-type"), 10)

        // Send POST request to back-end that submits command.
        fetch("/backend/submit", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ index: idx, type: type })
        })
        .then(data => {
            if (type == 0)
                console.log("Launched application successfully!")
            else
                console.log("Stopped application successfully!")
        })
        .catch((err) => {
            console.error(err)
        })
    }

    // Handle on clicks for app start buttons.
    var appStarts = document.getElementsByClassName("app-start")

    for (var i = 0; i < appStarts.length; i++) {
        appStarts[i].addEventListener("click", parseCmd)
    }

    // Handle on clicks for app stop buttons.
    var appStops = document.getElementsByClassName("app-stop")

    for (var i = 0; i < appStops.length; i++) {
        appStops[i].addEventListener("click", parseCmd)
    }
}