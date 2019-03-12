package main

import (
	"fmt"
	"math"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"time"
)

func main() {
	logs, recordCount, err := GetLogs("123", "0000-01-02-15:04:05", "2319-08-23-17:00:00", 3, 3)
	fmt.Println("logs:", logs)
	fmt.Println("recordCount:", recordCount)
	fmt.Println("err:", err)
	// db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8")
	// gorm.Open("mysql", argsStr)

	// checkErr(err)
	// 插入数据
	// stmt, err := db.Prepare("INSERT user_info SET username=?,departname=?,created=?")
	// checkErr(err)
	// res, err := stmt.Exec("test", " 研发部门", "2017-12-09")
	// checkErr(err)
	// id, err := res.LastInsertId()
	// checkErr(err)
	// fmt.Println(id)

	// // 更新数据
	// stmt, err = db.Prepare("update user_info set username=? where uid=?")
	// checkErr(err)
	// res, err = stmt.Exec("test", id)
	// checkErr(err)
	// affect, err := res.RowsAffected()
	// checkErr(err)
	// fmt.Println(affect)

	// // 查询数据
	// rows, err := db.Query("SELECT * FROM user_info")
	// checkErr(err)
	// for rows.Next() {
	// 	var uid int
	// 	var username string
	// 	var department string
	// 	var created string
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(department)
	// 	fmt.Println(created)
	// }

	// // 删除数据
	// stmt, err = db.Prepare("delete from user_info where uid=?")
	// checkErr(err)
	// //res, err = stmt.Exec(id)
	// checkErr(err)
	// affect, err = res.RowsAffected()
	// checkErr(err)
	// fmt.Println(affect)
	// db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Log
type Log struct {
	// gorm.Model
	Date   string `json:"date"`
	Detail string `json:"detail"`
	Type   string `json:"type"`
	User   string `json:"user"`
}

// TableName 设置表名
func (Log) TableName() string {
	return "sys_log"
}

// GetLogs 根据查询条件，获取数据（带分页）
func GetLogs(logType, startDate, endDate string, currentPage, pageSize int) ([]Log, int, error) { // (pageCount int) {
	logs := make([]Log, 0)
	var count int
	var recordCount int
	// 获取数据库连接
	Db, err := gorm.Open("mysql", "root:123456@/test?charset=utf8") //db.GetGorm()
	if err != nil {
		return logs, recordCount, err
	}

	// 获取数据总数量
	Db.Table("sys_log").Count(&count)
	fmt.Println("count", count)
	recordCount = int(math.Ceil(float64(count) / float64(pageSize)))

	// 计算总页数
	if count > 0 && pageSize > 0 {
		recordCount = int(math.Ceil(float64(count) / float64(pageSize)))
		fmt.Println("recordCount", recordCount)
	} else {
		return logs, 0, nil
	}

	// 根据时间排序
	Db = Db.Order("date")

	// 限制开始结束时间
	Db = Db.Where("date >= ? and date <= ?", startDate, endDate)

	// 限制到选择的页数
	if currentPage > 0 && pageSize > 0 {
		Db = Db.Limit(pageSize).Offset((currentPage - 1) * pageSize)
	}

	// 获取数据
	if err := Db.Find(&logs).Error; err != nil {
		fmt.Println(err.Error())
	}
	return logs, recordCount, nil
}
