package lib

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Configuration struct {
	FilePath string
	Debug    bool
	Driver   string
}

func (c Configuration) DebugMessage() string {
	if c.Debug {
		return "in debug mode"
	}
	return ""
}

func (c Configuration) FilePathResult() string {
	i, j := strings.LastIndex(c.FilePath, "/"), strings.LastIndex(c.FilePath, path.Ext(c.FilePath))
	if j <= 0 {
		return ""
	}
	if i < 0 {
		return fmt.Sprintf("%s_result.json", c.FilePath[0:j])
	}
	return fmt.Sprintf("%s_result.json", c.FilePath[i:j])
}

func (c Configuration) HasFilePathResult() bool {
	return len(c.FilePathResult()) > 0
}

func (c Configuration) CheckFile() {
	if _, err := os.Stat(c.FilePath); os.IsNotExist(err) {
		panic("filepath not exist")
	}
}
