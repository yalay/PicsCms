{{define "ad"}}
<!DOCTYPE html>
<html>
<head>
    {{template "top"}}
    <title>{{.title}} - {{.keywords}}</title>
	<meta name="keywords" content="{{.keywords}}">
	<meta http-equiv="refresh" content="8;url={{.nextUrl}}">
</head>
<body class="home blog body_top">
    {{template "header" .}}
    <div class="main">
        <div class="main_inner">
            <div class="main_left" style="width:100%">
                <div class="item_title">
                    <h1> {{.title}}(<span>赞助商页面 8秒后自动退出</span>)</h1>
	                <div class="single-cat"><span><a href="/">首页</a> / </span> <a href="{{cateUrl .cEngName}}" rel="category tag">{{.cName}}</a><a style="float: right; color: red" href="{{.nextUrl}}">跳过赞助商页面</a></div>
                </div>
                <div class="content" id="content">
                    <div style="margin: 30px auto; width: 600px; vertical-align: middle; display: block;">
	                    <script type="text/javascript" data-idzone="2847708" src="https://ads.exosrv.com/nativeads.js"></script>
                    </div>
                </div>
                </div>
                <section class="single-post-comment">
                    <div class="single-post-comment-reply" id="respond" >
                    </div>
                </section>
            </div>
        </div>
    </div>
    {{template "footer" .}}
</body>
</html>
{{end}}
