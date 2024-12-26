{{ template "common/header.tpl" . }}
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title d-flex " style="height: 60px;">
                    <h5>角色管理</h5>
                    <div class="ibox-tools align-items-center">
                        <a data-toggle="modal" class="btn btn-primary" href="/add_role">新增角色</a>
                    </div>
                </div>

                <div class="ibox-content">

                    <table class="table table-bordered">
                        <thead>
                        <tr>
                            <th class="center">id</th>
                            <th class="center">角色名称</th>
                            <th class="center">操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{ if .role_list }}
                        {{ range .role_list }}
                        <tr class="gradeA">
                            <td>{{ .ID }}</td>
                            <td>{{  .RoleName }}</td>
                            <td>
                                <a href="/role_bind_right?id={{.ID }}" class="btn btn-success btn-outline">分配权限</a>
                                <a href="/up_role?id={{.ID }}" class="btn btn-success btn-outline">编辑角色</a>

                            </td>
                        </tr>
                        {{ end }}
                        {{ end }}
                        </tbody>
                    </table>

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

</html>
