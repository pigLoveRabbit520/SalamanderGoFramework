package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"github.com/spf13/viper"
	"github.com/goware/cors"
	"fmt"
	"os"
	"context"
	"os/signal"
	"time"
	"github.com/go-chi/chi/middleware"
	"github.com/salamander/SalamanderGo/middlewares"
	"github.com/salamander/SalamanderGo/components/logger"
)

var (
	log = logger.DefaultLogger()
)

func initConfig() {
	// info.BuildInfo()
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("未找到配置文件或解析失败")
		fmt.Println(err)
		os.Exit(1)
	}
}

func initServer() {
	//server.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
	//Getter: middleware.MethodFromForm("_method"),
	//}))
	r := chi.NewRouter()
	r.Use(cors.Default().Handler)
	r.Use(middleware.RealIP)
	r.Use(middlewares.Logger())
	r.Use(middlewares.Recoverer)

	checkUserToken := viper.GetBool("check_token")
	apiRouter := InitAPIRouter(checkUserToken)
	r.Mount("/", apiRouter)


	port := viper.GetString("port")
	if port == "" {
		fmt.Println("请设定端口")
		os.Exit(1)
	}

	url := "localhost:" + port
	server := &http.Server{Addr: url, Handler: r}

	// graceful
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		// sig is a ^C, handle it
		log.Warn("app is interrupt")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		if err := server.Shutdown(ctx); err != nil {
			log.WithError(err).Error("shutdown failed")
		}
	}()

	fmt.Println("start at http://" + url + "\n")
	err := server.ListenAndServe()
	log.WithError(err).Error("start server failed")
}


func main() {
	initConfig()
	initServer()
}
