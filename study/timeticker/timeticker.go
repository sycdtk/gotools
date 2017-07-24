package main

import (
	"fmt"
	"time"
)

func main() {
	call0()
	//	call1()

}

//持续周期执行
func call0() {
	t := time.NewTicker(time.Second)

	go func() {
		for t := range t.C { //for range t.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5) //阻塞，则执行次数为sleep的休眠时间/ticker的时间
	t.Stop()
	fmt.Println("Ticker stopped")
}

//仅执行几次
//func call1() {
//	//初始化断续器,间隔2s
//	var ticker *time.Ticker = time.NewTicker(100 * time.Millisecond)

//	//num为指定的执行次数
//	num := 2
//	c := make(chan int, num)
//	go func() {
//		for t := range ticker.C {
//			c <- 1
//			fmt.Println("Tick at", t)
//		}
//	}()

//	time.Sleep(time.Millisecond * 1500)
//	ticker.Stop()
//	fmt.Println("Ticker stopped")
//}
