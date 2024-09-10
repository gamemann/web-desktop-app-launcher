window.onload = function() {
    console.log("Window loaded!")

    function parseCmd() {
        // Get command.
        var cmd = this.getAttribute("data-command");

        // Send POST request to back-end that submits command.
        fetch("/backend/submit", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ command: cmd })
        })
        .then(data => {
            console.log("Command '" + cmd + "' executed successfully!")
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