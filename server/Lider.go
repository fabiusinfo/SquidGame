package main

import (
	"context"
	"fmt"
	"net"
	"log"
	"time"
	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSquidGameServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Wena los k, quien quiere wones " + in.GetName()}, nil
}

func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	return &pb.JoinReply{Message: "Jugador" + in.GetPlayer() + "se unio al Juego" + in.GetState() + ", suerte calamar, o algo asi no vi la serie "}, nil
}

func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	return &pb.SendReply{Message: "El jugador " + in.GetPlayer() + " hizo una jugada " + in.GetPlay()}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":8080")

	conn, err2 := grpc.Dial("10.6.43.43:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	if err2 != nil {
		panic("cannot connect with pozo " + err.Error())
	}

	serv := grpc.NewServer()
	serviceClient := pb.NewSquidGameServiceClient(conn)
	pb.RegisterSquidGameServiceServer(serv, &server{})
	if err = serv.Serve(listner); err != nil {
		panic("cannot initialize the server" + err.Error())
	}
	var first string

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()
	fmt.Scanln(&first)
    fmt.Println("ingresa la letra a para solicitar monto: ")


	if (first=="a"){
		//aqui primer intento del consultar desde el servidor a otra entidad.
	r, err := serviceClient.AmountCheck(ctx, &pb.AmountRequest{Message: first})
	if err != nil {
		log.Fatalf("no se pudo solicitar el monto: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	}


}
