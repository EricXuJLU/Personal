//mysql的增删改查
package db

import (
	. "MiniDNS/define"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

//插入一对DNS映射，若已存在，则跳过
func Insert(db *sql.DB, table, name, ip string) (id int64){
	exist := Query(db, table, name)
	for _, i := range exist{
		if i == ip {
			fmt.Println(name, ip, "has exist!")
			id = -1
			return
		}
	}
	query := "insert into "+table+" set Name=\""+name+"\", IP=\""+ip+"\", Modify=Now();"
	//fmt.Println(query)
	stmt, err := db.Prepare(query)
	Check(err,2001)
	res, err := stmt.Exec()
	Check(err,2002)
	id, err = res.LastInsertId()
	Check(err,2003)
	fmt.Println("Successfully insert a row whose ID is", id)
	return
}

//按域名查询DNS，返回此域名对应的所有IP
func Query(db *sql.DB, table, name string) (res []string){
	query := "select IP from "+table+" where Name=\""+name+"\";"
	rows, err := db.Query(query)
	Check(err, 2011)
	for rows.Next(){
		var tmp string
		err = rows.Scan(&tmp)
		Check(err, 2012)
		res = append(res, tmp)
	}
	return
}

//更新一对映射，若ipsrc=*，则移除该域名的所有其他映射
func Update(db *sql.DB, table, namesrc, ipsrc, namedst, ipdst string)(aff int64){
	aff = Delete(db, table, namesrc, ipsrc)
	if aff > 0{
		Insert(db, table, namedst, ipdst)
	}
	return
}

//删除一对映射，若ip=*，则该域名所有映射全删
func Delete(db *sql.DB, table, name, ip string)(aff int64){
	if ip == "*" {
		query := "delete from "+table+" where Name=\""+name+"\";"
		stmt, err := db.Prepare(query)
		Check(err, 2031)
		res, err := stmt.Exec()
		Check(err, 2032)
		aff, err = res.RowsAffected()
		Check(err, 2033)
		fmt.Println(aff, "rows has been deleted!")
		return
	}
	query := "delete from "+table+" where Name=\""+name+"\" and IP=\""+ip+"\";"
	stmt, err := db.Prepare(query)
	Check(err, 2034)
	res, err := stmt.Exec()
	Check(err, 2035)
	aff, err = res.RowsAffected()
	Check(err, 2036)
	fmt.Println(aff, "rows has been deleted!")
	return
}
