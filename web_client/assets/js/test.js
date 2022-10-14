
function testSend() {
    // 测试输出crc32加密
    writeToScreen(GetCrc32("LoginReq"))

    // if (!websocket) {
    //     return false;
    // }

    // var msg = document.getElementById("input_window")
    // if (msg.value == "") {
    //     return false;
    // }
    // var payload = {
    //     TestStr: msg.value
    // }
    // sendMessage(String(MESSAGE_ID.Test), JSON.stringify(payload))
    // msg.value = ""
    return false
}

function TestBack(payload) {
    writeToScreen("测试服务器回发" + payload.TestStr)
}