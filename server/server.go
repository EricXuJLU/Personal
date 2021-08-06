package main

import (
	db2 "MiniDNS/db"
	. "MiniDNS/define"
	"MiniDNS/protos"
	"context"
	"database/sql"
	_ "gorm.io/driver/mysql"
	"strconv"
)

type Server struct {}

func (server *Server) GetIP(ctx context.Context, name *protos.STR) (*protos.STRs, error) {
	db, err := sql.Open("mysql", "root:root@/MiniDNS")
	Check(err,1001)
	defer db.Close()
	ip := new(protos.STRs)
	ip.Contents = db2.Query(db, "dns", name.Content)
	return ip, err
}

func (server *Server) AddDNS(ctx context.Context, pair *protos.Pair) (*protos.STR, error) {
	db, err := sql.Open("mysql", "root:root@/MiniDNS")
	Check(err,1001)
	defer db.Close()
	res := new(protos.STR)
	id := db2.Insert(db, "dns", pair.Name, pair.IP)
	if id == -1 {
		res.Content = "当前映射已存在，插入失败"
	} else{
		res.Content = "插入成功"
	}
	return res, err
}

func (server *Server) DeleteDNS(ctx context.Context, pair *protos.Pair) (*protos.STR, error) {
	db, err := sql.Open("mysql", "root:root@/MiniDNS")
	Check(err,1001)
	defer db.Close()
	aff := db2.Delete(db, "dns", pair.Name, pair.IP)
	res := new(protos.STR)
	res.Content = strconv.FormatInt(aff, 10) + "条记录被删除"
	return res, err
}

func (server *Server) ModifyDNS(ctx context.Context, pairs *protos.Pairs) (*protos.STR, error) {
	db, err := sql.Open("mysql", "root:root@/MiniDNS")
	Check(err,1001)
	defer db.Close()
	aff := db2.Update(db, "dns", pairs.SRC.Name, pairs.SRC.IP, pairs.DST.Name, pairs.DST.IP)
	//aff = int64(math.Max(float64(aff), 1))
	res := new(protos.STR)
	res.Content = "修改成功，"+strconv.FormatInt(aff, 10)+"条记录被影响到"
	return res, err
}
