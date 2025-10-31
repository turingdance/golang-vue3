package cmd

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/turingdance/infra/filekit"
	"github.com/turingdance/infra/logger"
	"turingdance.com/turing/internal/app/sys"
	"turingdance.com/turing/internal/conf"
	"turingdance.com/turing/internal/pkg/cache"
	"turingdance.com/turing/internal/pkg/log"
	"turingdance.com/turing/internal/pkg/storage"
	"turingdance.com/turing/internal/server"
	"turingdance.com/turing/internal/server/auth"
	"turingdance.com/turing/internal/server/middleware"
	"turingdance.com/turing/site"
)

// 程序 运行
func Start(appConf *conf.BootConf) (svr *server.HttpServer, err error) {

	//初始化pid
	if !filekit.Exists(appConf.Engin.Pidfile) {
		dir := filepath.Dir(appConf.Engin.Pidfile)
		err = os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			return
		}
	}
	// 获得当前程序的pid
	pid := os.Getpid()
	err = os.WriteFile(appConf.Engin.Pidfile, []byte(fmt.Sprintf("%d", pid)), 0755)
	if err != nil {
		return
	}
	//初始化日志
	log.Initialize(appConf.Log)

	// 初始化缓存
	cache.Initialize(appConf.Redis)
	// 初始化应用
	sys.Initialize(appConf)

	httpserver := server.NewHttpServer(appConf.Http, server.UseBanner(conf.Banner))
	// 中间件,
	httpserver.UseMiddleWare(middleware.Cros, middleware.AccessLog)
	// rpc 服务,如果需要,主要考虑分布式
	// server.Handle("/rpc.api/", handle.NewRpcxHandler("/rpc.api/", appConf.Discover))
	// api 服务
	httpserver.Handle("/sys/", sys.CreateRouter("/sys/"), auth.NewAuthorize(appConf.Auth.WhiteList).Handle)
	// httpserver.Handle
	// 下面是资源文件服务
	for _, config := range appConf.Storage {
		if config.Driver == storage.DriverLocal {
			// 不需要鉴权
			if !config.AuthRequired {
				httpserver.Handle("/"+config.Bucket+"/", http.StripPrefix("/"+config.Bucket+"/", http.FileServer(http.Dir(config.Datapath))))
			} else {
				// 签名鉴权
				httpserver.Handle("/"+config.Bucket+"/", http.StripPrefix("/"+config.Bucket+"/", http.FileServer(http.Dir(config.Datapath))), middleware.NewSigner(appConf.Enpryt).Handle)
			}
		}

	}
	// 静态资源服务,该目录是console 项目下 npm run build:prod 指令编译后生成
	httpserver.Handle("/", http.FileServer(http.FS(site.Assets)))

	//ecls
	//httpserver.Handle("/ecls/", cls.CreateRouter("/ecls/"), auth.NewAuthorize(appConf.Auth.WhiteList).Handle)
	err = httpserver.Start()
	return httpserver, err
}

func SendSignal(pid int, s syscall.Signal) error {
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	// 发送
	return process.Signal(s)
}

// 停止程序
func SendReload(appConf *conf.BootConf) error {
	data, err := os.ReadFile(appConf.Engin.Pidfile)
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	return SendSignal(pid, syscall.SIGHUP)
}

// 停止程序
func SendStop(appConf *conf.BootConf) error {
	data, err := os.ReadFile(appConf.Engin.Pidfile)
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	// 发送
	err = process.Signal(syscall.SIGTERM)
	return err
}

func Run(file, signalstr string) {
	appConf := conf.ParseConf(file)
	switch signalstr {
	case "stop":
		err := SendStop(appConf)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	case "reload":
		err := SendReload(appConf)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	default:
		svr, err := Start(appConf)
		if err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
		sig := make(chan os.Signal, 2)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		for s := range sig {
			switch s {
			case syscall.SIGTERM, syscall.SIGKILL:
				svr.Shutdown()
				return
			case syscall.SIGHUP:
				// 加载配置文件
				appConf := conf.ParseConf(file)
				Start(appConf)
			default:
				svr.Shutdown()
				return
			}
		}
	}
}
