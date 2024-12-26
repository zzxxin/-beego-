{{ template "common/header.tpl" . }}
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title d-flex " style="height: 60px;">
                    <h5>用户管理</h5>
                    <div class="ibox-tools align-items-center">
                        <a data-toggle="modal" class="btn btn-primary" href="/user_add">新增用户</a>
                    </div>
                </div>

                <div class="ibox-content">

                    <table class="table table-bordered">
                        <thead>
                        <tr>
                            <th class="center">id</th>
                            <th class="center">用户名称</th>
                            <th class="center">真实姓名</th>
                            <th class="center">手机号</th>
                            <th class="center">状态</th>
                            <th class="center">是否超管</th>
                            <th class="center">最后登录时间</th>
                            <th class="center">操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{ if .user_list }}
                        {{ range .user_list }}
                        <tr class="gradeA">
                            <td>{{ .ID }}</td>
                            <td>{{  .UserName }}</td>
                            <td>{{  .RealName }}</td>
                            <td>{{  .Mobile }}</td>
                            <td>{{  .StatusName }}</td>
                            <td>{{  .IsSuper }}</td>
                            <td>{{ FormatTime .LastLogin }}</td>
                            <td>
                                <a href="/user_edit?id={{ .ID }}" class="btn btn-success btn-outline">编辑</a>

                                {{ if eq .Status 1 }}
                                <a onclick='hidden_compose("{{ .ID }}", "{{ .UserName }}", 2)' href="#" class="btn btn-success btn-outline">停用</a>
                                {{ else }}
                                <a onclick='hidden_compose("{{ .ID }}", "{{ .UserName }}", 1)' href="#" class="btn btn-success btn-outline">启用</a>
                                {{ end }}

                                <a href="/allot_role?id={{ .ID }}" class="btn btn-success btn-outline">分配角色</a>
                            </td>
                        </tr>
                        {{ end }}
                        {{ end }}
                        </tbody>
                    </table>
                    <!-- 可以在这里添加分页等其他功能 -->
                    <div class="pagination-wrapper">
                        {{ if .Pagination }}
                        {{ PageLinks .Pagination }}
                        {{ end }}
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
</body>
<script>
    function hidden_compose(id, user_name, status) {
        var msg = status == 2 ? '停用' : '启用';
        if (confirm('是否要' + msg + '用户：' + user_name + '？')) {
            location.href = "/user_status?id=" + id + "&status=" + status;
        }
    }
</script>
</html>