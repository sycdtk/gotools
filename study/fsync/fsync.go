package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type FTest struct {
	name    string
	datas   map[string]string
	f       *os.File
	mu      *sync.RWMutex
	appData map[string]string
}

func (ft *FTest) sync() {
	ticker := time.NewTicker(time.Duration(5) * time.Second) //设置定时器
	defer ticker.Stop()

	for range ticker.C {

		func() {
			ft.mu.Lock()
			defer ft.mu.Unlock() //通过闭包释放锁
			if len(ft.appData) > 0 {
				for k, v := range ft.appData {
					ft.f.Write([]byte(k + "=" + v + "\n"))
				}
				ft.f.Sync()
				fmt.Println("data file sync!")
				ft.appData = make(map[string]string)
			}
		}()

	}

}

func (ft *FTest) set(key, value string) {
	ft.datas[key] = value
	ft.appData[key] = value
	fmt.Println("set value:", key, value)
}

func (ft *FTest) get(key string) string {
	ft.mu.RLock()
	defer ft.mu.RUnlock()
	if v, ok := ft.datas[key]; ok {
		return v
	}
	return ""

}

func NetFTest() *FTest {
	ft := &FTest{name: "test"}
	ft.datas = make(map[string]string)
	ft.appData = make(map[string]string)
	ft.mu = new(sync.RWMutex)
	var err error
	ft.f, err = os.OpenFile("file.dat", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	go ft.sync()

	return ft
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func main() {

	ft := NetFTest()

	keys := []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", "a9", "a10", "a11", "a12", "a13", "a14", "a15", "a16", "a17", "a18", "a19", "a20", "a21", "a22", "a23", "a24", "a25", "a26", "a27", "a28", "a29", "a30"}

	for i := 1; i <= 30; i++ {
		time.Sleep(time.Second)
		ft.set(keys[i], GetRandomString(6))
	}

}
