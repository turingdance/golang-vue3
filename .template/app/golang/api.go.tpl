package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/kardianos/service"
	"github.com/sirupsen/logrus"
	"github.com/turingdance/infra/ipckit"
	"turingdance.com/turing/internal/app/{{.App.Name|lcfirst}}"
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

func (p *App) Start(_ service.Service) error {
	appConf := conf.ParseConf(p.Config)

	go p.starthttp(appConf)
	go p.startipcserver()
	return nil
}

// 监听信号
func (p *App) startipcserver() {
	ipcservice := ipckit.NewUnixDomainService(p.Name)
	chdata, cherr := ipcservice.Serve()
	for {
		select {
		case data := <-chdata:
			line := string(data)
			if line == "stop" {
				p.Stop(nil)
			}
			if line == "reload" {
				p.Reload()
			}
		case err := <-cherr:
			logrus.Error(err.Error())
		}
	}

}

// 监听
func (p *App) starthttp(appConf *conf.BootConf) (svc *server.HttpServer, err error) {

	//初始化日志
	log.Initialize(appConf.Log)

	// 初始化缓存
	cache.Initialize(appConf.Redis)
	logrus.Info("cache.Initialize end")
	// 初始化应用
	sys.Initialize(appConf)
	logrus.Info("sys.Initialize end")
	httpserver := server.NewHttpServer(appConf.Http, server.UseBanner(conf.Banner()))
	p.httpserver = httpserver
	// 中间件,
	httpserver.UseMiddleWare(middleware.Cros, middleware.AccessLog)
	// api 服务
	httpserver.Handle("/sys/", sys.CreateRouter("/sys/"), auth.NewAuthorize(appConf.Auth.WhiteList).Handle)

	// 这里是生成的模板
	{{.App.Name|lcfirst}}.Initialize(appConf)
	httpserver.Handle("/{{.App.Name|lcfirst}}/", {{.App.Name|lcfirst}}.CreateRouter("/{{.App.Name|lcfirst}}/"), auth.NewAuthorize(appConf.Auth.WhiteList).Handle)

	logrus.Info("httpserver.Handle end")
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
			logrus.Info("Storage end")
		}

	}

	// 静态资源服务,该目录是console 项目下 npm run build:prod 指令编译后生成
	httpserver.Handle("/", http.FileServer(http.FS(site.Assets)))

	//ecls
	//httpserver.Handle("/ecls/", cls.CreateRouter("/ecls/"), auth.NewAuthorize(appConf.Auth.WhiteList).Handle)
	err = httpserver.Start()
	logrus.Info("httpserver.Start end")
	return httpserver, err
}

func (p *App) Stop(_ service.Service) error {
	if p.httpserver != nil {
		p.httpserver.Shutdown()
		p.httpserver = nil
	}
	os.Exit(0)
	return nil
}

func (p *App) Reload() error {
	conf.Reload()
	return nil
}

func (p *App) Run() (err error) {
	svcConfig := &service.Config{
		Name:             p.Name,        //服务显示名称
		DisplayName:      p.DisplayName, //服务名称
		Description:      p.Description, //服务描述
		WorkingDirectory: p.WorkingDirectory,
		Arguments:        []string{"-c", p.Config},
	}
	s, _ := service.New(p, svcConfig)
	return s.Run()
}

func (p *App) DispatchCmd(cmd string) (err error) {
	switch cmd {
	case "version":
		err = p.Version()
	case "install":
		err = p.Install()
	case "uninstall":
		err = p.Uninstall()
	case "stop":
		err = p.SendStop()
	case "reload":
		err = p.SendReload()
	default:
		err = p.Run()
	}
	if err != nil {
		fmt.Printf("%s error : %s", cmd, err.Error())
	} else {
		fmt.Printf("%s ok", cmd)
	}
	return err
}

func (p *App) Install() (err error) {
	svcConfig := &service.Config{
		Name:             p.Name,        //服务显示名称
		DisplayName:      p.DisplayName, //服务名称
		Description:      p.Description, //服务描述
		WorkingDirectory: p.WorkingDirectory,
		Arguments:        []string{},
	}
	s, _ := service.New(p, svcConfig)
	return s.Install()
}
func (p *App) Uninstall() (err error) {
	svcConfig := &service.Config{
		Name:             p.Name,        //服务显示名称
		DisplayName:      p.DisplayName, //服务名称
		Description:      p.Description, //服务描述
		WorkingDirectory: p.WorkingDirectory,
		Arguments:        []string{"-c", p.Config},
	}
	s, _ := service.New(p, svcConfig)
	return s.Uninstall()
}

func (p *App) Version() (err error) {
	fmt.Println(conf.Banner())
	return nil
}

// 停止程序
func (p *App) SendReload() error {
	ipcservice := ipckit.NewUnixDomainService(p.Name)
	_, err := ipcservice.WriteLine("reload")
	return err
}

// 停止程序
func (p *App) SendStop() error {
	ipcservice := ipckit.NewUnixDomainService(p.Name)
	_, err := ipcservice.WriteLine("stop")
	return err
}
