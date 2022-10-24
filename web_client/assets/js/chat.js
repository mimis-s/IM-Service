var worker;

$(function () {

    /**
     * Some examples of how to use features.
     *
     **/

    var SohoExamle = {
        Message: {
            add: function (message, type) {
                var chat_body = $('.layout .content .chat .chat-body');
                if (chat_body.length > 0) {

                    type = type ? type : '';
                    message = message ? message : 'I did not understand what you said!';

                    $('.layout .content .chat .chat-body .messages').append(`<div class="message-item ` + type + `">
                        <div class="message-avatar">
                            <figure class="avatar">
                                <img src="./dist/media/img/` + (type == 'outgoing-message' ? 'women_avatar5.jpg' : 'man_avatar3.jpg') + `" class="rounded-circle">
                            </figure>
                            <div>
                                <h5>` + (type == 'outgoing-message' ? 'Mirabelle Tow' : 'Byrom Guittet') + `</h5>
                                <div class="time">14:50 PM ` + (type == 'outgoing-message' ? '<i class="ti-check"></i>' : '') + `</div>
                            </div>
                        </div>
                        <div class="message-content">
                            ` + message + `
                        </div>
                    </div>`);

                    setTimeout(function () {
                        chat_body.scrollTop(chat_body.get(0).scrollHeight, -1).niceScroll({
                            cursorcolor: 'rgba(66, 66, 66, 0.20)',
                            cursorwidth: "4px",
                            cursorborder: '0px'
                        }).resize();
                    }, 200);
                }
            }
        }
    };

    setTimeout(function () {
        // $('#disconnected').modal('show');
        // $('#call').modal('show');
        // $('#videoCall').modal('show');
        $('#pageTour').modal('show');
    }, 1000);

    $(document).on('submit', '.layout .content .chat .chat-footer form', function (e) {
        e.preventDefault();

        var input = $(this).find('input[type=text]');
        var message = input.val();

        message = $.trim(message);

        if (message) {

            // 读取输入, 发送给服务器
            var payload = {
                Data: {
                    SenderID: 1,
                    ReceiverID: 2,
                    MessageID: 3,
                    MessageType: 1,
                    SendTimeStamp: 1234145,
                    MessageStatus: 1,
                    Data: message,
                },
            }

            var json = { "msg_id": String(MESSAGE_ID.Chat.req), "payload": JSON.stringify(payload) };

            worker.port.postMessage(json)

            // SohoExamle.Message.add(message, 'outgoing-message');
            input.val('');

            // setTimeout(function () {
            //     SohoExamle.Message.add();
            // }, 1000);
        } else {
            input.focus();
        }
    });

    $(document).on('click', '.layout .content .sidebar-group .sidebar .list-group-item', function () {
        if (jQuery.browser.mobile) {
            $(this).closest('.sidebar-group').removeClass('mobile-open');
        }

        // 消除新消息提示
        $(this).find("div[class='new-message-count']").remove()

        // 把当前item状态设置为open-chat
        function hasClass(element, cName) {
            return element.className.match(RegExp('(\\s|^)' + cName + '(\\s|$)')); //使用正则检测是否有相同的样式
        }
        if (!hasClass(this, 'open-chat')) {
            this.className += ' open-chat'
        }
        // 其他的item设置为常态
        var listItem = this.parentNode.children
        var listLenght = listItem.length
        for (var i = 0; i < listLenght; i++) {
            if (listItem[i] !== this) {
                //利用正则捕获到要删除的样式的名称，然后把他替换成一个空白字符串，就相当于删除了
                listItem[i].className = listItem[i].className.replace(RegExp('(\\s|^)' + 'open-chat' + '(\\s|$)'), '');
            }
        }
    });

    if (!worker) {
        worker = new SharedWorker('assets/assets/js/websocket.js');
        worker.port.onmessage = function (e) {
            if (MESSAGE_ID.Chat.res == e.data.msg_id) {
                console.log(e.data.payload)
                SohoExamle.Message.add(e.data.payload.Data.Data, 'outgoing-message');
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
