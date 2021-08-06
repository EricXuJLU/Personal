package main

import (
	. "MiniDNS/define"
	pb "MiniDNS/protos"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main(){
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDNSServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
