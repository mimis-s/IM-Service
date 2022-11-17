var worker;

// 一个新消息sider
var newSiderItem = `<li class="list-group-item">
    <figure class="avatar">
    <span class="avatar-title bg-warning bg-success rounded-circle">
        <i class="fa fa-users"></i>
    </span>
    </figure>
    <div class="users-list-body">
        <div>
            <h5>😍 My Family 😍</h5>
            <p><strong>Maher Ruslandi: </strong>Hello!!!</p>
        </div>
        <div class="users-list-action">
            <small class="text-muted">Yesterday</small>
            <div class="action-toggle">
                <div class="dropdown">
                    <a data-toggle="dropdown" href="#">
                        <i data-feather="more-horizontal"></i>
                    </a>
                    <div class="dropdown-menu dropdown-menu-right">
                        <a href="#" class="dropdown-item">Open</a>
                        <a href="#" data-navigation-target="contact-information"
                            class="dropdown-item">Profile</a>
                        <a href="#" class="dropdown-item">Add to archive</a>
                        <div class="dropdown-divider"></div>
                        <a href="#" class="dropdown-item text-danger">Delete</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</li > `

function newSideChatItem() {
    var group = document.getElementById("sidebar_item_group")
    group.insertAdjacentHTML("beforeEnd", newSiderItem)
}

// 定义worker
$(function () {
    if (!worker) {
        worker = new SharedWorker('assets/assets/js/websocket.js');
        worker.port.onmessage = function (e) {
            InitChat(e)
            InitFriends(e)
        };
    }
})

// 好友消息
function InitFriends(e) {
    // 好友申请回发
    if (MESSAGE_ID.ApplyFriends.res == e.data.msg_id) {
        console.log("好友申请回发" + e.data.payload)
    }

    // 对方同意好友申请
    if (MESSAGE_ID.ApplyFriends.relay == e.data.msg_id) {
        console.log("对方同意好友申请" + e.data.payload)
    }

    console.log("同意好友申请的id:" + MESSAGE_ID.ApplyFriends.relay)

    // 获取好友列表
    if (MESSAGE_ID.GetFriendsList.res == e.data.msg_id) {
        console.log("获取好友列表" + e.data.payload)
    }
}

// // 聊天消息
// function InitChat(e) {
//     // 聊天消息回发
//     if (MESSAGE_ID.Chat.res == e.data.msg_id) {
//         console.log(e.data.payload)
//         SohoExamle.Message.add(e.data.payload.Data.Data, 'outgoing-message');
//     }

//     // 接收对端消息
//     if (MESSAGE_ID.Chat.relay == e.data.msg_id) {
//         console.log(e.data.payload)
//         SohoExamle.Message.add(e.data.payload.Data.Data, '');
//     }
// }
