{{ template "common/header.tpl" . }}
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title">
                    <h5>{{ if eq .act "update" }} 修改{{ else }} 新增{{ end }}Cron</h5>
                </div>
                <div class="ibox-content">
                    <form method="post" {{ if eq .act "add" }} action="/add_cron" {{ else }} action="/up_cron" {{ end }} class="form-horizontal">
                    <input type="hidden" name="cron_id" value="{{ if .cron_info }}{{ .cron_info.ID }}{{ end }}">

                    <div class="form-group">
                        <label class="col-sm-3 control-label">任务名称(执行时候的标识)</label>
                        <div class="col-sm-6">
                            <input type="text" class="form-control w250" name="task_name" value="{{if .cron_info}}{{ .cron_info.TaskName }}{{end}}" mod="isempty" msg="任务名称不能为空!">
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-3 control-label">Cron 表达式（秒、分、时、日、月、星期）</label>
                        <div class="col-sm-6">
                            <input type="text" class="form-control w250" name="cron_expression" value="{{if .cron_info}}{{ .cron_info.CronExpression }}{{end}}" mod="isempty" msg="表达式不能为空!">
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-3 control-label">任务描述</label>
                        <div class="col-sm-6">
                            <input type="text" class="form-control w250" name="task_desc" value="{{if .cron_info}}{{ .cron_info.TaskDesc }}{{end}}" mod="isempty" msg="任务描述不能为空!">
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-3 control-label">是否开启</label>
                        <div class="col-sm-6">
                            <div class="radio radio-info radio-inline">
                                <input type="radio" id="inlineRadio1" value="1" name="task_status" {{if .cron_info}}{{ if eq .cron_info.TaskStatus 1 }}checked{{ end }}{{ end }}>
                                <label for="inlineRadio1"> 启用 </label>
                            </div>
                            <div class="radio radio-inline">
                                <input type="radio" id="inlineRadio2" value="2" name="task_status" {{if .cron_info}}{{ if eq .cron_info.TaskStatus 2 }}checked{{ end }}{{ end }}>
                                <label for="inlineRadio2"> 禁用 </label>
                            </div>
                        </div>
                    </div>

                    <div class="form-group">
                        <div class="col-sm-4 col-sm-offset-2">
                            <button class="btn btn-primary" type="submit">保存内容</button>
                            <button class="btn btn-white" type="button" onclick="redirectToPage()">取消</button>
                        </div>
                    </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

<script>
    function redirectToPage() {
        window.location.href = "/cron_list";
    }
</script>
</body>
</html>