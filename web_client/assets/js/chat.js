
function SendChat() {
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
    sendMessage(String(MESSAGE_ID.Chat), JSON.stringify(payload))
    msg.value = ""
    return false
}

// 服务器回发的聊天消息
function SignleChat(payload) {
    writeToScreen("我:" + payload.TestStr)
}