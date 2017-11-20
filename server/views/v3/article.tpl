{{define "article"}}
<!DOCTYPE html>
<html>
<head>
    {{template "top"}}
    <title>{{.title}} - {{.keywords}}</title>
    <meta name="keywords" content="{{.keywords}}">
    <link rel="canonical" href="{{articleUrl .id}}"/>
</head>
<body class="home blog body_top">
    {{template "header" .}}
    <div class="main">
        <div class="main_inner">
            <div class="main_left" style="width:100%">
                <div class="item_title">
                    <h1> {{.title}}(<span>{{.pageId}} / {{.attachNum}}</span>)</h1>
                    <div class="single-cat"><span>分类:</span> <a href="{{cateUrl .cEngName}}" rel="category tag">{{.cName}}</a> / <span>发布于</span>{{.publishTime}}</div>
                </div>
                <div class="content" id="content">
                    <div class="content_left">
                        <a href="{{.preUrl}}" title="上一页" class="pre-cat"><i class="fa fa-chevron-left"></i></a>
                        <a href="{{.nextUrl}}" title="下一页" class="next-cat"><i class="fa fa-chevron-right"></i></a>
                        <div class="image_div" id="image_div">
                            <p><a href="{{.nextUrl}}"><img src="{{attachUrl .file}}" alt="{{.title}}" title="点击图片查看下一张"></a></p>
                            <div class="nav-links page_imges">{{if .pagination}}{{.pagination}}{{end}}</div>
                        </div>
                    </div>
                </div>
                <div class="content_right_title">相关资源：
                    <span class="single-tags">
                    {{range .tags}}
                    <a href="{{tagUrl .}}">{{.}} </a>
                    {{end}}
                    </span>
                </div>
                </div>
                <ul class="xg_content">
                    {{range .relates}}
                    {{template "list" .}}
                    {{end}}
                </ul>
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
