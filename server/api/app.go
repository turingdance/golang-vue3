package api

import (
	"os"
	"path/filepath"

	"turingdance.com/turing/internal/server"
)

type App struct {
	Name             string
	DisplayName      string
	Description      string
	WorkingDirectory string
	Config           string
	Arguments        []string
	httpserver       *server.HttpServer
}
type AppOption func(*App)

func NewApp(name string, options ...AppOption) *App {
	var exepath, _ = os.Executable()
	var exeabspath, _ = filepath.Abs(exepath)
	var currentDir = filepath.Dir(exeabspath)

	var prg = &App{
		Name:             name,
		DisplayName:      name,
		Description:      name,
		WorkingDirectory: currentDir,
		Arguments:        []string{},
	}
	for _, v := range options {
		v(prg)
	}
	return prg
}
func (prg *App) WithOption(options ...AppOption) *App {
	for _, v := range options {
		v(prg)
	}
	return prg
}
func Name(name string) AppOption {
	return func(a *App) {
		a.Name = name
		a.DisplayName = name
		a.Description = name
	}
}
func DisplayName(name string) AppOption {
	return func(a *App) {
		a.DisplayName = name

	}
}
func Description(name string) AppOption {
	return func(a *App) {
		a.Description = name
	}
}
func WorkingDirectory(dir string) AppOption {
	return func(a *App) {
		a.WorkingDirectory = dir
	}
}

// 配置文件
func Config(file string) AppOption {
	return func(a *App) {
		a.Config = file
		if !filepath.IsAbs(file) {
			a.Config = filepath.Join(a.WorkingDirectory, file)
		}
	}
}
