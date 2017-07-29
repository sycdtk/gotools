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

	fmt.Println(DB("").TableExist("users"))

	DB("").Execute("DROP TABLE users;")

	fmt.Println(DB("").TableExist("users"))

	DB("").Execute(`CREATE TABLE
	    users
	    (
	        id INTEGER,
	        name TEXT,
	        age INTEGER
	    );`)

	fmt.Println(DB("").TableExist("users"))

	DB("").Execute("INSERT INTO users(name,age) values(?,?)", "lirui", 20)

	DB("").Execute("UPDATE users SET age = ? WHERE name = ?", 18, "lirui")

	results := DB("").Query("SELECT name,age FROM users")

	for _, row := range results {

		for _, cell := range row {

			if len(string(cell)) > 8 {
			} else {
				fmt.Println(string(cell))
			}

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

	DB("").Execute("DELETE FROM users WHERE name = ?", "lirui")
}
