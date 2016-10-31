package modules

import (
	//	"bytes"
	//	"encoding/base64"
	//	"encoding/json"
	//	"io/ioutil"
	//	"log"
	//	"net/http"
	//	"os"
	//	"strings"
	"database/sql"
	//	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type SlowInfo struct {
	StartTime   string
	QueryTime   int16
	LockTime    int16
	RowsSent    int32
	rowExamined int32
	Db          string
	SqlText     string
}

type QueryModule struct {
}

//慢查询列表
func (docModule *QueryModule) SlowQuery(date string, page int, count int) (list []SlowInfo, total int, err error) {
	the_time, err := time.Parse("20060102", date)
	if err != nil {
		checkErr(err)
	}

	start_time := strconv.FormatInt(the_time.Unix()-28800, 10)
	end_time := strconv.FormatInt(the_time.Unix()-28800+86399, 10)

	db, err := sql.Open("mysql", "root:123456@/log?charset=utf8")
	if err != nil {
		checkErr(err)
	}
	defer db.Close()

	sql_total := "SELECT COUNT(*) AS `total` FROM `slow_query` WHERE `start_time` BETWEEN '" + start_time + "' AND '" + end_time + "'"
	rows, err := db.Query(sql_total)
	for rows.Next() {
		err = rows.Scan(&total)
		checkErr(err)
	}

	offset := (page - 1) * count
	sql_list := "SELECT `start_time`, `query_time`, `lock_time`, `rows_sent`, `rows_examined`, `db`, `sql_text` FROM `slow_query` WHERE `start_time` BETWEEN '" + start_time + "' AND '" + end_time + "' LIMIT " + strconv.Itoa(offset) + ", " + strconv.Itoa(count)
	rows, err = db.Query(sql_list)
	checkErr(err)
	for rows.Next() {
		var start_time int64
		var query_time int16
		var lock_time int16
		var rows_sent int32
		var rows_examined int32
		var db string
		var sql_text string
		var slow_info SlowInfo
		err = rows.Scan(&start_time, &query_time, &lock_time, &rows_sent, &rows_examined, &db, &sql_text)
		checkErr(err)

		slow_info.StartTime = time.Unix(start_time, 0).Format("2006-01-02 15:04:05")
		slow_info.QueryTime = query_time
		slow_info.LockTime = lock_time
		slow_info.RowsSent = rows_sent
		slow_info.rowExamined = rows_examined
		slow_info.Db = db
		slow_info.SqlText = sql_text

		list = append(list, slow_info)
	}

	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
