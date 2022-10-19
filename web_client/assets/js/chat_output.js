var output;

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
