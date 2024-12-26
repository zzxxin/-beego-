{{ template "common/header.tpl" . }}
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title">
                    <h5>{{if eq .act "update"}} 修改{{else}} 新增{{end}}角色</h5>
                </div>
                <div class="ibox-content">
                    <form method="post" action="{{if eq .act "add"}}/add_role{{else}}/up_role{{end}}" class="form-horizontal">
                    <input type="hidden" name="role_id" value="{{ if .roleinfo }}{{.roleinfo.ID}}{{ end }}">
                    <div class="form-group">
                        <label class="col-sm-3 control-label">角色名称</label>
                        <div class="col-sm-4">
                            <input type="text" class="form-control w250" name="role_name" value="{{if .roleinfo}}{{.roleinfo.RoleName}}{{end}}" mod="isempty" msg="角色名称不能为空!">
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
        window.location.href = "/role_list";
    }
</script>
</body>
</html>