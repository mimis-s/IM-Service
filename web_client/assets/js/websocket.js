// 修改主要用于实现iframe+websocket多界面切换的想法

var MESSAGE_ID = {
    "HeartCheack": "-1",
    "Chat": {
        "req": CRC32("ChatSingleReq", 10),
        "res": CRC32("ChatSingleRes", 10),
    },
    "Login": {
        "req": CRC32("LoginReq", 10),
        "res": CRC32("LoginRes", 10),
    },
    "Register": {
        "req": CRC32("RegisterReq", 10),
        "res": CRC32("RegisterRes", 10),
    },
    "ErrCode": CRC32("CommonError", 10),
    "Test": "2",
}

// str是要变换的字符串,radix是生成编码的进制
function CRC32(str, radix = 10) {
    const Utf8Encode = function (string) {
        string = string.replace(/\r\n/g, "\n");
        let text = "";
        for (let n = 0; n < string.length; n++) {
            const c = string.charCodeAt(n);
            if (c < 128) {
                text += String.fromCharCode(c);
            } else if ((c > 127) && (c < 2048)) {
                text += String.fromCharCode((c >> 6) | 192);
                text += String.fromCharCode((c & 63) | 128);
            } else {
                text += String.fromCharCode((c >> 12) | 224);
                text += String.fromCharCode(((c >> 6) & 63) | 128);
                text += String.fromCharCode((c & 63) | 128);
            }
        }
        return text;
    }

    const makeCRCTable = function () {
        let c;
        const crcTable = [];
        for (let n = 0; n < 256; n++) {
            c = n;
            for (let k = 0; k < 8; k++) {
                c = ((c & 1) ? (0xEDB88320 ^ (c >>> 1)) : (c >>> 1));
            }
            crcTable[n] = c;
        }
        return crcTable;
    }

    const crcTable = makeCRCTable();
    const strUTF8 = Utf8Encode(str);
    let crc = 0 ^ (-1);
    for (let i = 0; i < strUTF8.length; i++) {
        crc = (crc >>> 8) ^ crcTable[(crc ^ strUTF8.charCodeAt(i)) & 0xFF];
    }
    crc = (crc ^ (-1)) >>> 0;
    return crc.toString(radix);
};


var portList = [];

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
            websocket.send(JSON.stringify({ "msg_id": String(MESSAGE_ID.HeartCheack), "payload": "" }));
            ts.serverTimer = setTimeout(function () {
                websocket.onclose();
            }, ts.timeout)
        }, this.timeout)
    }
}

// 创建WebSocket链接
function createWebSocket() {
    if (!url) return
    websocket = new WebSocket(url);  // WebSocket事件方法    
    return wsHandle();

}

/* 初始化ws各个事件的方法 */
function wsHandle(url) {
    websocket.onopen = function (evt) {
        console.log("Connection open");
        heartCheck.start();
    }
    websocket.onclose = function (evt) {
        console.log('websocket 断开: ' + evt.code + ' ' + evt.reason + ' ' + evt.wasClean);
        reconnect();
    }
    websocket.onmessage = function (evt) {
        var data = JSON.parse(evt.data)
        var msg_id = String(data.msg_id)
        if (msg_id == MESSAGE_ID.HeartCheack) {
            console.log("received heartCheack")
        } else if (msg_id == MESSAGE_ID.ErrCode) {
            // 出错了
            portList[0].postMessage(data);
        } else {
            console.log("received msg:" + evt.data);
            // RegisterFuncMap.get(msg_id)(data.payload);
            portList.forEach(port => {
                port.postMessage(data);
            });
        }
        heartCheck.reset();
    }

    websocket.onerror = function (evt) {
        console.log('websocket 出错: ' + evt.data + '正在重连')
        reconnect();
    }
    return websocket
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

// 连接服务器
function testOnConn() {
    console.log("加速连接服务器中......")
    // document.getElementById("connect_btn").disabled = "disabled"
    // document.getElementById("close_connect_btn").removeAttribute("disabled")

    createWebSocket()
}

// 断开服务器
function testUnConn() {
    console.log("服务器断开连接")
    // document.getElementById("close_connect_btn").disabled = "disabled"
    // document.getElementById("connect_btn").removeAttribute("disabled")

    websocket.close()
}

// sharedworker每一次创建都会触发onconnect
onconnect = function (e) {
    // 当触发了这个事件时，就建立对应的websocket连接
    if (!websocket) {
        console.log("重连服务器")
        websocket = createWebSocket()
    }

    var port = e.ports[0];
    // 然后把这个连接对应的port放在portList中
    portList.push(port);

    port.addEventListener('message', function (evt) {
        sendMessage(evt.data.msg_id, evt.data.payload)
    });

    port.start();
};