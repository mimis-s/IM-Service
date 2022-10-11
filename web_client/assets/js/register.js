// 这里面定义了所有消息的ID和对应调用的函数

var MESSAGE_ID = {
    "HeartCheack": "-1",
    "Chat": "1",
    "Test": "2",
}

const RegisterFuncMap = new Map();

window.addEventListener("load", init, false);

function init() {
    init_window()
}

function init_window() {
    output = document.getElementById("message_output");
    RegisterFuncMap.set(MESSAGE_ID.Test, TestBack);   // ID: 2 => 测试用
    RegisterFuncMap.set(MESSAGE_ID.Chat, SignleChat);   // ID: 1 => 私聊函数
}