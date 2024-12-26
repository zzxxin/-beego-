{{ template "common/header.tpl" . }}
<body class="gray-bg">
<div class="wrapper wrapper-content animated fadeInRight">
    <div class="row">
        <div class="col-sm-12">
            <div class="ibox float-e-margins">
                <div class="ibox-title d-flex " style="height: 60px;">
                    <h5>权限管理</h5>
                    <div class="ibox-tools align-items-center">
                        <a data-toggle="modal" class="btn btn-primary" href="/add_right">新增权限</a>
                    </div>
                </div>
                <div class="ibox-content">
                    <table class="table table-bordered">
                        <thead>
                        <tr>
                            <th class="center">id</th>
                            <th class="center">所属分组</th>
                            <th class="center">权限名称</th>
                            <th class="center">权限标识(<span style="color: #0000cc">接口路由</span>)</th>
                            <th class="center">操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{ if .right_list }}
                        {{ range .right_list }}
                        <tr class="gradeA">
                            <td>{{ .ID }}</td>
                            <td>{{ .GroupName }}</td>
                            <td>{{ .RightName }}</td>
                            <td>{{ .RightLogo }}</td>
                            <td>
                                <a href="/up_right?id={{ .ID }}" class="btn btn-success btn-outline">编辑权限</a>
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
</body>

</html>