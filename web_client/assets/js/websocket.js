// 这里面定义了websocket检测心跳包

var url = "ws://localhost:8998/ws"

var lockReconnect = false; //避免重复连接
var websocket;// 心跳检测
var heartCheck = {
    timeout: 5000,
    timer: null,
    serverTimer: null,
    reset: function () {
        clearTimeout(this.timer);
        clearTimeout(this.serverTimer);
        this.start();
    },
    start: function () {
        let ts = this;
        this.timer = setTimeout(function () {
            console.log("send heartCheack")
            websocket.send(JSON.stringify({ "msg_id": "-1", "payload": "" }));
            ts.serverTimer = setTimeout(function () {
                websocket.onclose();
            }, ts.timeout)
        }, this.timeout)
    }
}

// 创建WebSocket链接
function createWebSocket() {
    if (window["WebSocket"]) {
        if (!url) return
        websocket = new WebSocket(url);  // WebSocket事件方法    
        wsHandle();
    } else {
        writeToScreen('浏览器不支持websocket')
    }
}

/* 初始化ws各个事件的方法 */
function wsHandle(url) {
    websocket.onopen = function (evt) {
        writeToScreen("Connection open");
        heartCheck.start();
    }
    websocket.onclose = function (evt) {
        console.log('websocket 断开: ' + evt.code + ' ' + evt.reason + ' ' + evt.wasClean);
    }
    websocket.onmessage = function (evt) {
        var data = JSON.parse(evt.data)
        if (data.msg_id == "-1") {
            console.log("received heartCheack")
        } else {
            console.log("received msg:" + evt.data);
            RegisterFuncMap.get(String(data.msg_id))(data.payload);
        }
        heartCheck.reset();
    }

    websocket.onerror = function (evt) {
        console.log('websocket 出错: ' + evt.data + '正在重连')
        reconnect();
    }
}

/* ws重新连接 */
function reconnect() {
    if (lockReconnect) return;
    lockReconnect = true;
    //没连接上会一直重连，设置延迟避免请求过多  
    setTimeout(function () {
        createWebSocket();
        lockReconnect = false;
    },
        2000);
}

/* 发送消息 */
function sendMessage(tag, payload) {
    var json = { "msg_id": tag, "payload": payload };
    var jsonStr = JSON.stringify(json);
    console.log("send msg:" + jsonStr)
    websocket.send(jsonStr);
    heartCheck.start()
}