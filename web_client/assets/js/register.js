// 这里面定义了所有消息的ID和对应调用的函数

const RegisterFuncMap = new Map();

window.addEventListener("load", init, false);

function init() {
    init_window()
}

function init_window() {
    output = document.getElementById("message_output");
    RegisterFuncMap.set("1", SignleChat);   // ID: 1 => 私聊函数
}