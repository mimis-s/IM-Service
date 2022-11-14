var worker;

$(function () {

    // 提交好友申请
    $(document).on('submit', 'addFriends', function (e) {
        e.preventDefault();

        var applyFriendID = $(this).find('input[id=applyFriendID]').val();
        applyFriendID = $.trim(applyFriendID);

        if (applyFriendID) {

            // 读取输入, 发送给服务器
            var payload = {
                ApplyFriendsID: applyFriendID,
            }

            var json = { "msg_id": String(MESSAGE_ID.ApplyFriends.req), "payload": JSON.stringify(payload) };

            worker.port.postMessage(json)

        } else {
            input.focus();
        }
    });

    if (!worker) {
        worker = new SharedWorker('assets/assets/js/websocket.js');
        worker.port.onmessage = function (e) {
            // 好友申请回发
            if (MESSAGE_ID.ApplyFriends.res == e.data.msg_id) {
                console.log(e.data.payload)
            }

            // 对方同意好友申请
            if (MESSAGE_ID.ApplyFriends.relay == e.data.msg_id) {
                console.log(e.data.payload)
            }
        };
    }

});

// 发送图片
// <div class="message-item">
// <div class="message-avatar">
//     <figure class="avatar">
//         <img src="assets/assets/jquery/chat/dist/media/img/man_avatar3.jpg"
//             class="rounded-circle" alt="image">
//     </figure>
//     <div>
//         <h5>Byrom Guittet</h5>
//         <div class="time">10:43 AM</div>
//     </div>
// </div>
// <figure>
//     <img src="dist/media/img/image1.jpg" class="w-25 img-fluid rounded" alt="image">
// </figure>
// </div>

// 中间顶格显示
// <div class="message-item messages-divider sticky-top" data-label="Today"></div>

// 标记新消息分段
// <div class="message-item messages-divider" data-label="1 message unread"></div>
