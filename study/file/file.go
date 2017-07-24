package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("file.dat", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	var buf []byte

	datas := []string{"你够浪1", "你够浪2", "你够浪3", "你够浪4", "你够浪5", "你够浪6", "你够浪7", "你够浪8", "你够浪9", "你够浪10"}

	for _, data := range datas {
		buf = []byte(data + "\n")
		f.Write(buf)
		f.Sync()
		fmt.Println(data)
		buf = buf[:0]
		time.Sleep(time.Second)
	}

	f.Close()

}
