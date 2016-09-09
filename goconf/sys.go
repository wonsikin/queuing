package goconf

import (
	"fmt"
	"os"
	"strings"
)

// WorkDir 当前应用工作目录
var WorkDir string

const (
	pkg = "github.com/wonsikin/seq"
)

// workDir 获取程序启动目录
func getWorkDir() {
	var err error
	WorkDir, err = os.Getwd()
	if err != nil {
		fmt.Printf("can not get work directory: %s\n", err)
		os.Exit(2)
	}

	if pos := strings.Index(WorkDir, pkg); pos >= 0 {
		WorkDir = WorkDir[:(pos + len(pkg))]
	}

	fmt.Printf("work directory:\t %s\n", WorkDir)
}
