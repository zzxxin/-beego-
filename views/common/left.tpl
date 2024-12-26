{{ template "common/header.tpl" . }}

<body class="fixed-sidebar full-height-layout gray-bg" style="overflow:hidden">
<div id="wrapper">
    <!--左侧导航开始-->
    <nav class="navbar-default navbar-static-side" role="navigation">
        <div class="nav-close"><i class="fa fa-times-circle"></i></div>
        <div class="sidebar-collapse">
            <ul class="nav" id="side-menu">
                <li class="nav-header">
                    <div class="dropdown profile-element">
                        <a data-toggle="dropdown" class="dropdown-toggle" href="#">
                            <span class="clear">
                                <span class="block m-t-xs"><strong class="font-bold">{{ .user_info.UserName }}</strong></span>
                                <span class="text-muted text-xs block">
                                    {{ if eq .user_info.IsSuper "Y" }}超管{{ else }}普管{{ end }} <b class="caret"></b>
                                </span>
                            </span>
                        </a>
                        <ul class="dropdown-menu animated fadeInRight m-t-xs">
                            <li class="divider"></li>
                            <li><a href="/login_out">安全退出</a></li>
                        </ul>
                    </div>
                    <div class="logo-element">H+</div>
                </li>
                {{ if .UserRight }}
                    {{ range $key, $item := .UserRight.right_list }}
                    <li>
                        <a href="#">
                            <i class="fa fa-home"></i>
                            <span class="nav-label">{{ $key }}</span>
                            <span class="fa arrow"></span>
                        </a>
                        <ul class="nav nav-second-level">
                            {{ range $v := $item }}
                            <li>
                                <a class="J_menuItem" href="/{{ $v.RightLogo }}" data-index="0">{{ $v.RightName }}</a>
                            </li>
                            {{ end }}
                        </ul>
                    </li>
                    {{ end }}
                {{ end }}
            </ul>
        </div>
    </nav>
    <!--左侧导航结束-->

    <!--右侧部分开始-->
    <div id="page-wrapper" class="gray-bg dashbard-1">
        <div class="row border-bottom">
            <nav class="navbar navbar-static-top" role="navigation" style="margin-bottom: 0">
                <div class="navbar-header"><a class="navbar-minimalize minimalize-styl-2 btn btn-primary " href="#"><i class="fa fa-bars"></i></a></div>
            </nav>
        </div>
        <div class="row content-tabs">
            <nav class="page-tabs J_menuTabs">
                <div class="page-tabs-content">
                    <a href="javascript:;" class="active J_menuTab" data-id="/web">首页</a>
                </div>
            </nav>
            <a href="/login_out" class="roll-nav roll-right J_tabExit"><i class="fa fa-sign-out"></i> 退出</a>
        </div>

        <div class="row J_mainContent" id="content-main">
            <iframe class="J_iframe" name="iframe0" width="100%" height="100%" src="/web" frameborder="0" data-id="/web" seamless></iframe>
        </div>
    </div>
</div>

<script src="{{.BaseUrl}}/js/jquery.min.js?v=2.1.4"></script>
<script src="{{.BaseUrl}}/js/bootstrap.min.js?v=3.3.5"></script>
<script src="{{.BaseUrl}}/js/plugins/metisMenu/jquery.metisMenu.js"></script>
<script src="{{.BaseUrl}}/js/plugins/slimscroll/jquery.slimscroll.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/layer/layer.min.js"></script>
<script src="{{.BaseUrl}}/js/hplus.min.js?v=4.0.0"></script>
<script type="text/javascript" src="{{.BaseUrl}}/js/contabs.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/pace/pace.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/jeditable/jquery.jeditable.js"></script>
<script src="{{.BaseUrl}}/js/plugins/dataTables/jquery.dataTables.js"></script>
<script src="{{.BaseUrl}}/js/plugins/dataTables/dataTables.bootstrap.js"></script>
<script src="{{.BaseUrl}}/js/content.min.js?v=1.0.0"></script>
<script type="text/javascript" src=" " charset="UTF-8"></script>
<script src="{{.BaseUrl}}/js/plugins/iCheck/icheck.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/sweetalert/sweetalert.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/chosen/chosen.jquery.js"></script>
<script src="{{.BaseUrl}}/js/plugins/jsKnob/jquery.knob.js"></script>
<script src="{{.BaseUrl}}/js/plugins/jasny/jasny-bootstrap.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/datapicker/bootstrap-datepicker.js"></script>
<script src="{{.BaseUrl}}/js/plugins/prettyfile/bootstrap-prettyfile.js"></script>
<script src="{{.BaseUrl}}/js/plugins/nouslider/jquery.nouislider.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/switchery/switchery.js"></script>
<script src="{{.BaseUrl}}/js/plugins/ionRangeSlider/ion.rangeSlider.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/iCheck/icheck.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/colorpicker/bootstrap-colorpicker.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/clockpicker/clockpicker.js"></script>
<script src="{{.BaseUrl}}/js/demo/form-advanced-demo.min.js"></script>

<script src="{{.BaseUrl}}/js/plugins/validate/jquery.validate.min.js"></script>
<script src="{{.BaseUrl}}/js/plugins/validate/messages_zh.min.js"></script>
<script src="{{.BaseUrl}}/js/demo/form-validate-demo.min.js"></script>

<style>
    .copyrights{text-indent:-9999px;height:0;line-height:0;font-size:0;overflow:hidden;}
</style>

</body>
</html>


        