package main

import (
	"fmt"
	"math/rand"
	"time"
)

//生成随机字符串
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

//windows 下time.Now().UnixNano()无法获取精确值
func GetRandomInt(i int) int {
	return rand.Intn(i)
}

//linux
func GetRandomInt1(i int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(i)
}

func main() {

	fmt.Println(string(rand.Intn(99))

}
