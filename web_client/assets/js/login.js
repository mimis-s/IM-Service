
function Login() {
    console.log("1")

    if (!websocket) {
        return false;
    }
    console.log("2")

    var userID = document.getElementById("input-username")
    var password = document.getElementById("input-password")
    if (userID.value == "") {
        return false;
    }
    var payload = {
        UserID: Number(userID.value),
        Password: password.value
    }
    console.log("3")

    sendMessage(String(MESSAGE_ID.Login.req), JSON.stringify(payload))
    console.log("发送登录成功")
    var iframe_temp = window.frames["iframeeducation"]
    iframe_temp.location = "chat"
    return false
}

// 服务器回发的聊天消息
function LoginBack(payload) {
    writeToScreen("登录:" + payload.UserID)
}