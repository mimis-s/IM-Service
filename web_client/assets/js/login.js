var worker;

// 跳转到注册界面
function TrunRegister() {
    window.parent.frames[0].location = "register";
}

function Login() {
    // console.log("1")

    // if (!websocket) {
    //     return false;
    // }
    // console.log("2")

    var userID = document.getElementById("input-username")
    var password = document.getElementById("input-password")
    if (userID.value == "") {
        return false;
    }
    var payload = {
        UserID: Number(userID.value),
        Password: password.value
    }

    var json = { "msg_id": String(MESSAGE_ID.Login.req), "payload": JSON.stringify(payload) };

    worker.port.postMessage(json)

    // sendMessage(String(MESSAGE_ID.Login.req), JSON.stringify(payload))

    return false
}

// 服务器回发的聊天消息
function LoginBack(payload) {
    console.log("登录成功:" + payload.UserID)

    // 修改这个页面的父iframe, 登录成功之后进入主界面
    window.parent.frames[0].location = "home";
}

function create_login_worker() {
    if (!worker) {
        worker = new SharedWorker('assets/assets/js/websocket.js');
        worker.port.onmessage = function (e) {
            if (MESSAGE_ID.Login.res == e.data.msg_id) {
                LoginBack(e.data.payload)
            }
        };
    }

    // document.getElementById('btn').addEventListener('click', function (e) {
    //     // 点击按钮时，通过port发送消息到shared worker
    //     worker.port.postMessage('nihao' + Math.random());
    // });
}