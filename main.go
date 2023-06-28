package main

import (
	"runtime"

	"github.com/alphayan/ipproxy-pool/api"
	"github.com/alphayan/ipproxy-pool/cmd"
	"github.com/alphayan/ipproxy-pool/middleware/config"
	"github.com/alphayan/ipproxy-pool/middleware/database"
	"github.com/alphayan/ipproxy-pool/middleware/logutil"
	"github.com/alphayan/ipproxy-pool/run"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 检查或设置命令行参数
	cmd.Execute()

	setting := config.ServerSetting

	// 将日志写入文件或打印到控制台
	logutil.InitLog(&setting.Log)
	// 初始化数据库连接
	database.InitDB(&setting.Database)

	// Start HTTP
	go func() {
		api.Run(&setting.System)
	}()

	// Start Task
	run.Task()
}
