package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/ohko/dd"
)

var (
	daemon = flag.Bool("d", false, "daemon")
	close  = flag.Bool("close", false, "close server")
)

func main() {
	flag.Parse()

	pidFile := "/tmp/dd_example.pid"

	if *daemon {
		// 启动守护进程
		dd.Daemon(pidFile, true, true, time.Second*3)
	} else if *close {
		// 关闭程序和守护进程
		if err := dd.Close(pidFile); err != nil {
			log.Println(err)
		}
		return
	}

	// 子进程
	// ...
	fmt.Println("Hello world!")

	// 等待结束信号
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	<-c
}
