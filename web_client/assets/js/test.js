
function testSend() {
    if (!websocket) {
        return false;
    }

    var msg = document.getElementById("input_window")
    if (msg.value == "") {
        return false;
    }
    var payload = {
        TestStr: msg.value
    }
    sendMessage("1", JSON.stringify(payload))
    msg.value = ""
    return false
}

function testOnConn() {
    writeToScreen("加速连接服务器中......")
    document.getElementById("connect_btn").disabled = "disabled"
    document.getElementById("close_connect_btn").removeAttribute("disabled")

    createWebSocket()
}

function testUnConn() {
    writeToScreen("服务器断开连接")
    document.getElementById("close_connect_btn").disabled = "disabled"
    document.getElementById("connect_btn").removeAttribute("disabled")

    websocket.close()
}

// 服务器回发的聊天消息
function SignleChat(payload) {
    writeToScreen("我:" + payload.TestStr)
}