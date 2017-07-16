package csv

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/sycdtk/gotools/errtools"
)

type CSV struct {
	name        string //文件名称
	column, row int    //行列
	Datas       [][]string
}

//读取csv文件
func (c *CSV) Reader() {

	file, err := os.Open(c.name)

	errtools.CheckErr(err, "csv文件打开失败！")

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		errtools.CheckErr(err, "csv文件读取失败！")

		data := []string{}
		for _, v := range record {
			data = append(data, v)
		}

		c.Datas = append(c.Datas, data)
	}
}

//写入csv文件
func (c *CSV) Writer(datas [][]string, append bool) {

	var f *os.File

	if append {
		file, err := os.OpenFile(c.name, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //追加
		errtools.CheckErr(err, "读取csv文件失败!")
		f = file
	} else {
		file, err := os.Create(c.name) //创建文件
		errtools.CheckErr(err, "创建csv文件失败!")
		f = file
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
