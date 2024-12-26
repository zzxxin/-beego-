<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>信息提示</title>
    <link rel="stylesheet" href="/static/css/bootstrap.min.css">
    <style>
        .modal-body {
            display: flex;
            flex-direction: column;
            justify-content: center;
            align-items: center;
            font-size: 20px; /* 调整字体大小 */
            min-height: 150px; /* 设置最小高度，让内容垂直居中 */
        }
    </style>
</head>
<body>
<div class="modal-dialog">
    <div class="modal-content">
        <div class="modal-header">
            <h3>信息提示</h3>
        </div>
        <div class="modal-body">
            <p style="{{ if eq .code 200 }}color: green{{ else }}color: red{{ end }}">{{ .message }}</p>
        </div>
        <div class="modal-footer">
            <a href="{{ .url }}" class="btn btn-primary">立即返回</a>
        </div>
    </div>
</div>
<script>
    var delay_time = 5;
    function delayAct() {
        delay_time--;

        var url = '{{ .url }}';
        if (delay_time < 0) {
            clearInterval(interval);
            return;
        }

        if (delay_time === 0) {
            clearInterval(interval);
            if (url === '') {
                window.history.go(-1);
                return;
            }
            window.location.href = url;
            return;
        }

        document.getElementById('return_time').innerHTML = delay_time;
    }
    var interval = setInterval(delayAct, 1000);
</script>
</body>
</html>