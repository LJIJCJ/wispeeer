package cmd

import (
	"fmt"
	"path"
	"time"

	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	logeer "github.com/ka1i/wispeeer/pkg/log"
)

// NewPost ...
func (c *CMD) NewPost(title string) error {
	logeer.Println("new").Infof("Location: %s", utils.GetWorkspace())

	// 检查发布文件夹状态
	utils.CheckDir(path.Join(utils.GetWorkspace(), c.Options.SourceDir))
	utils.CheckDir(path.Join(utils.GetWorkspace(), c.Options.SourceDir, c.Options.PostDir))

	title = utils.SafeFormat(title, " ", "", "")
	var safeName = utils.SafeFormat(title, "_", "md", ".")
	var filePath = path.Join(utils.GetWorkspace(), c.Options.SourceDir, c.Options.PostDir, safeName)
	if utils.IsExist(filePath) {
		return fmt.Errorf("article %v is exist", safeName)
	}
	// 创建文章文件
	err := tools.CreateMarkdown(filePath, title, "void")
	if err != nil {
		return fmt.Errorf("create article %s is failed", safeName)
	}
	showInfo(title, safeName)
	return nil
}

func (c *CMD) NewPage(title string) error {
	logeer.Println("new").Infof("Location: %s", utils.GetWorkspace())

	// 检查发布文件夹状态
	title = utils.SafeFormat(title, " ", "", "")
	utils.CheckDir(path.Join(utils.GetWorkspace(), c.Options.SourceDir))

	utils.CheckDir(path.Join(utils.GetWorkspace(), c.Options.SourceDir, title))

	var safeName = "index.md"
	var filePath = path.Join(utils.GetWorkspace(), c.Options.SourceDir, title, safeName)
	if utils.IsExist(filePath) {
		return fmt.Errorf("page %v is exist", safeName)
	}
	// 创建文件
	err := tools.CreateMarkdown(filePath, title, title)
	if err != nil {
		return fmt.Errorf("create page %s is failed", safeName)
	}
	showInfo(title, safeName)
	return nil
}

func showInfo(title string, safeName string) {
	fmt.Printf("title  : %s\n", title)
	fmt.Printf("posted : %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("Created: %s\n", safeName)
}
