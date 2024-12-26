{{ template "common/header.tpl" . }}
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title">
                    <h5>{{ if eq .act "update" }} 修改{{ else }} 新增{{ end }}权限</h5>
                </div>
                <div class="ibox-content">
                    <form method="post" {{ if eq .act "add" }} action="/add_right" {{ else }} action="/up_right" {{ end }} class="form-horizontal">
                    <input type="hidden" name="right_id" value="{{ if .right_info }}{{ .right_info.ID }}{{ end }}">

                    <div class="form-group">
                        <label class="col-sm-3 control-label">目录名称</label>
                        <div class="col-sm-6">
                            <input type="text" class="form-control w250" name="group_name" value="{{if .right_info}}{{ .right_info.GroupName }}{{end}}" mod="isempty" msg="角色名称不能为空!">
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-3 control-label">权限标识（真实路由）</label>
                        <div class="col-sm-6">
                            <input type="text" class="form-control w250" name="right_logo" value="{{if .right_info}}{{ .right_info.RightLogo }}{{end}}" mod="isempty" msg="权限标识不能为空!">
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-3 control-label">权限名称</label>
                        <div class="col-sm-6">
                            <input type="text" class="form-control w250" name="right_name" value="{{if .right_info}}{{ .right_info.RightName }}{{end}}" mod="isempty" msg="权限名称不能为空!">
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-3 control-label">是否校验权限</label>
                        <div class="col-sm-6">
                            <div class="radio radio-info radio-inline">
                                <input type="radio" id="inlineRadio1" value="1" name="is_right" {{if .right_info}}{{ if eq .right_info.IsRight 1 }}checked{{ end }}{{ end }}>
                                <label for="inlineRadio1"> 是 </label>
                            </div>
                            <div class="radio radio-inline">
                                <input type="radio" id="inlineRadio2" value="2" name="is_right" {{if .right_info}}{{ if eq .right_info.IsRight 2 }}checked{{ end }}{{ end }}>
                                <label for="inlineRadio2"> 否 </label>
                            </div>
                        </div>
                    </div>

                    <div class="form-group">
                        <label class="col-sm-3 control-label">是否菜单</label>
                        <div class="col-sm-6">
                            <div class="radio radio-info radio-inline">
                                <input type="radio" id="inlineRadio11" value="1" name="is_menu" {{if .right_info}}{{ if eq .right_info.IsMenu 1 }}checked{{ end }}{{ end }}>
                                <label for="inlineRadio11"> 是 </label>
                            </div>
                            <div class="radio radio-inline">
                                <input type="radio" id="inlineRadio22" value="2" name="is_menu" {{if .right_info}}{{ if eq .right_info.IsMenu 2 }}checked{{ end }}{{ end }}>
                                <label for="inlineRadio22"> 否 </label>
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
        window.location.href = "/right_list";
    }
</script>
</body>
</html>