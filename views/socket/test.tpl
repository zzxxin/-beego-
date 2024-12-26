{{ template "common/header.tpl" . }}
<style>
    .left {
        text-align: left;
    }

    .right {
        text-align: right;
    }

    .chat-message {
        background-color: #f3f3f4;
        padding: 10px;
        border-radius: 5px;
        display: inline-block;
        max-width: 70%;
    }

    .right .chat-message {
        background-color: #a3d063;
    }

    .author-name {
        font-weight: bold;
    }

    .chat-date {
        color: #888;
        font-size: 0.85em;
        margin-left: 10px;
    }
</style>
<body class="gray-bg">
<div class="wrapper wrapper-content  animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox chat-view">

                <div class="ibox-title">
                    <small class="pull-right text-muted">最新消息：2024-09-10 18:39:23</small> 群聊窗口
                </div>
                <div class="ibox-content">
                    <div class="row">
                        <div class="col-md-9">
                            <div class="chat-discussion" id="chat-discussion">
                                <!-- 动态插入聊天消息的容器 -->
                            </div>
                        </div>
                        <div class="col-md-3">
                            <div class="chat-users">
                                <div class="users-list">
                                    <div class="chat-user">
                                        <div class="online-users">
                                            <h4>在线用户</h4>
                                            <ul id="online-users">
                                                <!-- 在线用户列表 -->
                                            </ul>
                                        </div>
                                    </div>
                                </div>

                            </div>
                        </div>
                    </div>

                    <div class="row">
                        <div class="col-sm-9">
                            <div class="chat-message-form">
                                <div class="form-group">
                                    <textarea id="message-input" class="form-control message-input" name="message"
                                              placeholder="输入消息内容，按回车键发送"></textarea>
                                </div>
                            </div>
                        </div>
                    </div>

                </div>

            </div>
        </div>
    </div>
</div>

<!-- 引入必要的JS库 -->
<script src="{{.BaseUrl}}/js/jquery.min.js?v=2.1.4"></script>
<script src="{{.BaseUrl}}/js/bootstrap.min.js?v=3.3.5"></script>
<script src="{{.BaseUrl}}/js/content.min.js?v=1.0.0"></script>

<script type="text/javascript">
    let socket;
    const chatDiscussion = $('#chat-discussion');
    const onlineUsersList = $('#online-users');
    const username = "{{ .user_info.UserName }}";  // 获取当前用户名
    const userId = "{{ .user_info.ID }}";  // 获取当前用户ID

    $(document).ready(function () {
        // 初始化 WebSocket 连接
        initWebSocket();

        // 页面加载完成后初始化聊天记录
        initChatHistory();

        // 按回车发送消息
        $('#message-input').keypress(function (e) {
            if (e.which === 13 && !e.shiftKey) { // 检查是否是回车键，并排除shift+enter的情况
                e.preventDefault(); // 禁用默认的回车行为
                sendMessage();
            }
        });
    });

    // 初始化聊天记录
    function initChatHistory() {
        const channelId = ""; // 当前群聊的是空
        $.ajax({
            url: '/chat/get_messages',
            type: 'GET',
            data: { channel_id: channelId },
            success: function (data) {
                if (data.error) {
                    console.error("获取聊天记录失败: ", data.error);
                } else {
                    // 遍历返回的消息数据并渲染
                    data.forEach(function (message) {
                        const isSelf = message.user_id == userId;  // 判断是否是自己发送的消息
                        displayMessage(message.username, message.content, message.timestamp, isSelf);
                    });
                }
            },
            error: function (xhr, status, error) {
                console.error("AJAX 请求错误: ", error);
            }
        });
    }

    // 初始化 WebSocket 连接
    function initWebSocket() {
        const wsUrl = `ws://127.0.0.1:8080/ws`;  // 替换为实际的 WebSocket URL
        socket = new WebSocket(wsUrl);
        socket.onopen = function () {
            console.log("WebSocket 连接成功");
        };

        // 接收消息时处理
        socket.onmessage = function (event) {
            const message = JSON.parse(event.data);
            console.log(message);
            if (message.type === "message") {
                const isSelf = message.userId === userId;  // 判断是否是自己发送的消息
                displayMessage(message.username, message.message, new Date().toLocaleTimeString(), isSelf);
            } else if (message.type === "online_users") {
                const onlineUsers = JSON.parse(message.message);
                updateOnlineUsers(onlineUsers);
            } else if (message.type === "system") {
                displayMessage("系统", message.message, new Date().toLocaleTimeString(), false);
            }
        };

        socket.onclose = function () {
            console.log("WebSocket 连接已关闭");
        };

        socket.onerror = function (error) {
            console.log("WebSocket 出现错误: ", error);
        };
    }

    // 发送消息逻辑
    function sendMessage() {
        var message = $('#message-input').val().trim();
        if (message === "") {
            alert("消息不能为空！");
            return;
        }

        // 定义发送消息体
        const msgObj = {
            username: username,
            userId: userId,
            message: message,
            type: "message",  // 设置消息类型为普通聊天消息
            targetUserId: 0,  // 群聊的目标ID是0
            chanel_id: ""  // 群聊的目标ID是0
        };

        // 通过 WebSocket 发送消息
        socket.send(JSON.stringify(msgObj));

        // 在界面上显示自己的消息 (isSelf = true)
        displayMessage(username, message, new Date().toLocaleTimeString(), true);

        // 清空输入框并自动滚动到底部
        $('#message-input').val("");
        chatDiscussion.scrollTop(chatDiscussion[0].scrollHeight);
    }

    // 更新在线用户列表
    function updateOnlineUsers(userList) {
        const onlineUsers = userList.bind_ids;  // 在线用户 ID 列表
        const allUsers = userList.user_list;    // 全部用户信息列表

        onlineUsersList.empty();  // 清空当前列表

        // 先渲染在线用户
        onlineUsers.forEach(function (onlineId) {
            const user = allUsers.find(user => user.id === onlineId);
            if (user) {
                onlineUsersList.append(`
                <div class="chat-user" data-userid="${user.id}">
                    <span class="pull-right label label-primary">在线</span>
                    <div class="chat-user-name">
                        <a href="#">${user.user_name}</a>
                    </div>
                </div>
            `);
            }
        });

        // 渲染离线用户
        allUsers.forEach(function (user) {
            if (!onlineUsers.includes(user.id)) {
                onlineUsersList.append(`
                <div class="chat-user">
                    <span class="pull-right label label-default">离线</span>
                    <div class="chat-user-name">
                        <a href="#">${user.user_name}</a>
                    </div>
                </div>
            `);
            }
        });

        // 如果没有在线用户
        if (onlineUsers.length === 0) {
            onlineUsersList.append('<li>暂无在线用户</li>');
        }
    }

    // 显示消息
    function displayMessage(username, message, time, isSelf) {
        const messagePosition = isSelf ? 'right' : 'left';  // 根据发送者调整消息位置
        const newMessage = `
    <div class="${messagePosition}">
        <div class="author-name">
            ${username}
            <small class="chat-date">
                ${time}
            </small>
        </div>
        <div class="chat-message ${isSelf ? 'active' : ''}">
            ${message}
        </div>
    </div>
`;
        chatDiscussion.append(newMessage);
        // 确保在元素更新后滚动
        setTimeout(() => {
            chatDiscussion.scrollTop(chatDiscussion[0].scrollHeight);
        }, 0);
    }

    // 给用户那里添加一个点击事件
    onlineUsersList.on('click', '.chat-user', function () {
        const targetUserId = $(this).data('userid'); // 获取目标用户ID
        createOrGetChannel(targetUserId);
    });

    function createOrGetChannel(targetUserId) {
        $.ajax({
            url: '/chat/get_channel',
            type: 'POST',
            data: { member_ids: [targetUserId], is_group: false }, // 私聊模式
            success: function (data) {
                if (data.error) {
                    console.error("创建频道失败: ", data.error);
                } else {
                    // 成功获取频道，更新聊天窗口
                    updateChatWindow(data.channel, data.messages);
                }
            },
            error: function (xhr, status, error) {
                console.error("AJAX 请求错误: ", error);
            }
        });
    }

    function updateChatWindow(channel, messages) {
        // 更新聊天窗口标题
        const chatTitle = channel.is_group ? channel.name : `与 ${getUserName(channel.members[1])} 的对话`;
        $('#chat-window-title').text(chatTitle);

        // 清空当前的聊天内容
        const chatMessagesContainer = $('#chat-messages');
        chatMessagesContainer.empty();

        // 渲染新的聊天消息
        messages.forEach(function (message) {
            const messageElement = createMessageElement(message);
            chatMessagesContainer.append(messageElement);
        });

        // 滚动到最新消息
        chatMessagesContainer.scrollTop(chatMessagesContainer[0].scrollHeight);
    }

    function createMessageElement(message) {
        const time = new Date(message.created_at).toLocaleTimeString();
        const userClass = message.sender_id === currentUser.id ? 'chat-message-self' : 'chat-message-other';

        return `
        <div class="chat-message ${userClass}">
            <span class="chat-message-user">${message.sender_name}</span>
            <span class="chat-message-time">${time}</span>
            <div class="chat-message-content">${message.content}</div>
        </div>
    `;
    }

</script>

</body>
</html>