var worker;

function Register() {

    var userName = document.getElementById("register_lineedit_username")
    var password = document.getElementById("register_lineedit_password")
    var passwordAgain = document.getElementById("register_lineedit_passwordagain")
    var email = document.getElementById("register_lineedit_email")

    if (userName.value == "" || password.value == "") {
        alert("用户名或者密码为空")
        return false;
    }

    var payload = {
        UserName: userName.value,
        Password: password.value
    }
    console.log("注册")

    var json = { "msg_id": String(MESSAGE_ID.Register.req), "payload": JSON.stringify(payload) };

    worker.port.postMessage(json)
    return false
}

// 服务器回发的聊天消息
function RegisterBack(payload) {
    console.log("注册成功:" + payload.UserID)

    // 修改这个页面的父iframe, 登录成功之后进入主界面
    window.parent.frames[0].location = "login";
}

function create_register_worker() {
    if (!worker) {
        worker = new SharedWorker('assets/assets/js/websocket.js');
        worker.port.onmessage = function (e) {
            if (MESSAGE_ID.Login.res == e.data.msg_id) {
                RegisterBack(e.data.payload)
            }
        };
    }

    // document.getElementById('btn').addEventListener('click', function (e) {
    //     // 点击按钮时，通过port发送消息到shared worker
    //     worker.port.postMessage('nihao' + Math.random());
    // });
}