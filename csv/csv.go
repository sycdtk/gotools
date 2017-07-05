package csv

import (
	"encoding/csv"
	"os"
)

type CSV struct {
	name        string //文件名称
	column, row int    //行列
}

//读取csv文件
func (c *CSV) Reader(datas *[][]string) {

}

//写入csv文件
func (c *CSV) Writer(datas [][]string) {

	f, err := os.Create(c.name) //创建文件

	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f) //创建一个新的写入文件流

	w.WriteAll(datas) //写入数据

	w.Flush()
}

//csv对象构建
func NewCSV(filename string) *CSV {
	return &CSV{name: filename}
}
