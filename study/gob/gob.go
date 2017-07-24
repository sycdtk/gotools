package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	//	"os"
)

type MyStruct struct {
	Id   int
	Name string
	age  int
}

func (m *MyStruct) GetName() string {
	return m.Name
}

func (m MyStruct) GetName1() string {
	return m.Name
}

func (m *MyStruct) getName2() string {
	return m.Name
}

func (m MyStruct) getName3() string {
	return m.Name
}

//仅能序列化导出字段

var buf bytes.Buffer
var buf1 bytes.Buffer

func main() {

	enc := gob.NewEncoder(&buf)
	err := enc.Encode("asdf")
	if err != nil {
		panic("err")
	}

	fmt.Println(buf.String())
	//buf.Reset()

	var str string
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(&str)
	if err != nil {
		panic("err")
	}

	fmt.Println(str)
	//	//序列化
	//	file, err := os.Create("gob.dat")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	user := MyStruct{Id: 1, Name: "Mike", age: 9}
	//	user1 := MyStruct{Id: 2, Name: "Jack", age: 16}

	//	u := []MyStruct{user, user1}
	//	enc := gob.NewEncoder(file) //绑定writer
	//	err2 := enc.Encode(u)       //输出
	//	fmt.Println(err2)

	//	//反序列化
	//	var du []MyStruct
	//	dfile1, err := os.Open("gob.dat")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	dec := gob.NewDecoder(dfile1)
	//	err2 = dec.Decode(&du)

	//	if err2 != nil {
	//		fmt.Println(err2)
	//		return
	//	}

	//	for _, user := range du {
	//		fmt.Println(user.Id)
	//		fmt.Println(user.GetName())
	//		fmt.Println(user.GetName1())
	//		fmt.Println(user.getName2())
	//		fmt.Println(user.getName3())
	//		fmt.Println(user.age)
	//	}
}
