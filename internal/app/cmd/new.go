package cmd

import (
	"fmt"
	"path"

	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
)

// NewPost ...
func (c *CMD) NewPost(title string) error {
	// 检查发布文件夹状态
	utils.CheckDir(c.Options.SourceDir)
	utils.CheckDir(path.Join(c.Options.SourceDir, c.Options.PostDir))

	title = utils.SafeFormat(title, " ", "", "")
	var safeName = utils.SafeFormat(title, "_", "md", ".")
	var filePath = path.Join(c.Options.SourceDir, c.Options.PostDir, safeName)
	if utils.IsExist(filePath) {
		return fmt.Errorf("article %v is exist", safeName)
	}
	// 创建文章文件
	err := tools.CreateMarkdown(filePath, title, "[void]")
	if err != nil {
		return fmt.Errorf("create article %s is failed", safeName)
	}
	return nil
}

func (c *CMD) NewPage(title string) error {
	// 检查发布文件夹状态
	title = utils.SafeFormat(title, " ", "", "")
	utils.CheckDir(c.Options.SourceDir)

	utils.CheckDir(path.Join(c.Options.SourceDir, title))

	var safeName = c.IndexStr
	var filePath = path.Join(c.Options.SourceDir, title, safeName)
	if utils.IsExist(filePath) {
		return fmt.Errorf("page %v is exist", safeName)
	}
	// 创建文件
	err := tools.CreateMarkdown(filePath, title, title)
	if err != nil {
		return fmt.Errorf("create page %s is failed", safeName)
	}
	return nil
}
