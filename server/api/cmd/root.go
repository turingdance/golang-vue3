package cmd

import (
	"github.com/spf13/cobra" // 安装依赖 go get -u github.com/spf13/cobra/cobra
)

// 这个是根命令定义
var rootCmd = &cobra.Command{
	Use:   "engin", // 这个就是你的自己定义的根命令
	Short: "系统核心引擎",
	Long:  `系统通过这里启动`,
	Run: func(cmd *cobra.Command, args []string) {
		Run(configfile, signalstr)
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configfile, "config", "c", "./app-prod.yml", "config file path")
	rootCmd.PersistentFlags().StringVarP(&signalstr, "signal", "s", "", "reload/start/stop")
}
