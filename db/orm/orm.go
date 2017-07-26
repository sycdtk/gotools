package orm

import (
	"bytes"
	"fmt"
	"reflect"
	"sync"

	"github.com/sycdtk/gotools/db"
	"github.com/sycdtk/gotools/logger"
)

var struMap map[string]*StruInfo //结构对象集合

var struMu sync.Mutex //struMap写锁

type StruInfo struct {
	TableName    string   //struct 对应表名,默认主键为ID，且必须有，表名配置在主键字段
	FieldNames   []string //struct 字段名
	FieldColumns []string //struct字段对应列名
}

func init() {
	struMap = make(map[string]*StruInfo) //结构对象初始化
}

//model注册
func Register(stru interface{}) {
	st := reflect.TypeOf(stru)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}

	pk, ok := st.FieldByName("ID")
	if !ok {
		logger.Err(st.PkgPath(), st.Name(), "获取主键ID失败！")
	}

	tn := pk.Tag.Get("table") //获取表名

	fn := []string{} //字段名
	fc := []string{} //列名

	for i := 0; i < st.NumField(); i++ {
		sc := st.Field(i).Tag.Get("column")
		if len(sc) > 0 {
			fn = append(fn, st.Field(i).Name)
			fc = append(fc, sc)
		}
	}
	struMap[st.PkgPath()+":"+st.Name()] = &StruInfo{TableName: tn, FieldNames: fn, FieldColumns: fc}

}

//创建数据库表
func CreateTable(stru interface{}) {
	st := reflect.TypeOf(stru)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}

	if si, ok := struMap[st.PkgPath()+":"+st.Name()]; ok { //结构体已注册
		sql := new(bytes.Buffer) //创建sql

		sql.WriteString("CREATE TABLE ")

		sql.WriteString(si.TableName)

		sql.WriteString(" (")

		for i := 0; i < len(si.FieldNames); i++ {

			sft, _ := st.FieldByName(si.FieldNames[i])

			sql.WriteString(si.FieldColumns[i])
			sql.WriteString(" ")
			sql.WriteString(typeToDB(sft.Type))

			if i != len(si.FieldNames)-1 {
				sql.WriteString(",")
			}
		}

		sql.WriteString(");")

		db.DefaultDB().Execute(sql.String())

		logger.Debug(sql.String())

	} else { //结构体未注册
		logger.Err(st.PkgPath(), st.Name(), "结构体未注册！")
	}
}

//删除数据库表
func DropTable(stru interface{}) {
	st := reflect.TypeOf(stru)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}

	if si, ok := struMap[st.PkgPath()+":"+st.Name()]; ok { //结构体已注册
		sql := new(bytes.Buffer) //创建sql

		sql.WriteString("DROP TABLE ")

		sql.WriteString(si.TableName)

		sql.WriteString(";")

		db.DefaultDB().Execute(sql.String())

		logger.Debug(sql.String())

	} else { //结构体未注册
		logger.Err(st.PkgPath(), st.Name(), "结构体未注册！")
	}
}

//保存数据对象
func Save(stru interface{}) {
	st := reflect.TypeOf(stru)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}

	if si, ok := struMap[st.PkgPath()+":"+st.Name()]; ok { //结构体已注册

		names := new(bytes.Buffer)  //字段名
		values := new(bytes.Buffer) //字段值，？占位
		args := []interface{}{}     //参数值

		sv := reflect.Indirect(reflect.ValueOf(stru)) //获取持久化对象的值
		for i := 0; i < len(si.FieldNames); i++ {

			sfv := sv.FieldByName(si.FieldNames[i])

			names.WriteString(",")
			values.WriteString(",")

			names.WriteString(si.FieldColumns[i])
			values.WriteString("?")
			args = append(args, realValue(sfv))
		}

		sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", si.TableName, names.String()[1:], values.String()[1:])

		db.DefaultDB().Execute(sql, args...)

		logger.Debug(sql)

	} else { //结构体未注册
		logger.Err(st.PkgPath(), st.Name(), "结构体未注册！")
	}

}

func SaveAll() {

}

func Query(stru interface{}) {
	st := reflect.TypeOf(stru)
	if st.Kind() == reflect.Ptr {
		st = st.Elem()
	}

	if si, ok := struMap[st.PkgPath()+":"+st.Name()]; ok { //结构体已注册

		names := new(bytes.Buffer) //字段名
		args := []interface{}{}    //参数值

		sv := reflect.Indirect(reflect.ValueOf(stru)) //获取持久化对象的值
		for i := 0; i < len(si.FieldNames); i++ {

			sfv := sv.FieldByName(si.FieldNames[i])

			names.WriteString(",")
			names.WriteString(si.FieldColumns[i])

			args = append(args, realValue(sfv))

		}

		sql := fmt.Sprintf("SELECT %s FROM %s ;", names.String()[1:], si.TableName)

		//		results := [][]interface{}{}

		db.DefaultDB().Execute(sql, args...)

		logger.Debug(sql)

	} else { //结构体未注册
		logger.Err(st.PkgPath(), st.Name(), "结构体未注册！")
	}
}

func QueryAll() {

}

func Update() {

}

func UpdateAll() {

}

func Delete() {

}

//判断各个类型的空值
func nilValue(rt reflect.Type) interface{} {

	switch rt.Kind() {
	//	case reflect.Bool:

	case reflect.Int:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return 0
		//	case reflect.Uint:
		//	case reflect.Uint8:
		//	case reflect.Uint16:
		//	case reflect.Uint32:
		//	case reflect.Uint64:
		//	case reflect.Uintptr:
		//	case reflect.Float32:
		//	case reflect.Float64:
		//	case reflect.Complex64:
		//	case reflect.Complex128:
		//	case reflect.Array:
		//	case reflect.Chan:
		//	case reflect.Func:
		//	case reflect.Interface:
		//	case reflect.Map:
		//	case reflect.Ptr:
		//	case reflect.Slice:
	case reflect.String:
		return ""
	//	case reflect.Struct:
	//	case reflect.UnsafePointer:
	default:
		return ""
	}
}

//根据数据类型，转换为数据库对应字段类型，此处对应关系将依赖于数据库
func typeToDB(rt reflect.Type) string {

	switch rt.Kind() {
	//	case reflect.Bool:

	case reflect.Int:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return "INTEGER"
		//	case reflect.Uint:
		//	case reflect.Uint8:
		//	case reflect.Uint16:
		//	case reflect.Uint32:
		//	case reflect.Uint64:
		//	case reflect.Uintptr:
		//	case reflect.Float32:
		//	case reflect.Float64:
		//	case reflect.Complex64:
		//	case reflect.Complex128:
		//	case reflect.Array:
		//	case reflect.Chan:
		//	case reflect.Func:
		//	case reflect.Interface:
		//	case reflect.Map:
		//	case reflect.Ptr:
		//	case reflect.Slice:
	case reflect.String:
		return "TEXT"
	//	case reflect.Struct:
	//	case reflect.UnsafePointer:
	default:
		return ""
	}
}

//reflect.Value值转换为实际的数据类型，目前仅考虑支持bool、整形、浮点、字符串、日期类型
func realValue(rv reflect.Value) interface{} {
	switch rv.Kind() {
	//	case reflect.Bool:

	case reflect.Int:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return rv.Int()
		//	case reflect.Uint:
		//	case reflect.Uint8:
		//	case reflect.Uint16:
		//	case reflect.Uint32:
		//	case reflect.Uint64:
		//	case reflect.Uintptr:
		//	case reflect.Float32:
		//	case reflect.Float64:
		//	case reflect.Complex64:
		//	case reflect.Complex128:
		//	case reflect.Array:
		//	case reflect.Chan:
		//	case reflect.Func:
		//	case reflect.Interface:
		//	case reflect.Map:
		//	case reflect.Ptr:
		//	case reflect.Slice:
	case reflect.String:
		return rv.String()
	//	case reflect.Struct:
	//	case reflect.UnsafePointer:
	default:
		return nil
	}
}
