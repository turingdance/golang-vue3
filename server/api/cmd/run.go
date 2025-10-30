package cmd

import (
	"os"

	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
	"github.com/turingdance/infra/logger"
)

// 命令执行方法
func Execute() {
	if len(os.Args) == 1 {
		Run(configfile, signalstr)
	} else {
		cobra.OnInitialize(func() {})
		// 执行命令 如果异常会返回错误信息
		if err := rootCmd.Execute(); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	}

}
