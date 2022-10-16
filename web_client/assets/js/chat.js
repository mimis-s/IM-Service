
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
    sendMessage(String(MESSAGE_ID.Chat.req), JSON.stringify(payload))
    msg.value = ""
    return false
}

// 服务器回发的聊天消息
function SignleChatBack(payload) {
    writeToScreen("我:" + payload.TestStr)
}