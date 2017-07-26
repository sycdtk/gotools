package db

import (
	"fmt"
	//	"reflect"
	//	"os"
	"testing"
)

func TestDB(t *testing.T) {

	//	file := "foo.db"       //源文件路径
	//	err := os.Remove(file) //删除文件test.txt
	//	if err != nil {
	//		//如果删除失败则输出 file remove Error!
	//		fmt.Println("file remove Error!")
	//		//输出错误详细信息
	//		fmt.Printf("%s", err)
	//	} else {
	//		//如果删除成功则输出 file remove OK!
	//		fmt.Print("file remove OK!")
	//	}

	//	DefaultDB().Execute(`CREATE TABLE
	//    users
	//    (
	//        id INTEGER,
	//        name TEXT,
	//        age INTEGER
	//    )`)

	//	DefaultDB().Execute(`CREATE TABLE
	//    Chair
	//    (
	//        ID INTEGER,
	//        Day TEXT
	//    )`)

	DefaultDB().Execute("INSERT INTO users(name,age) values(?,?)", "lirui", 20)

	DefaultDB().Update("UPDATE users SET age = ? WHERE name = ?", 18, "lirui")

	results := DefaultDB().Query("SELECT name,age FROM users WHERE name = ?", "lirui")

	for _, row := range results {

		for _, cell := range row {

			fmt.Println(string(cell))

			//			rt := reflect.TypeOf(cell)
			//			if rt.Kind() == reflect.Ptr {
			//				rt = rt.Elem()
			//			}

			//			fmt.Println("--", rt.Name())

			//			if i, ok := cell.(*int); ok {
			//				fmt.Print(i)
			//				fmt.Print(" ")
			//			}

			//			if s, ok := cell.(string); ok {
			//				fmt.Print(s)
			//				fmt.Print(" -")
			//			}

		}
		fmt.Println("")
	}

	DefaultDB().Delete("DELETE FROM users WHERE name = ?", "lirui")
}
