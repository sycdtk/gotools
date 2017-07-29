package db

import (
	"database/sql"

	"github.com/sycdtk/gotools/config"
	"github.com/sycdtk/gotools/errtools"
	"github.com/sycdtk/gotools/logger"

	_ "github.com/mattn/go-sqlite3" // sqlite3 dirver
)

var dbPool map[string]*DBContext //多类型数据库连接池

var defalutDB string //默认数据库

//初始化数据库
func init() {

	//初始化数据库驱动1
	db1 := config.Read("db", "db1")
	defalutDB = db1 //默认数据库

	driver1 := config.Read("db", "driver1")
	conn1 := config.Read("db", "conn1")

	dbPool = make(map[string]*DBContext)

	dbc1, err := connectDB(driver1, conn1)
	errtools.CheckErr(err, "连接创建失败！")

	dbPool[db1] = dbc1

}

//选择数据库连接
func DB(dbName string) *DBContext {

	if dbName == "" {
		dbName = defalutDB
	}

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
// UPDATE  "UPDATE users SET age = ? WHERE id = ?"
// DELETE "DELETE FROM users WHERE id = ?"
// Create "CREATE TABLE(...)"
// Drop "DROP TABLE..."
func (c *DBContext) Execute(execSql string, args ...interface{}) {
	stmt, err := c.db.Prepare(execSql)

	errtools.CheckErr(err, "SQL Prepare失败:", execSql, args)

	result, err := stmt.Exec(args...)

	errtools.CheckErr(err, "SQL 执行失败:", execSql, args)

	lastID, _ := result.LastInsertId()

	affectNum, _ := result.RowsAffected()

	logger.Debug("SQL执行完成：", execSql, args, "，最后插入ID：", lastID, "，受影响行数：", affectNum)
}

// Query  "SELECT * FROM users"
func (c *DBContext) Query(querySql string, args ...interface{}) [][]sql.RawBytes {

	rows, err := c.db.Query(querySql, args...)
	errtools.CheckErr(err, "SQL 查询失败:", querySql, args)

	defer rows.Close()

	cols, err := rows.Columns() // 获取列数
	errtools.CheckErr(err, "SQL 获取结果失败:", querySql, args)

	var results [][]sql.RawBytes

	rowValue := make([]sql.RawBytes, len(cols))

	row := make([]interface{}, len(cols))

	for i := range rowValue {
		row[i] = &rowValue[i]
	}

	for rows.Next() {
		err = rows.Scan(row...)
		errtools.CheckErr(err, "SQL 结果解析失败:", querySql, args)

		results = append(results, rowValue)
	}

	logger.Debug("SQL 查询完成：", querySql, args)

	return results
}

//检查数据库表是否存在
func (c *DBContext) TableExist(tableName string) bool {
	return len(c.Query("select name from sqlite_master where type='table' and name=?", tableName)) > 0
}
