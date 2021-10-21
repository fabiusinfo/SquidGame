package main

import (
	"context"
	"fmt"
	"log"
	"net"
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
	return &pb.SendReply{Message: "El jugador " + in.GetPlayer() + " hizo una jugada " + in.GetPlay() + "en la etapa" + in.GetEtapa()}, nil
}

func main() {

	// nos convertimos en servidor (LIDER)
	listner, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterSquidGameServiceServer(serv, &server{})

	var first string
	//enviar

	fmt.Println("ingresa la letra a para eviar jugadas: ")
	fmt.Scanln(&first)

	// NAMENOOOOOOOOOOOOOOOOOOOOOODEEEEEEEEEEEEEEE
	conn, err := grpc.Dial("10.6.43.42:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceLider := pb.NewSquidGameServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	playerID := "1"
	etapa := "1"
	jugada := "5"

	r, err := serviceLider.SendPlays(ctx, &pb.SendRequest{Player: playerID, Play: jugada, Stage: stage})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	//////////////////////

	// POZOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO

	var second string
	message := "solicitar"

	fmt.Println("ingresa la letra a para solicitar monto: ")
	fmt.Scanln(&second)

	//if true {} aqui parte el POZO
	conn, err2 := grpc.Dial("10.6.43.43:8080", grpc.WithInsecure())

	if err2 != nil {
		panic("cannot connect with pozo " + err.Error())
	}
	//serviceClient := pb.NewSquidGameServiceClient(conn)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//aqui primer intento del consultar desde el servidor a otra entidad.
	r, err := serviceClient.AmountCheck(ctx, &pb.AmountRequest{Message: message})
	if err != nil {
		log.Fatalf("no se pudo solicitar el monto: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	//Aquí llamar el sendplays del Namenode javier.
	if true {
		conn, err2 := grpc.Dial("10.6.43.42:8080", grpc.WithInsecure())

		if err2 != nil {
			panic("cannot connect with nameNode " + err.Error())
		}
		serviceClient := pb.NewSquidGameServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		//aqui primer intento del consultar desde el servidor a otra entidad.
		r, err := serviceClient.SendPlays(ctx, &pb.SendRequest{Message: message})
		if err != nil {
			log.Fatalf("no se pudo enviar la jugada: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}

	if err = serv.Serve(listner); err != nil {
		log.Printf("paso por el fallo")
		panic("cannot initialize the server" + err.Error())
	}

}
