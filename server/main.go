package main

import (
	"flag"

	"turingdance.com/turing/api"
)

//go:generate  codectl router  -s ./internal/app/$(app)/rest
var (
	version   bool
	signalstr string
)

func main() {
	flag.StringVar(&api.AppConfig, "c", "app-prod.yaml", "/path/to/configfile,only yaml/yml(app-prod.yaml)")
	flag.BoolVar(&version, "v", false, "version of app")
	flag.StringVar(&signalstr, "s", "start", `
	start:   run app 
	stop:    stop app
	install: install a service 
	uninstall:  uninstall the service
	`)
	flag.Parse()
	// 如果为0
	app := api.NewApp(api.AppName)
	if api.AppMemo != "" {
		app.WithOption(api.Description(api.AppMemo))
	}
	if api.AppConfig != "" {
		app.WithOption(api.Config(api.AppConfig))
	}
	if version {
		signalstr = "version"
	}
	app.DispatchCmd(signalstr)

}
