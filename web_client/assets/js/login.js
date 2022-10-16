
function Login() {
    if (!websocket) {
        return false;
    }

    var userID = document.getElementById("input-username")
    var password = document.getElementById("input-password")
    if (userID.value == "") {
        return false;
    }
    var payload = {
        UserID: Number(userID.value),
        Password: password.value
    }
    sendMessage(String(MESSAGE_ID.Login.req), JSON.stringify(payload))
    return false
}

// 服务器回发的聊天消息
function LoginBack(payload) {
    writeToScreen("登录:" + payload.UserID)
}