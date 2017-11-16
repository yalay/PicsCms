package conf

import (
	"fmt"
	"path/filepath"
	"path"
	"strings"
)

func GenArticleUrl(articleId, pageId int) string {
	if pageId == 0 || pageId == 1 {
		return fmt.Sprintf("/article-%d.html", articleId)
	} else {
		return fmt.Sprintf("/article-%d-%d.html", articleId, pageId)
	}
}

func GenCateUrl(name, pageId int) string {
	if pageId == 0 || pageId == 1 {
		return fmt.Sprintf("/%s.html", name)
	} else {
		return fmt.Sprintf("/%s-%d.html", name, pageId)
	}
}

func GenAttachUrl(attachPath string) string {
	attachPath = strings.TrimPrefix(attachPath, RootPath())
	oriPath := filepath.ToSlash(attachPath)
	if 	strings.HasPrefix(oriPath, "http://") ||
		strings.HasPrefix(oriPath, "https://") {
		return oriPath
	}
	return path.Join("/static", oriPath)
}
