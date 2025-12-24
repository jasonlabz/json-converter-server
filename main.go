package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/jasonlabz/json-converter-server/bootstrap"
)

// @title		    TODO: ***********服务
// @version		    1.0
// @description	    TODO: 旨在***********
// @host			TODO: localhost:port
// @contact.name	TODO: your name
// @contact.url	    TODO: http://www.*****.io/support
// @contact.email	TODO: mail_name@qq.com
// @BasePath		TODO: /base_path
func main() {
	// context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// bootstrap init
	bootstrap.MustInit(ctx)

	serverConfig := bootstrap.GetConfig()

	startFileServer(serverConfig)

	log.Println("Server exiting")
}

// startFileServer 文件服务
func startFileServer(c *bootstrap.Config) {
	config := c.GetServerConfig().Static
	// 创建 HTTP 服务器
	if config.Path == "" {
		return
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(config.Path)))
	if config.Username != "" && config.Password != "" {
		// 使用基本认证保护文件下载路由
		authMux := basicAuth(mux, config.Username, config.Password)
		// 启动 HTTP 服务器
		// log.Printf("Starting file server at :%d", config.GetConfig().Application.Port+1)
		err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), authMux)
		if err != nil {
			log.Fatalf("file server listen: %s\n", err)
		}
		return
	}
	// 启动 HTTP 服务器
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux)
	if err != nil {
		log.Fatalf("file server listen: %s\n", err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	log.Println("shutdown file server ...")
	return
}

// basicAuth 认证检查
func basicAuth(handler http.Handler, username, password string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != username || pass != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
