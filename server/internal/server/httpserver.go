package server

import (
	"context"
	"fmt"
	"net/http"
)

type HttpConf struct {
	Port int
	Host string
}
type ServerConf struct {
	Name    string
	Datadir string
	HttpConf
}
type MiddleWare func(http.Handler) http.Handler
type HttpServer struct {
	port       int
	host       string
	mux        *http.ServeMux
	server     *http.Server
	banner     string
	handle     http.Handler
	middleware []MiddleWare
}
type HttpServerOption func(*HttpServer)

func UseHandle(h http.Handler) HttpServerOption {
	return func(hs *HttpServer) {
		hs.handle = h
	}
}
func UseBanner(str string) HttpServerOption {
	return func(hs *HttpServer) {
		hs.banner = str
	}
}
func UseMiddleWare(mw ...MiddleWare) HttpServerOption {
	return func(hs *HttpServer) {
		hs.UseMiddleWare(mw...)
	}
}
func NewHttpServer(conf HttpConf, options ...HttpServerOption) *HttpServer {
	s := &HttpServer{
		port: conf.Port,
		host: conf.Host,
		mux:  &http.ServeMux{},
	}
	for _, opt := range options {
		opt(s)
	}
	return s
}

// 全局拦截器
func (s *HttpServer) UseMiddleWare(mw ...MiddleWare) {
	s.middleware = append(s.middleware, mw...)
}

// 局部拦截器
// h1(h2(h3(h4(h5(h6(h7(h8)))))))
// 如上所属拦截器组
// 拦截器运行顺序h1->h2->h3->h4->h5->h6->h7->h8
func (s *HttpServer) Handle(patern string, handle http.Handler, mw ...MiddleWare) {
	//
	handle0 := handle
	// 后处理的放后面
	for i := len(mw) - 1; i >= 0; i-- {
		nextfun := mw[i]
		handle0 = nextfun(handle0)
	}
	// 先处理的放前面
	for _, f := range s.middleware {
		handle0 = f(handle0)
	}
	s.mux.Handle(patern, handle0)
}
func (s *HttpServer) Start() error {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	fmt.Println(s.banner)
	fmt.Println("run @" + addr)
	s.server = &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}
	err := s.server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
	//<-sigs
}
func (s *HttpServer) Shutdown() {
	s.server.Shutdown(context.Background())
}
