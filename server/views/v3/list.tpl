{{define "list"}}
<li class="i_list list_n2">
    <a target="_blank" href="{{articleUrl .Id}}" title="{{.Title}}">
        <img class="waitpic" src="/img/loading.gif" data-original="{{attachUrl .Cover}}" width="270" height="370" alt="{{.Title}}" style="display: inline;">
    </a>
    <div class="case_info">
        <div class="meta-title"> {{.Title}} </div>
        <div class="meta-post"><i class="fa fa-clock-o"></i> {{.PublishTime}} </div>
    </div>
</li>
{{end}}
