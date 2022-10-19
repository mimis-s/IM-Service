var worker;

window.addEventListener("load", init, false);

function init() {
    init_window()
}

function init_window() {
    output = document.getElementById("message_output");
}

function SendChat() {

    var msg = document.getElementById("input_window")
    if (msg.value == "") {
        return false;
    }
    var payload = {
        TestStr: msg.value
    }

    var payload = {
        TestStr: msg.value
    }

    var json = { "msg_id": String(MESSAGE_ID.Chat.req), "payload": JSON.stringify(payload) };

    worker.port.postMessage(json)

    return false
}

// 服务器回发的聊天消息
function SignleChatBack(payload) {
    writeToScreen("我:" + payload.TestStr)

    // 清空输入框
    var msg = document.getElementById("input_window")
    msg.value = ""
}


function create_chat_worker() {
    if (!worker) {
        worker = new SharedWorker('assets/assets/js/websocket.js');
        worker.port.onmessage = function (e) {
            if (MESSAGE_ID.Chat.res == e.data.msg_id) {
                SignleChatBack(e.data.payload)
            }
        };
    }
}