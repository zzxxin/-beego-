{{ template "common/header.tpl" . }}
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title">
                    <h5> 用户【{{.result.user_name}}】分配角色</h5>
                </div>
                <div class="ibox-content">
                    <form method="post" class="form-horizontal" action="/allot_role">
                    <input type="hidden" name="user_id" value="{{.result.user_id}}">
                    <div class="form-group">
                        <label class="col-sm-2 control-label">请选择分配的角色：</label>
                        <div class="col-sm-10">
                            {{range .result.all_role}}
                            <label class="checkbox-inline i-checks">
                                <input type="checkbox" name="bind_checked[]" {{ if inArray (str .ID) $.result.bind_role }} checked="checked"{{ end }} value="{{ .ID }}">&nbsp;&nbsp;{{ .RoleName }}
                            </label>
                            {{end}}
                        </div>
                    </div>
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

<script src="{{ .BaseUrl }}/js/jquery.min.js"></script>
<script src="{{ .BaseUrl }}/js/plugins/iCheck/icheck.min.js"></script>
<script>
    $(document).ready(function(){
        $(".i-checks").iCheck({
            checkboxClass: "icheckbox_square-green",
            radioClass: "iradio_square-green",
        });
    });

    function redirectToPage() {
        window.location.href = "/user_list";
    }
</script>
</body>
</html>