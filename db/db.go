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
func (c *DBContext) Execute(sql string, args ...interface{}) {
	stmt, err := c.db.Prepare(sql)

	errtools.CheckErr(err, "创建Prepare失败:", sql, args)

	result, err := stmt.Exec(args...)

	errtools.CheckErr(err, "创建执行失败:", sql, args)

	lastID, err := result.LastInsertId()

	errtools.CheckErr(err, "创建失败:", sql, args)

	logger.Debug("创建完成：", sql, args, lastID)
}

// Query  "SELECT * FROM users"
func (c *DBContext) Query(sql string, args ...interface{}) {

	//	rows, err := c.db.Query(sql, args...)
	//	errtools.CheckErr(err, "查询失败:", sql, args)

	//	defer rows.Close()

	//	for rows.Next() {
	//		err := rows.Scan(&result)
	//		errtools.CheckErr(err, "查询失败:", sql, args)
	//	}
}

// UPDATE  "UPDATE users SET age = ? WHERE id = ?"
func (c *DBContext) Update(sql string, args ...interface{}) {

	stmt, err := c.db.Prepare(sql)
	errtools.CheckErr(err, "更新Prepare失败:", sql, args)

	result, err := stmt.Exec(args...)
	errtools.CheckErr(err, "更新执行失败:", sql, args)

	affectNum, err := result.RowsAffected()

	errtools.CheckErr(err, "更新失败:", sql, args)

	logger.Debug("更新完成：", sql, args, affectNum)
}

// DELETE "DELETE FROM users WHERE id = ?"
func (c *DBContext) Delete(sql string, args ...interface{}) {
	stmt, err := c.db.Prepare(sql)
	errtools.CheckErr(err, "删除Prepare失败:", sql, args)

	result, err := stmt.Exec(args...)
	errtools.CheckErr(err, "删除执行失败:", sql, args)

	affectNum, err := result.RowsAffected()
	errtools.CheckErr(err, "删除失败:", sql, args)

	logger.Debug("删除完成：", sql, args, affectNum)
}
