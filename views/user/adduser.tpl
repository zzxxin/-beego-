{{ template "common/header.tpl" . }}
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{if eq .act "update"}}修改{{else}}新增{{end}}用户</title>
    <link rel="stylesheet" href="{{.BaseUrl}}/css/bootstrap.min.css">
</head>
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title">
                    <h5>{{if eq .act "update"}}修改{{else}}新增{{end}}用户</h5>
                </div>
                <div class="ibox-content">
                    <form method="post" action="{{if eq .act "add"}}/user_add{{else}}/user_edit{{end}}" class="form-horizontal">
                    <input type="hidden" name="user_id" value="{{if .user}}{{.user.ID}}{{end}}">
                    <div class="form-group">
                        <label class="col-sm-3 control-label">用户名</label>
                        <div class="col-sm-6">
                            <input id="lastname" name="user_name" class="form-control" type="text" aria-required="true" value="{{if .user}}{{.user.UserName}}{{end}}">
                        </div>
                    </div>
                    {{if eq .act "add"}}
                    <div class="form-group">
                        <label class="col-sm-3 control-label">密码</label>
                        <div class="col-sm-6">
                            <input id="password" name="passwd" class="form-control" type="password" required>
                        </div>
                    </div>
                    <div class="form-group">
                        <label class="col-sm-3 control-label">确认密码：</label>
                        <div class="col-sm-6">
                            <input id="confirm_password" name="confirm_password" class="form-control" type="password" required>
                            <span class="help-block m-b-none"><i class="fa fa-info-circle"></i> 请再次输入您的密码</span>
                        </div>
                    </div>
                    {{end}}
                    <div class="form-group">
                        <label class="col-sm-3 control-label">真实姓名</label>
                        <div class="col-sm-6">
                            <input id="real_name" name="real_name" class="form-control" type="text" aria-required="true" value="{{if .user}}{{.user.RealName}}{{end}}">
                        </div>
                    </div>
                    <div class="form-group">
                        <label class="col-sm-3 control-label">手机号</label>
                        <div class="col-sm-6">
                            <input type="text" class="form-control w250" name="mobile" value="{{if .user}}{{.user.Mobile}}{{end}}" mod="isempty" msg="权限名称不能为空!">
                        </div>
                    </div>
                    <div class="form-group">
                        <label class="col-sm-3 control-label">账号状态</label>
                        <div class="col-sm-6">
                            <div class="radio radio-info radio-inline">
                                <input type="radio" id="inlineRadio1" value="1" name="status" {{if .user}}{{if eq .user.Status 1}}checked{{end}}{{end}}>
                                <label for="inlineRadio1"> 启用 </label>
                            </div>
                            <div class="radio radio-inline">
                                <input type="radio" id="inlineRadio2" value="2" name="status" {{if .user}}{{if eq .user.Status 2}}checked{{end}}{{end}}>
                                <label for="inlineRadio2"> 停用 </label>
                            </div>
                        </div>
                    </div>
                    <div class="form-group">
                        <label class="col-sm-3 control-label">是否为超级管理员</label>
                        <div class="col-sm-6">
                            <div class="radio radio-info radio-inline">
                                <input type="radio" id="inlineRadio11" value="Y" name="is_super" {{if .user}}{{if eq .user.IsSuper "Y"}}checked{{end}}{{end}}>
                                <label for="inlineRadio1"> 是 </label>
                            </div>
                            <div class="radio radio-inline">
                                <input type="radio" id="inlineRadio22" value="N" name="is_super" {{if .user}}{{if eq .user.IsSuper "N"}}checked{{end}}{{end}}>
                                <label for="inlineRadio2"> 否 </label>
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

<script src="{{.BaseUrl}}/js/jquery.min.js?v=2.1.4"></script>
<script src="{{.BaseUrl}}/js/bootstrap.min.js?v=3.3.5"></script>
<script src="{{.BaseUrl}}/js/content.min.js?v=1.0.0"></script>
<script src="{{.BaseUrl}}/js/plugins/validate/jquery.validate.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/validate/messages_zh.min.js"></script>
<script src="{{.BaseUrl}}/js/demo/form-validate-demo.min.js"></script>

<script>
    function redirectToPage() {
        window.location.href = "/user_list";
    }
</script>
<script>
    $(document).ready(function() {
        // 自定义验证规则：验证密码和确认密码是否一致
        $.validator.addMethod("checkPasswordMatch", function(value, element) {
            var password = $("#password").val();
            return password === value;
        }, "两次输入的密码不一致");

        // 表单验证规则
        $("form").validate({
            rules: {
                passwd: {
                    required: true,
                    minlength: 6
                },
                confirm_password: {
                    required: true,
                    minlength: 6,
                    checkPasswordMatch: true
                }
            },
            messages: {
                passwd: {
                    required: "请输入密码",
                    minlength: "密码长度不能少于6个字符"
                },
                confirm_password: {
                    required: "请再次输入密码",
                    minlength: "密码长度不能少于6个字符"
                }
            },
            errorPlacement: function(error, element) {
                error.addClass("text-danger");  // 添加红色提示
                error.insertAfter(element);     // 将错误消息插入到元素后面
            },
            highlight: function(element) {
                $(element).closest('.form-group').addClass('has-error'); // 标红输入框
            },
            unhighlight: function(element) {
                $(element).closest('.form-group').removeClass('has-error'); // 移除标红
            }
        });

        // 当确认密码框的光标移出时，立即验证两次密码是否一致
        $("#confirm_password").on('blur', function() {
            $(this).valid(); // 触发 jQuery Validate 插件的即时校验
        });
    });
</script>

</body>
</html>