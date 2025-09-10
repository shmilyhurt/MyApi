package main

import (
	"MyApi/internal/router"
	"MyApi/pkg/database"
	"context"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)

	}
}

func main() {
	// 1. 初始化配置
	initConfig()
	// 2. 初始化数据库
	database.InitMysql()
	// 3. 初始化路由
	r := router.SetupRouter()
	// 4. 启动 HTTP 服务 (不直接用 r.Run)
	port := viper.GetString("server.port")
	srv := http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	// 5. 启动服务 (异步)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 6. 监听退出信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	// 7. 优雅关闭 (设置超时时间，比如 5 秒)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown1:", err)
	}
}
