<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>后台管理 - 登录</title>
    <link rel="shortcut icon" href="favicon.ico">
    <link href="{{.BaseUrl}}/css/bootstrap.min.css?v=3.3.5" rel="stylesheet">
    <link href="{{.BaseUrl}}/css/font-awesome.min.css?v=4.4.0" rel="stylesheet">
    <link href="{{.BaseUrl}}/css/animate.min.css" rel="stylesheet">
    <link href="{{.BaseUrl}}/css/style.min.css?v=4.0.0" rel="stylesheet">
    <base target="_blank">
    <!--[if lt IE 8]>
    <meta http-equiv="refresh" content="0;ie.html" />
    <![endif]-->
    <script>
        if (window.top !== window.self) {
            window.top.location = window.location;
        }
    </script>
</head>

<body class="gray-bg">

<div class="middle-box text-center loginscreen animated fadeInDown">
    <div>
        <div>
            <h1 class="logo-name">H+</h1>
        </div>
        <h3>欢迎使用 后台管理</h3>

        <form class="m-t" role="form">
            <div class="form-group">
                <input type="text" class="form-control" placeholder="用户名 超管账号 ：admin" id="user_name" required="">
            </div>
            <div class="form-group">
                <input type="password" class="form-control" id="passwd" placeholder="密码 超管密码 ：admin" required="">
            </div>
            <button type="button" class="btn btn-primary block full-width m-b" id="submit">登 录</button>
            <p class="text-muted text-center">
                <a href="#"><small>忘记密码了？</small></a> | <a href="#">注册一个新账号</a>
            </p>
        </form>
    </div>
</div>

<script src="{{.BaseUrl}}/js/jquery.min.js?v=2.1.4"></script>
<script src="{{.BaseUrl}}/js/bootstrap.min.js?v=3.3.5"></script>

<script>
    $('#submit').on('click', function() {
        var url = '/do_login';
        var params = {
            user_name: $("#user_name").val(),
            passwd: $("#passwd").val()
        };

        $.post(url, params, function(res) {
            if (res.code == 200) {
                console.log("{{ .BaseURL }}")
                window.location.href = "/";
                return;
            }
            alert('错误：[' + res.code + '] ' + (res.msg || '网络错误'));
        }, 'json');
    });
</script>

</body>
</html>