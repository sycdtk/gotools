package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Data struct {
	mu   *sync.RWMutex
	data string
}

func (d *Data) change(v string) {
	d.mu.Lock()
	defer d.mu.Unlock()

	time.Sleep(time.Duration(5) * time.Second)
	fmt.Println("data from", d.data, " to ", v)
	d.data = v
}

func (d *Data) read(i int) {
	d.mu.RLock()
	defer d.mu.RUnlock()

	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println(i, " 读取：", d.data)

}

func main() {
	a := &Data{data: "a", mu: new(sync.RWMutex)}

	for i := 0; i < 10; i++ {
		go func(d int) {
			a.change(strconv.Itoa(d))
		}(i)
	}

	for i := 0; i < 20; i++ {
		go func(dt int) {
			a.read(dt)
		}(i)
	}

	time.Sleep(time.Duration(10) * time.Minute)

}
