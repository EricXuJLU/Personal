package main

import (
	. "MiniDNS/define"
	pb "MiniDNS/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"strings"
)

const (
	//format:
	//get name
	//add name ip
	//delete name ip(*)
	//update namesrc ipsrc(*) namedst ipdst
	ORDER = "delete www.google.com *"
)

func main() {
	//客户端连接服务端
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	Check(err, 6001)
	defer  conn.Close()
	//获得grpc句柄
	c := pb.NewDNSClient(conn)
	handle(c)

}
func handle(c pb.DNSClient){
	order := strings.Split(ORDER, " ")
	switch strings.ToLower(order[0]) {
	case "get": handleGet(c, order[1]);
	case "add": handleAdd(c, order[1], order[2]);
	case "delete": handleDelete(c, order[1], order[2]);
	case "update": handleUpdate(c, order[1], order[2], order[3], order[4]);
	default:
		fmt.Println("Invalid Order!")
	}
}

func handleGet(c pb.DNSClient, name string){
	re1, err := c.GetIP(context.Background(), &pb.STR{Content: name})
	Check(err, 6002)
	fmt.Println("此域名对应的IP有：")
	for _, i := range re1.Contents{
		fmt.Println(i)
	}
}
func handleAdd(c pb.DNSClient, name, ip string){
	re2, err := c.AddDNS(context.Background(), &pb.Pair{Name: name, IP: ip})
	Check(err, 6003)
	fmt.Println(re2)
}
func handleDelete(c pb.DNSClient, name, ip string){
	re2 ,err := c.DeleteDNS(context.Background(), &pb.Pair{
		Name: name,
		IP:   ip,
	})
	Check(err, 6004)
	fmt.Println(re2)
}
func handleUpdate(c pb.DNSClient, namesrc, ipsrc, namedst, ipdst string){
	re2 ,err := c.ModifyDNS(context.Background(), &pb.Pairs{
		SRC: &pb.Pair{Name: namesrc, IP: ipsrc},
		DST: &pb.Pair{Name: namedst, IP: ipdst},
	})
	Check(err, 6005)
	fmt.Println(re2)
}
/*
re2, err := c.AddDNS(context.Background(), &pb.Pair{Name: "www.baidu.com", IP: "2.2.2.2"})
	Check(err, 6003)
	fmt.Println("re2=",re2)
	re2, err = c.AddDNS(context.Background(), &pb.Pair{Name: "www.baidu.com", IP: "4.4.4.4"})
	Check(err, 6004)
	fmt.Println("re2=",re2)
	re2, err = c.AddDNS(context.Background(), &pb.Pair{Name: "www.baidu.com", IP: "3.3.3.3"})
	Check(err, 6005)
	fmt.Println("re2=",re2)

	//通过句柄进行调用服务端函数


	re2 ,err = c.ModifyDNS(context.Background(), &pb.Pairs{
		SRC: &pb.Pair{Name: "www.baidu.com", IP: "*"},
		DST: &pb.Pair{Name: "www.google.com", IP: "11.11.11.11"},
	})
	Check(err, 6007)
	fmt.Println(re2)
	//检查结果
	re1, err = c.GetIP(context.Background(), &pb.STR{Content: "www.baidu.com"})
	Check(err, 6006)
	fmt.Println("此域名对应的IP有：")
	for _, i := range re1.Contents{
		fmt.Println(i)
	}

	re1, err = c.GetIP(context.Background(), &pb.STR{Content: "www.google.com"})
	Check(err, 6006)
	fmt.Println("此域名对应的IP有：")
	for _, i := range re1.Contents{
		fmt.Println(i)
	}

re2 ,err = c.DeleteDNS(context.Background(), &pb.Pair{
Name: "www.baidu.com",
IP:   "4.4.4.4",
})
Check(err, 6007)
fmt.Println(re2)
//检查结果
re1, err = c.GetIP(context.Background(), &pb.STR{Content: "www.baidu.com"})
Check(err, 6006)
fmt.Println("此域名对应的IP有：")
for _, i := range re1.Contents{
fmt.Println(i)
}

re2 ,err = c.DeleteDNS(context.Background(), &pb.Pair{
Name: "www.baidu.com",
IP:   "*",
})
Check(err, 6007)
fmt.Println(re2)
//检查结果
re1, err = c.GetIP(context.Background(), &pb.STR{Content: "www.baidu.com"})
Check(err, 6006)
fmt.Println("此域名对应的IP有：")
for _, i := range re1.Contents{
fmt.Println(i)
}
//modify和delete都待测，尤其是批量
 */