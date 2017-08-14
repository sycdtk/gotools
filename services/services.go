package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func main() {
	cmdStr := flag.String("exec", "", "请输入进程包含程序名称，多个进程服务以逗号分隔！例如：services.exe -exec=\"grabdata.exe -start,grabdataview.exe\"")

	flag.Parse()

	if len(*cmdStr) == 0 {
		flag.Usage()
		os.Exit(-1)
	}

	cmds := strings.Split(*cmdStr, ",")

	var wg = &sync.WaitGroup{}

	for _, cmd := range cmds {
		wg.Add(1)
		go start(cmd)
	}

	wg.Wait()
}

//启动服务
func start(command string) {
	ticker := time.NewTicker(time.Duration(1) * time.Second) //设置定时器
	defer ticker.Stop()

	for range ticker.C {
		cmd := exec.Command(command)

		err := cmd.Start()
		if err != nil {
			fmt.Printf("%s %s 启动失败:\n%s\n", time.Now().Format("2006-01-02 15:04:05"), command, err)
			continue
		}

		fmt.Printf("%s %s 进程启动!\n", time.Now().Format("2006-01-02 15:04:05"), command)
		err = cmd.Wait()
		fmt.Printf("%s %s 进程退出:\n%s\n", time.Now().Format("2006-01-02 15:04:05"), command, err)
	}
}
