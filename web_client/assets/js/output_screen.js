var output;

window.addEventListener("load", init, false);

function init() {
    init_window()
}

function init_window() {
    output = document.getElementById("message_output");
}


function writeToScreen(message) {
    var pre = document.createElement("p");
    pre.style.wordWrap = "break-word";
    timeNow = new Date()
    year = timeNow.getFullYear()
    month = timeNow.getMonth() + 1
    day = timeNow.getDate()
    // day2=timeNow.getDay()
    hour = timeNow.getHours()
    min = timeNow.getMinutes()
    sec = timeNow.getSeconds()
    msec = timeNow.getMilliseconds()
    pre.innerHTML = "[" + year + "/" + month + "/" + day + " " + hour + ":" + min + ":" + sec + "." + msec + "]<br>"
    pre.innerHTML += message;
    output.appendChild(pre);
    output.scrollTop = output.scrollHeight
}
