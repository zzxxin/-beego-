<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>私聊聊天室</title>
    <!-- 引入 Bootstrap 样式 -->
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
    <style>
        body {
            padding: 20px;
        }
        #user-list {
            height: 500px;
            overflow-y: auto;
        }
        #chat-box {
            height: 400px;
            overflow-y: auto;
            border: 1px solid #ddd;
            padding: 10px;
        }
        .chat-message {
            margin-bottom: 15px;
        }
        .chat-message .username {
            font-weight: bold;
        }
        .chat-message .message {
            margin-left: 10px;
        }
    </style>
</head>
<body>

<div class="container">
    <div class="row">
        <!-- 左侧用户列表 -->
        <div class="col-md-4">
            <h5>在线用户</h5>
            <ul id="user-list" class="list-group">
                {{range .user_list}}
                <li class="list-group-item user-item" data-user-id="{{.ID}}" data-username="{{.UserName}}">
                    {{.UserName}}
                </li>
                {{end}}
            </ul>
        </div>

        <!-- 右侧聊天框 -->
        <div class="col-md-8">
            <h5 id="chat-with">与群聊对话</h5>
            <div id="chat-box"></div>

            <!-- 输入框和发送按钮 -->
            <div class="input-group">
                <input type="text" id="message-input" class="form-control" placeholder="输入消息">
                <button class="btn btn-primary" id="send-btn">发送</button>
            </div>
        </div>
    </div>
</div>

<!-- 引入 jQuery 和 Bootstrap JS -->
<script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"></script>

<script>
    var socket;
    var targetUserId = 0;  // 默认是群聊模式
    var username = "{{ .user_info.UserName }}";  // 当前用户的用户名

    // 建立 WebSocket 连接
    function connectWebSocket() {
        socket = new WebSocket("ws://127.0.0.1:8080/ws");

        // 接收消息
        socket.onmessage = function(event) {
            var msg = JSON.parse(event.data);

            if (msg.type === "message") {
                if (msg.userId === "{{ .user_info.ID }}") {
                    return;
                } else {
                    $('#chat-box').append('<div class="chat-message text-start"><span class="username text-success">' + msg.username + ':</span><span class="message">' + msg.message + '</span></div>');
                }
            } else if (msg.type === "system") {
                $('#chat-box').append('<div class="chat-message text-muted text-center"><span>' + msg.message + '</span></div>');
            } else if (msg.type === "online_users") {
                try {
                    var onlineData = JSON.parse(msg.message);  // 解析message内容
                    var bindIds = onlineData.bind_ids || [];   // 检查bind_ids是否存在
                    updateUserList(bindIds);                   // 更新用户列表
                } catch (error) {
                    console.error("解析在线用户消息时出错: ", error);
                }
            }

            $('#chat-box').scrollTop($('#chat-box')[0].scrollHeight);
        };


        socket.onclose = function() {
            console.log("连接关闭，尝试重连...");
            setTimeout(connectWebSocket, 1000);  // 连接断开后尝试重连
        };
    }

    // 发送消息
    $('#send-btn').on('click', function() {
        var message = $('#message-input').val();
        if (message && socket.readyState === WebSocket.OPEN) {
            // 构造发送的消息
            var msg = {
                userId: "{{ .user_info.ID }}", // 当前用户ID
                username: username,
                message: message,
                targetUserId: targetUserId  // 发送目标用户ID
            };

            // 发送消息到 WebSocket 服务器
            socket.send(JSON.stringify(msg));

            // 清空输入框
            $('#message-input').val('');

            // 立即将自己的消息显示在聊天框中（右对齐）
            $('#chat-box').append('<div class="chat-message text-end"><span class="username text-primary">' + username + ':</span><span class="message">' + message + '</span></div>');

            // 保持聊天框滚动到最新消息
            $('#chat-box').scrollTop($('#chat-box')[0].scrollHeight);
        }
    });


    // 选择用户进行私聊
    $('.user-item').on('click', function() {
        var userId = $(this).data('user-id');
        var selectedUsername = $(this).data('username');
        targetUserId = userId;  // 设置私聊的目标用户ID

        $('#chat-with').text('与 ' + selectedUsername + ' 私聊中');
        $('#chat-box').empty();  // 清空聊天记录（可选）
    });

    // 初始化 WebSocket 连接
    connectWebSocket();

    // 更新在线用户列表
    function updateUserList(bindIds) {
        $('.user-item').each(function() {
            var userId = $(this).data('user-id');
            if (bindIds.includes(userId)) {
                $(this).addClass('list-group-item-success');  // 在线用户显示绿色背景
            } else {
                $(this).removeClass('list-group-item-success');
            }
        });
    }
</script>

</body>
</html>