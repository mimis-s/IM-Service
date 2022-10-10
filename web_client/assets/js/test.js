
function testSend() {
    if (!websocket) {
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

    websocket.send(jsonStr);
    heartCheck.start()
    return false
}

function testOnConn() {
    createWebSocket()
}

function testUnConn() {
    writeToScreen("Connection Closed.")
    websocket.close()
}
