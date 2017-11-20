{{define "error"}}
<!DOCTYPE html>
<html>
<head>
    {{template "top"}}
    <title>404 {{.error}} - {{.webName}}</title>
    <meta name="keywords" content="404">
    <meta name="description" content="404 找不到您要的图集">
</head>
<body class="home blog body_top">
    {{template "header" .}}
    <div class="cat_bg">
        <div class="cat_bg_img" style="background-image:url(/img/tags.png);">
            <div><span style="font-size: 18px;color: #F14141;font-weight: 600;">404</span><br>{{.error}}</div>
        </div>
    </div>
    <!--分类导航-->
    <div class="fl flbg">
        <div class="fl_title"><div class="fl01">以下是推荐给您的相似图集</div></div>
    </div>
    <div class="update_area">
        <div class="update_area_content">
            <ul class="update_area_lists cl">
                {{range .tArticles}}
                {{template "list" .}}
                {{end}}
            </ul>
        </div>
    </div>
    {{template "footer" .}}
</body>
</html>
{{end}}
