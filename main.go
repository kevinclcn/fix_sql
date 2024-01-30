package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// 过滤掉错误的sql
func main() {
	sqlFile, err := os.Open("/opt/tmp/数据/dml/customer_refund_order.sql")
	if err != nil {
		panic(err)
	}
	defer sqlFile.Close()

	sqlScanner := bufio.NewScanner(sqlFile)
	sqlScanner.Split(bufio.ScanLines)

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	sqlFileTarget, err := os.OpenFile("/opt/tmp/数据/dml/customer_refund_order_fixed.sql", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer sqlFileTarget.Close()

	for sqlScanner.Scan() {
		sql := sqlScanner.Text()

		_, err := db.Exec(sql)
		if err != nil {
			fmt.Println(err)
		} else {
			sqlFileTarget.WriteString(sql + "\n")
		}
	}
}
