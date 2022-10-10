
var conn;

function testSend() {
    if (!conn) {
        return false;
    }

    var msg = document.getElementById("input_window")
    if (msg.value == "") {
        return false;
    }
    var json = { "msg_id": "1", "payload": msg.value }; //创建对象；
    var jsonStr = JSON.stringify(json);       //转为JSON字符串
    writeToScreen(jsonStr)
    msg.value = "";

    conn.send(jsonStr);
    return false
}

function testOnConn() {
    if (window["WebSocket"]) {
        conn = new WebSocket("ws://localhost:8998/ws");
        conn.onopen = function (evt) {
            writeToScreen("Connection open")
        }
        conn.onclose = function (evt) {
            writeToScreen('websocket 断开: ' + evt.code + ' ' + evt.reason + ' ' + evt.wasClean)
        }
        conn.onmessage = function (evt) {
            writeToScreen(evt.data)
        }

        conn.onerror = function (evt) {
            writeToScreen('websocket 出错: ' + evt.data)
        }
    } else {
        writeToScreen("浏览器不支持websocket")
    }

}

function testUnConn() {
    writeToScreen("Connection Closed.")
    conn.close()
}

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
