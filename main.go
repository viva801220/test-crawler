package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"test-crawler/router"
	"time"

	"github.com/astaxie/beego/logs"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

const (
	GRACEFUL_TIMEOUT = 10
)

var helpInfo = "help\n  -h				help\n  -c conf/conf.toml	config file's path, default: conf/conf.toml\n"
var cmdConf = flag.Bool("c", false, "config file's path")
var cmdHelp = flag.Bool("h", false, "help")

func graceful(e *echo.Echo) {

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), GRACEFUL_TIMEOUT*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatalf("Error: %v", err)
	} else {
		e.Logger.Info("Server stopped")
	}
}

func enableMiddleware(e *echo.Echo) {

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Recover())
}

func createServer() {

	e := echo.New()
	enableMiddleware(e)
	e.Logger.SetLevel(log.INFO)
	router.Routers(e)
	graceful(e)
}

func main() {

	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)

	confFilePath := ""
	flag.Parse() // Scans the arg list and sets up flags
	if *cmdConf {
		if flag.NArg() == 1 {
			confFilePath = flag.Arg(0)
			fmt.Printf("run with conf:%s\n", confFilePath)
		} else {
			fmt.Printf("-c parameters error\n" + helpInfo)
			return
		}
	} else if *cmdHelp {
		fmt.Printf(helpInfo)
		return
	}

	var err error
	if err = initConfig(confFilePath); err != nil {
		log.Fatal("failed to read in config file:", err)
	}

	createServer()
}
