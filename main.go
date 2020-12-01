package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init(){
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}

//	@title				博客系统
//	@version 			1.0
//	@description		GO 语言学习之旅:博客系统
//	@termOfService		https://github.com/reation/blog-service
func main(){
	global.Logger.Infof("%s: go-log/%s", "reation", "book_blog")
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr					: ":" + global.ServerSetting.HttpPort,
		Handler					: router,
		ReadHeaderTimeout		: global.ServerSetting.ReadTimeOut,
		WriteTimeout			: global.ServerSetting.WriteTimeOut,
		MaxHeaderBytes			: 1 << 20,
	}

	s.ListenAndServe()
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupSetting() error{
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeOut 	*= time.Second
	global.ServerSetting.WriteTimeOut 	*= time.Second

	return nil
}

func setupLogger() error{
	t 		:= time.Now()
	date 	:= fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())

	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + "-" + date + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename	: fileName,
		MaxSize		: 600,
		MaxAge		: 10,
		LocalTime	: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}