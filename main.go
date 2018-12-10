package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"github.com/sevlyar/go-daemon"
)

type MyMux struct {
	routers map[string]func(res http.ResponseWriter, req *http.Request)
}

func (p *MyMux) ServeHTTP(res http.ResponseWriter, req *http.Request){
	//遍历routers，寻找匹配到的path
	for path, f := range p.routers{
		if ok, _ := regexp.MatchString("^" + path + "$", req.URL.Path); ok {
			f(res, req)
			return
		}
	}
	fmt.Fprintf(res, "Error: Don`t match URL '%s'", req.URL.Path)
}

func Html ( res http.ResponseWriter, req *http.Request) {
	//用于保存数据的map
	//data := make(map[string]string)
	//解析指定模板文件html
	t, err := template.ParseFiles("view/index.html")
	if err != nil {
		fmt.Printf("Html Error: %+v", nil)
	}
	//data["Name"] = "Gao"

	//输出到浏览器
	t.Execute(res, nil)
}

func _daemon() {
	cntxt := &daemon.Context{
		PidFileName: "pid",
		PidFilePerm: 0644,
		LogFileName: "log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")
}

func Static( res http.ResponseWriter, req *http.Request){
	fmt.Println("Deal Static: ", req.URL.Path)
	res.Header().Set("Content-Type", "text/css")
	http.ServeFile(res, req, "." + req.URL.Path)
}

func main () {
	_daemon()
	mux := &MyMux{}
	mux.routers = make(map[string]func(res http.ResponseWriter, req *http.Request))
	mux.routers["/"] = Html
	mux.routers["/static/.+"] = Static
	err := http.ListenAndServe(":80", mux)

	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
}
//resume/wmutong_go/WMuTong-vue-go-/main
//
//[program:golang-http-server]
//command=/root/resume/wmutong_go/WMuTong-vue-go-/main
//autostart=true
//autorestart=true
//startsecs=10
//stdout_logfile=/var/log/simple_http_server.log
//stdout_logfile_maxbytes=1MB
//stdout_logfile_backups=10
//stdout_capture_maxbytes=1MB
//stderr_logfile=/var/log/simple_http_server.log
//stderr_logfile_maxbytes=1MB
//stderr_logfile_backups=10
//stderr_capture_maxbytes=1MB