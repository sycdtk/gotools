package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	name string `fname:"1111";lname:"2222"`
	Age  int
}

func (m MyStruct) GetName() string {
	return "getname:" + m.name
}

func (m *MyStruct) GetName1() string {
	return "getname1:" + m.name
}

func (m *MyStruct) GetName2(str string) string {
	return "getname2:" + m.name + str
}

func main() {

	m1 := new(MyStruct) //对象的引用  &MyStruct{}，需要使用Elem()
	m1.name = "m1"
	m1.Age = 19

	m2 := MyStruct{} //对象
	m2.name = "m2"
	m2.Age = 9

	m1t := reflect.TypeOf(m1)  //涉及字段定义、方法名称获取
	m1v := reflect.ValueOf(m1) //涉及对象方法调用、数值修改

	fmt.Println(m1t.Kind() == reflect.Ptr) //判断是对象还是对象的引用reflect.Struct

	fmt.Println(m1t.Elem().Field(0).Tag.Get("fname")) //读取tag值

	nf := m1t.Elem().NumField() //获取字段数量
	fmt.Println(nf)
	for i := 0; i < nf; i++ {
		fmt.Println(m1t.Elem().Field(i).Name) //输出字段名称
	}

	nm := m1t.Elem().NumMethod() //获取方法数量，关联对象
	fmt.Println(nm)
	for i := 0; i < nm; i++ {
		fmt.Println(m1t.Elem().Method(i).Name) //输出方法名称
	}

	nm = m1t.NumMethod() //获取方法数量，关联对象引用
	fmt.Println(nm)
	for i := 0; i < nm; i++ {
		fmt.Println(m1t.Method(i).Name) //输出方法名称
	}

	fmt.Println(m1v.Elem().FieldByName("name")) //获取字段值
	fmt.Println(m1v.Elem().FieldByName("Age"))

	fmt.Println(m1v.Elem().FieldByName("name").CanSet()) //能够修改值
	fmt.Println(m1v.Elem().FieldByName("Age").CanSet())

	m1v.Elem().FieldByName("Age").SetInt(99) //修改值
	fmt.Println(m1v.Elem().FieldByName("Age"))

	m1v.Elem().FieldByName("Age").Set(reflect.ValueOf(66)) //修改值
	fmt.Println(m1v.Elem().FieldByName("Age"))

	fmt.Println(m1v.MethodByName("GetName").Call(nil))                                      //调用对象关联的函数
	fmt.Println(m1v.MethodByName("GetName1").Call(nil))                                     //调用对象引用关联的函数
	fmt.Println(m1v.MethodByName("GetName2").Call([]reflect.Value{reflect.ValueOf("你好！")})) //调用带参数的函数

	r := m1v.MethodByName("GetName2").Call([]reflect.Value{reflect.ValueOf("你好！")})
	fmt.Println(r[0].String()) //处理返回值

	fmt.Println("--------------------------------------")

	m2t := reflect.TypeOf(m2)  //涉及字段定义、方法名称获取
	m2v := reflect.ValueOf(m2) //涉及对象方法调用、数值修改

	fmt.Println(m2t.Kind() == reflect.Struct) //判断是对象还是对象的引用reflect.Ptr

	fmt.Println(m2t.Field(0).Tag.Get("fname")) //读取tag值

	nf = m2t.NumField() //获取字段名称
	fmt.Println(nf)
	for i := 0; i < nf; i++ {
		fmt.Println(m2t.Field(i).Name) //输出字段名称
	}

	fmt.Println(m2v.FieldByName("name")) //获取字段值
	fmt.Println(m2v.FieldByName("Age"))

	fmt.Println(m2v.FieldByName("name").CanSet()) //能够修改值
	fmt.Println(m2v.FieldByName("Age").CanSet())

	fmt.Println(m2v.MethodByName("GetName").Call(nil)) //调用对象关联的函数
	//	fmt.Println(m2v.MethodByName("GetName1").Call(nil))                                     //调用对象引用关联的函数  对象调用不成功
	//	fmt.Println(m2v.MethodByName("GetName2").Call([]reflect.Value{reflect.ValueOf("你好！")})) //调用带参数的函数 对象调用不成功

	nm = m2t.NumMethod() //获取方法数量，关联对象，无法调用关联对象的引用
	fmt.Println(nm)
	for i := 0; i < nm; i++ {
		fmt.Println(m2t.Method(i).Name) //输出方法名称
	}

}
