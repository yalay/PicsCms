{{define "list"}}
<li class="i_list list_n2">
    <a target="_blank" href="" title="{{.Title}}">
        <img class="waitpic" src="/img/loading.gif" data-original="{{.Cover}}?s=270x370" width="270" height="370" alt="{{.Title}}" style="display: inline;">
    </a>
    <div class="case_info">
        <div class="meta-title"> {{.Title}} </div>
        <div class="meta-post"><i class="fa fa-clock-o"></i> {{.PublishTime}} </div>
    </div>
</li>
{{end}}
