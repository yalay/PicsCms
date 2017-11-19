{{define "header"}}
<!-- 头部代码begin -->
<div class="index_header nav_headertop">
    <header id="luxbar" class="luxbar-default">
        <input type="checkbox" class="luxbar-checkbox" id="luxbar-checkbox"/>
        <div class="luxbar-menu luxbar-menu-right luxbar-menu-light">
            <ul class="luxbar-navigation">
                <li class="luxbar-header">
                    <a href="/" class="luxbar-brand"><img src="/img/logo.png"></a>
                    <label class="luxbar-hamburger luxbar-hamburger-spin"
                           id="luxbar-hamburger" for="luxbar-checkbox"> <span></span> </label>
                </li>
                <li {{if eq $.cid 0}}class="luxbar-item active"{{else}}class="luxbar-item"{{end}}><a href="/">网站首页</a></li>
                {{range .totalCates}}
                <li {{if eq $.cid .Id}}class="luxbar-item active"{{else}}class="luxbar-item"{{end}}><a href="{{cateUrl .EngName}}">{{.Name}}</a></li>
                {{end}}
            </ul>
        </div>
    </header>
    <!--<div class="header_inner">-->
        <!--<div class="logo"><a href="{{.webUrl}}"><img src="/img/logo.png" alt="{{.webName}}"></a></div>-->
        <!--<div class="header_menu">-->
            <!--<ul>-->
                <!--<li {{if eq $.cid 0}} class="current-menu-item"{{end}}><a href="/">网站首页</a></li>-->
                <!--{{range .totalCates}}-->
                <!--<li {{if eq $.cid .Id}} class="current-menu-item"{{end}}><a href="{{cateUrl .EngName}}">{{.Name}}</a></li>-->
                <!--{{end}}-->
            <!--</ul>-->
        <!--</div>-->
    <!--</div>-->
</div>
<!-- 头部代码end -->
{{end}}
