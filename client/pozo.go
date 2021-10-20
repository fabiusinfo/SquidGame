package main

import (
	"context"
	//"fmt"
	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"
	//"math/rand"
	//"strconv"
	//"time"
	//"log"
	"net"
)

type server struct {
	pb.UnimplementedSquidGameServiceServer
}


func (s *server) AmountCheck(ctx context.Context, in *pb.AmountRequest) (*pb.AmountReply, error) {
	
	return &pb.AmountReply{Message: "el monto del pozo actual es de 5000 " }, nil
}


func main() {
	listner, err := net.Listen("tcp", ":8080")
	//conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterSquidGameServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())

	

}
}
