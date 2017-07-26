package db

import (
	"database/sql"

	"github.com/sycdtk/gotools/config"
	"github.com/sycdtk/gotools/errtools"
	"github.com/sycdtk/gotools/logger"

	_ "github.com/mattn/go-sqlite3" // sqlite3 dirver
)

var dbPool map[string]*DBContext //多类型数据库连接池

//初始化数据库
func init() {

	driver1 := config.Read("db", "driver1")
	conn1 := config.Read("db", "conn1")

	dbPool = make(map[string]*DBContext)

	dbc1, err := connectDB(driver1, conn1)
	errtools.CheckErr(err, "连接创建失败！")

	dbPool[driver1] = dbc1
}

//默认数据库
func DefaultDB() *DBContext {
	defaultDB := config.Read("db", "default")
	if dbc, ok := dbPool[defaultDB]; ok {
		return dbc
	}
	return nil
}

//选择数据库连接
func ChooseDB(dbName string) *DBContext {
	if dbc, ok := dbPool[dbName]; ok {
		return dbc
	}
	return nil
}

type DBContext struct {
	db *sql.DB
}

//连接数据库
func connectDB(driverName string, dbName string) (*DBContext, error) {
	db, err := sql.Open(driverName, dbName)
	errtools.CheckErr(err, "连接创建失败:", driverName, dbName)

	if err := db.Ping(); err != nil {
		return nil, errtools.NewErr("数据库连接无效！")
	}

	return &DBContext{db}, nil
}

// Execute  "INSERT INTO users(name,age) values(?,?)"
func (c *DBContext) Execute(execSql string, args ...interface{}) {
	stmt, err := c.db.Prepare(execSql)

	errtools.CheckErr(err, "创建Prepare失败:", execSql, args)

	result, err := stmt.Exec(args...)

	errtools.CheckErr(err, "创建执行失败:", execSql, args)

	lastID, err := result.LastInsertId()

	errtools.CheckErr(err, "创建失败:", execSql, args)

	logger.Debug("创建完成：", execSql, args, lastID)
}

// Query  "SELECT * FROM users"
func (c *DBContext) Query(querySql string, args ...interface{}) [][]sql.RawBytes {

	rows, err := c.db.Query(querySql, args...)
	errtools.CheckErr(err, "查询失败:", querySql, args)

	defer rows.Close()

	cols, err := rows.Columns() // 获取列数
	errtools.CheckErr(err, "查询失败:", querySql, args)

	var results [][]sql.RawBytes

	rowValue := make([]sql.RawBytes, len(cols))

	row := make([]interface{}, len(cols))

	for i := range rowValue {
		row[i] = &rowValue[i]
	}

	for rows.Next() {
		err = rows.Scan(row...)
		errtools.CheckErr(err, "查询失败:", querySql, args)

		results = append(results, rowValue)
	}

	logger.Debug("查询完成：", querySql, args)

	return results
}

// UPDATE  "UPDATE users SET age = ? WHERE id = ?"
func (c *DBContext) Update(updateSql string, args ...interface{}) {

	stmt, err := c.db.Prepare(updateSql)
	errtools.CheckErr(err, "更新Prepare失败:", updateSql, args)

	result, err := stmt.Exec(args...)
	errtools.CheckErr(err, "更新执行失败:", updateSql, args)

	affectNum, err := result.RowsAffected()

	errtools.CheckErr(err, "更新失败:", updateSql, args)

	logger.Debug("更新完成：", updateSql, args, affectNum)
}

// DELETE "DELETE FROM users WHERE id = ?"
func (c *DBContext) Delete(delSql string, args ...interface{}) {
	stmt, err := c.db.Prepare(delSql)
	errtools.CheckErr(err, "删除Prepare失败:", delSql, args)

	result, err := stmt.Exec(args...)
	errtools.CheckErr(err, "删除执行失败:", delSql, args)

	affectNum, err := result.RowsAffected()
	errtools.CheckErr(err, "删除失败:", delSql, args)

	logger.Debug("删除完成：", delSql, args, affectNum)
}
