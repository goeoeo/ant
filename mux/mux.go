package mux

import (
	"errors"
	"log"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var ErrAbort = errors.New("User stop run")

//控制器接口
type ControllerInterface interface {
	Init(w http.ResponseWriter, r *http.Request)
}

//控制器方法
type ControllerFunc struct {
	C ControllerInterface //控制器
	F string              //方法名
}

//自定义路由
type AntMux struct {
	staticDir string
	http.ServeMux
	debug bool
	mu    map[string]ControllerFunc
}

//初始化路由
func NewAntMux() *AntMux {
	this := new(AntMux)
	this.mu = make(map[string]ControllerFunc)
	return this
}

//Debug
func (this *AntMux) Debug(debug bool) *AntMux {
	this.debug = debug

	return this
}

//设置静态文件目录
func (this *AntMux) StaticDir(dir string) *AntMux {
	this.staticDir = dir
	return this
}

func (this *AntMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		//恢复程序的控制权
		if err := recover(); err != nil {
			if err == ErrAbort {
				return
			} else {
				panic(err)
			}
		}
	}()

	//静态文件目录
	staticDir := http.Dir(this.staticDir)
	staticHandler := http.FileServer(staticDir)
	if strings.HasPrefix(r.URL.Path, "/"+this.staticDir) {
		http.StripPrefix("/"+this.staticDir, staticHandler).ServeHTTP(w, r)
		return
	}

	cf, ok := this.mu[r.URL.Path]
	if !ok {
		http.NotFound(w, r)

		return
	}
	cf.C.Init(w, r)

	c := reflect.ValueOf(cf.C)
	if this.debug {
		log.Println("请求地址：", r.URL.Path)
	}

	c.MethodByName(cf.F).Call(nil)

	return
}

//自动注册路由
func (this *AntMux) AutoRouter(c ControllerInterface) *AntMux {

	reflectVal := reflect.ValueOf(c)
	rt := reflectVal.Type()
	ct := reflect.Indirect(reflectVal).Type()
	controllerName := strings.TrimSuffix(ct.Name(), "Controller")

	for i := 0; i < rt.NumMethod(); i++ {

		//排除一些方法
		if inArray(rt.Method(i).Name, []string{"Init"}) {
			continue
		}

		key := "/" + strings.ToLower(controllerName) + "/" + strings.ToLower(rt.Method(i).Name)

		cf := ControllerFunc{
			C: c,
			F: rt.Method(i).Name,
		}

		this.mu[key] = cf
	}

	return this
}

//启动http服务
func (this *AntMux) Serve(port string, readTimeout time.Duration, writeTimeout time.Duration) error {
	var (
		listener   net.Listener //网络监听
		err        error
		httpServer *http.Server
	)

	//启动监听
	if listener, err = net.Listen("tcp", port); err != nil {
		return err
	}

	//创建http服务
	httpServer = &http.Server{
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Handler:      this,
	}

	if err = httpServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func inArray(s string, ss []string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}

	return false
}
