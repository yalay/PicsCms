{{define "header"}}
<!-- 头部代码begin -->
<div class="index_header nav_headertop">
    <div class="header_inner">
        <div class="logo"><a href="{{.webUrl}}"><img src="/img/logo.png" alt="{{.webName}}"></a></div>
        <div class="header_menu">
            <ul>
                <li {{if eq $.cid 0}} class="current-menu-item"{{end}}><a href="/">网站首页</a></li>
                <li class="megamenu toke {{if eq $.cid 90}} current-menu-ancestor{{end}}"><a>精品专题</a>
                    <ul class="sub-menu">
                        <li><a>人气模特</a>
                            <ul class="sub-menu">
                            </ul>
                        </li>
                        <li><a>热门标签</a>
                            <ul class="sub-menu">
                            </ul>
                        </li>
                    </ul>
                </li>
            </ul>
        </div>
        <div class="login_text hide">
            <ul id="mobile_menu">
            <li {{if eq $.cid 0}} class="current-menu-item"{{end}}><a href="{{.webUrl}}">网站首页</a></li>
        </div>
    </div>
</div>
<!-- 头部代码end -->
{{end}}
