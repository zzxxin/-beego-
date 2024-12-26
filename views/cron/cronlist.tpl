{{ template "common/header.tpl" . }}
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title d-flex " style="height: 80px;">
                    <h5>Cron管理（<span style="color: #0000cc">修改cron状态 最晚10秒后执行最新变更状态 检测时间为5秒一次</span>）</h5>
                    <div class="ibox-tools align-items-center">
                        <a data-toggle="modal" class="btn btn-primary" href="/add_cron">新增Cron</a>
                        {{ if eq .cron_status 1 }}
                        <a onclick='hidden_compose("cron_status",  0)' href="#" class="btn btn-success btn-outline">停用</a>
                        {{ else }}
                        <a onclick='hidden_compose("cron_status",  1)' href="#" class="btn btn-success btn-outline">启用</a>
                        {{ end }}
                    </div>
                </div>


                <div class="ibox-content">
                    <table class="table table-bordered">
                        <thead>
                        <tr>
                            <th class="center">id</th>
                            <th class="center">任务名称</th>
                            <th class="center">Cron 表达式</th>
                            <th class="center">任务状态</th>
                            <th class="center">任务描述</th>
                            <th class="center">操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{ if .cron_list }}
                        {{ range .cron_list }}
                        <tr class="gradeA">
                            <td>{{ .ID }}</td>
                            <td>{{ .TaskName }}</td>
                            <td>{{ .CronExpression }}</td>
                            <td>{{ if eq .TaskStatus 1 }}开启{{ else }}关闭{{ end }}</td>
                            <td>{{ .TaskDesc }}</td>
                            <td>
                                {{ if eq .TaskStatus 1 }}
                                <a onclick='hidden_compose_status("{{ .ID }}", "{{ .TaskName }}", 2)' href="#" class="btn btn-success btn-outline">停用</a>
                                {{ else }}
                                <a onclick='hidden_compose_status("{{ .ID }}", "{{ .TaskName }}", 1)' href="#" class="btn btn-success btn-outline">启用</a>
                                {{ end }}
                                <a href="/up_cron?id={{ .ID }}" class="btn btn-success btn-outline">编辑Cron</a>
                            </td>
                        </tr>
                        {{ end }}
                        {{ end }}
                        </tbody>
                    </table>
                    <div class="panel-footer clearfix">
                        {{ if .Pagination }}
                        {{ PageLinks .Pagination }}
                        {{ end }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<script>
    function hidden_compose(type, status) {
        var msg = status == 0 ? '停用' : '启用';
        if (confirm('是否要' + msg + 'Cron ？')) {
            location.href = "/cron_set?type=" + type + "&status=" + status;
        }
    }

    function hidden_compose_status(id, cron_name, status) {
        var msg = status == 2 ? '停用' : '启用';
        if (confirm('是否要' + msg + '计划任务：' + cron_name + '？')) {
            location.href = "/up_cron_status?id=" + id + "&status=" + status;
        }
    }

</script>
</body>

</html>