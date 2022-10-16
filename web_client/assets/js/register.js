// 这里面定义了所有消息的ID和对应调用的函数

var MESSAGE_ID

var RegisterFuncMap;

window.addEventListener("load", init, false);

function init() {
    init_args()
    init_window()
}

function init_args() {
    MESSAGE_ID = {
        "HeartCheack": "-1",
        "Chat": {
            "req": CRC32("SignleChatReq", 10),
            "res": CRC32("SignleChatRes", 10),
        },
        "Login": {
            "req": CRC32("LoginReq", 10),
            "res": CRC32("LoginRes", 10),
        },
        "Test": "2",
    }
    RegisterFuncMap = new Map()
}

function init_window() {
    output = document.getElementById("message_output");
    RegisterFuncMap.set(MESSAGE_ID.Test, TestBack);             // 测试用
    RegisterFuncMap.set(MESSAGE_ID.Chat.res, SignleChatBack);   // 私聊
    RegisterFuncMap.set(MESSAGE_ID.Login.res, LoginBack);       // 登录
}