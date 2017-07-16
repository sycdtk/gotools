package db

import (
	//	"fmt"
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

	DefaultDB().Create(`CREATE TABLE
    users
    (
        id INTEGER,
        name TEXT,
        age INTEGER
    )`)

	DefaultDB().Create("INSERT INTO users(name,age) values(?,?)", "lirui", 20)

	DefaultDB().Update("UPDATE users SET age = ? WHERE name = ?", 18, "lirui")

	DefaultDB().Delete("DELETE FROM users WHERE name = ?", "lirui")
}
