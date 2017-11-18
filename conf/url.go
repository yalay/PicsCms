package conf

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func GenArticleUrl(articleId int, args ...int) string {
	if articleId < 0 {
		return "/"
	}

	if len(args) == 0 || args[0] == 1 {
		return fmt.Sprintf("/article-%d.html", articleId)
	} else {
		return fmt.Sprintf("/article-%d-%d.html", articleId, args[0])
	}
}

func GenCateUrl(name string, args ...int) string {
	if len(args) == 0 || args[0] == 1 {
		return fmt.Sprintf("/%s.html", name)
	} else {
		return fmt.Sprintf("/%s-%d.html", name, args[0])
	}
}

func GenAttachUrl(attachPath string) string {
	attachPath = strings.TrimPrefix(attachPath, RootPath())
	oriPath := filepath.ToSlash(attachPath)
	if strings.HasPrefix(oriPath, "http://") ||
		strings.HasPrefix(oriPath, "https://") {
		return oriPath
	}
	return "//" + path.Join("cdn.tengmm.com", oriPath)
}

func GenTagUrl(tag string, args ...int) string {
	if len(args) == 0 || args[0] == 1 {
		return fmt.Sprintf("/tags-%s.html", tag)
	} else {
		return fmt.Sprintf("/tags-%s-%d.html", tag, args[0])
	}
}
