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
	flag.StringVar(&api.ProjectConfig, "c", "app-prod.yaml", "/path/to/configfile,only yaml/yml(app-prod.yaml)")
	flag.BoolVar(&version, "v", false, "version of app")
	flag.StringVar(&signalstr, "s", "start", `
	start:   run app 
	stop:    stop app
	install: install a service 
	uninstall:  uninstall the service
	`)
	flag.Parse()
	// 如果为0
	app := api.NewApp(api.ProjectName)
	if api.ProjectTitle != "" {
		app.WithOption(api.Description(api.ProjectTitle))
	}
	if api.ProjectConfig != "" {
		app.WithOption(api.Config(api.ProjectConfig))
	}
	if version {
		signalstr = "version"
	}
	app.DispatchCmd(signalstr)

}
