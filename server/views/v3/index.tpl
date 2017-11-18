{{define "home"}}
<!DOCTYPE html>
<html>
<head>
    {{template "top"}}
    <title>{{.webName}} - {{.webKeywords}}</title>
    <meta name="keywords" content="{{.webKeywords}}">
    <meta name="description" content="{{.webDesc}}">
    <link rel="stylesheet" href="/css/jquery.bxslider.min.css" type="text/css" media="all">
    <link rel="stylesheet" href="/css/backtotop.css" type="text/css" media="all">
</head>
<body class="home blog body_top">
    {{template "header" .}}
    <!--效果html开始-->
    <div class="site-wrap hide">
        <ul class="bxslider">
            {{range .sliderArticles}}
            <li><a target="_blank" href="{{articleUrl .Id}}"><img src="{{.Cover}}" title="{{.Title}}"></a></li>
            {{end}}
        </ul>
    </div>
    {{range .totalCates}}
    <div class="home-filter">
        <div class="h-screen-wrap">
            <ul class="h-screen"><li class="current-menu-item"><a href="{{cateUrl .EngName}}"> {{.Name}} </a></li></ul>
        </div>
        <ul class="h-soup cl">
            <li class="open"><i class="fa fa-coffee"></i><a href="{{cateUrl .EngName}}" title="{{.Name}}">  查看更多 </a></li>
        </ul>
    </div>
    <div class="update_area">
        <div class="update_area_content">
            <ul class="update_area_lists cl">
                {{range (index $.cateArticles .Id)}}
                {{template "list" .}}
                {{end}}
            </ul>
        </div>
    </div>
    {{end}}
    {{template "footer" .}}
    <script type="text/javascript" src="/js/jquery.bxslider.min.js"></script>
    <script type="text/javascript">
        $(document).ready(function(){
            $('.site-wrap').removeClass('hide');
            $('.bxslider').bxSlider({
                moveSlides: 1,
                slideMargin: 5,
                infiniteLoop: true,
                slideWidth: 590,
                minSlides: 1,
                maxSlides: 6,
                pager: false,
                controls: true,
                auto: true,
            });
        });
    </script>
</body>
</html>
{{end}}
