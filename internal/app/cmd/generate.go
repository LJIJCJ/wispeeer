package cmd

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"

	"github.com/ka1i/wispeeer/internal/pkg/tools"
	"github.com/ka1i/wispeeer/internal/pkg/utils"
	logeer "github.com/ka1i/wispeeer/pkg/log"
)

func (c *CMD) Generate() error {
	var err error
	logeer.Task("generate").Infof("Location : %v", utils.GetWorkspace())

	staticAssets := path.Join(utils.GetWorkspace(), c.ThemeStr, c.Options.Theme, "static")
	if utils.IsExist(staticAssets) {
		logeer.Task("generate").Info("copy static assets")
		err = tools.FileCopy(staticAssets, c.Options.PublicDir)
		if err != nil {
			return err
		}
	}

	logeer.Task("generate").Infof("public in: %v", c.Options.PublicDir)

	// kids! run
	err = c.prepare(c.Options.SourceDir)
	if err != nil {
		return err
	}
	return nil
}

func (c *CMD) prepare(startDIR string) error {
	files, err := ioutil.ReadDir(startDIR)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.Name()[0] == 46 {
			continue
		}
		filefullName := path.Join(startDIR, f.Name())
		pathLevelSlice := strings.Split(filepath.ToSlash((filepath.Dir(filefullName))), "/")
		pathLevel := len(pathLevelSlice)
		if utils.IsFile(filefullName) {
			if pathLevel == 1 {
				fmt.Printf("[COPY] ")
			}
			suffix := path.Ext(f.Name())
			title := strings.TrimSuffix(f.Name(), suffix)

			if pathLevel == 2 && suffix == ".md" {
				assetRoot := path.Join(startDIR, title)
				if pathLevelSlice[1] == c.Options.PostDir {
					fmt.Printf("[POST] ")
					fmt.Println(pathLevel, "FILE", filefullName)
					if utils.IsDir(assetRoot) {
						fmt.Printf("[COPY] ")
						fmt.Println(pathLevel, "DIR", assetRoot)
					}
				} else {
					fmt.Printf("[PAGE] ")
					fmt.Println(pathLevel, "FILE", filefullName)
					if utils.IsDir(assetRoot) {
						fmt.Printf("[COPY] ")
						fmt.Println(pathLevel, "DIR", assetRoot)
					}
				}
			}
		} else {
			if pathLevel == 2 {
				continue
			}
			err := c.prepare(filefullName)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
