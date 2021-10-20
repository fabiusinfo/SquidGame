package main

import (
	"context"
	//"fmt"
	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedSquidGameServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Wena los k, quien quiere wones " + in.GetName()}, nil
}

func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	return &pb.JoinReply{Message: "Jugador" + in.GetPlayer() +"se unio al Juego"+ in.GetState() + ", suerte calamar, o algo asi no vi la serie " }, nil
}

func (s *server) SendPlay(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	return &pb.SendReply{Message:"El jugador "+ in.GetPlayer()+" hizo una jugada " + in.GetPlay()}, nil
}




func main() {
	listner, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterSquidGameServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

}