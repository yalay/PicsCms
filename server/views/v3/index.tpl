<!DOCTYPE html>
<html>
<head>
    {{template "top"}}
    <title>{{.webName}} - {{.webKeywords}}</title>
    <meta name="keywords" content="{{.webKeywords}}">
    <meta name="description" content="{{.webDesc}}">
    <link rel="stylesheet" href="//cdn.bootcss.com/bxslider/4.2.12/jquery.bxslider.min.css" type="text/css" media="all">
    <link rel="stylesheet" href="/css/backtotop.css" type="text/css" media="all">
</head>
<body class="home blog body_top" youdao="bind">
    {{template "header" .}}
    <!--效果html开始-->
    <div class="site-wrap hide">
        <ul class="bxslider">
        </ul>
    </div>
    {{template "footer" .}}
    <script type="text/javascript" src="//cdn.bootcss.com/bxslider/4.2.12/jquery.bxslider.min.js"></script>
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
