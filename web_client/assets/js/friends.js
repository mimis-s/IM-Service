$(function () {

    // 提交好友申请
    $(document).on('click', '#addFriends button.btn', function (e) {
        console.log("提交好友申请")
        e.preventDefault();

        var applyFriendID = $('#addFriends').find('input[id=applyFriendID]').val();
        applyFriendID = $.trim(applyFriendID);

        if (applyFriendID) {

            // 读取输入, 发送给服务器
            var payload = {
                ApplyFriendsID: Number(applyFriendID),
            }

            var json = { "msg_id": String(MESSAGE_ID.ApplyFriends.req), "payload": JSON.stringify(payload) };

            worker.port.postMessage(json)

        } else {
            input.focus();
        }
    });
});

function InitFriends(e) {
    // 好友申请回发
    if (MESSAGE_ID.ApplyFriends.res == e.data.msg_id) {
        console.log("好友申请回发" + e.data.payload)
    }

    // 接收对方好友申请
    if (MESSAGE_ID.ApplyFriends.relay == e.data.msg_id) {
        console.log("接收对方好友申请" + e.data.payload)
    }

    // 获取好友列表
    if (MESSAGE_ID.GetFriendsList.res == e.data.msg_id) {
        console.log("获取好友列表" + e.data.payload)
    }
}


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
