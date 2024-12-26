{{template "common/header.tpl" .}}
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title">
                    <h5> 角色【<span style="color: green">{{ .result.role_name }}</span>】分配权限</h5>
                </div>
                <div class="ibox-content">
                    <form method="post" class="form-horizontal" action="/role_bind_right">
                    <input type="hidden" name="role_id" value="{{ .result.role_id }}">

                    {{range $key, $val := .result.right_list}}
                    <div class="form-group">
                        <label class="col-sm-2 control-label">{{ $key }}</label>
                        <div class="col-sm-10">
                            {{range $item := $val}}
                            <label class="checkbox-inline i-checks">
                                <input type="checkbox" name="bind_checked[]" {{if inArray (str .ID) $.result.bind_right}}checked="checked"{{end}} value="{{ $item.ID }}">&nbsp;{{ $item.RightName }}
                            </label>
                            {{end}}
                        </div>
                    </div>
                    {{end}}

                    <div class="hr-line-dashed"></div>
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

<script src="{{.BaseUrl}}/js/jquery.min.js?v=2.1.4"></script>
<script src="{{.BaseUrl}}/js/plugins/iCheck/icheck.min.js"></script>
<script>
    $(document).ready(function(){$(".i-checks").iCheck({checkboxClass:"icheckbox_square-green",radioClass:"iradio_square-green",})});
    function redirectToPage() {
        window.location.href = "/role_list";
    }
</script>
</body>
</html>