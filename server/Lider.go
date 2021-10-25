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

func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	return &pb.JoinReply{Signed: true}, nil
}

func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	conn, err := grpc.Dial("10.6.43.42:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceLider := pb.NewSquidGameServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := serviceLider.SendPlays(ctx, &pb.SendRequest{Player: in.GetPlayer(), Play: in.GetPlay(), Stage: in.GetStage()})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	return &pb.SendReply{Message: "El Lider Recibió tu jugada con éxito, se las envía al nameNode"}, nil
}
	//"El jugador " + in.GetPlayer() + " hizo una jugada " + in.GetPlay() + "en la etapa" + in.GetStage()

func (s *server) AmountCheck(ctx context.Context, in *pb.AmountRequest) (*pb.AmountReply, error) {
	message:="solicito monto"
	//conexión con el pozo
	conn, err := grpc.Dial("10.6.43.43:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with pozo " + err.Error())
	}
	serviceClient := pb.NewSquidGameServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//aqui primer intento del consultar desde el servidor a otra entidad.
	r, err := serviceClient.AmountCheck(ctx, &pb.AmountRequest{Message: message})
	if err != nil {
		log.Fatalf("no se pudo solicitar el monto: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMonto())
	return &pb.AmountReply{Monto: r.GetMonto()}, nil
	}

func main() {
	//códigos Etapas
	//1rv
	//2tc
	//3tn

	go func() {
		// nos convertimos en servidor (LIDER)
		listner, err := net.Listen("tcp", ":8080")

		if err != nil {
			panic("cannot create tcp connection" + err.Error())
		}

		serv := grpc.NewServer()
		pb.RegisterSquidGameServiceServer(serv, &server{})

		//esto es lo que estaba al final, no sé donde ponerlo
		if err = serv.Serve(listner); err != nil {
			log.Printf("paso por el fallo")
			panic("cannot initialize the server" + err.Error())
		}

	}()

	///////////// Interfaz

	var start string
	var playerAmount int
	var stage string
	stage ="1"
	fmt.Println("ingresa la cantidad de jugadores: ")
	fmt.Scanln(&playerAmount)

	if playerAmount == 16 {
		//se da inicio al juego
		fmt.Println("escribe start para comenzar: ")
		fmt.Scanln(&start)
		if start == "start"{
			fmt.Println("Ha comenzado la etapa: "+stage)
		}
		fmt.Println("se ha muerto ste men: 2")
		fmt.Println("los jugadores vivos que pasan a la siguiente ronda son 16")
		fmt.Println("los ganadores de la ronda son 1,2,3 ")
		stage="2"

		fmt.Println("escribe start para comenzar: ")
		fmt.Scanln(&start)
		if start == "start"{
			fmt.Println("Ha comenzado la etapa: "+stage)
		}
		fmt.Println("se ha muerto ste men: 2")
		fmt.Println("los jugadores vivos que pasan a la siguiente ronda son 16")
		fmt.Println("los ganadores de la ronda son 1,2,3 ")
		stage="3"

		fmt.Println("escribe start para comenzar: ")
		fmt.Scanln(&start)
		if start == "start"{
			fmt.Println("Ha comenzado la etapa: "+stage)
		}
		fmt.Println("se ha muerto ste men: 2")
		fmt.Println("los jugadores vivos que pasan a la siguiente ronda son 16")
		fmt.Println("los ganadores de la ronda son 1,2,3 ")
		stage="4"

	}
	//enviar

	

	// NAMENOOOOOOOOOOOOOOOOOOOOOODEEEEEEEEEEEEEEE
	/*conn, err := grpc.Dial("10.6.43.42:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceLider := pb.NewSquidGameServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	playerID := "1"
	stage := "1"
	jugada := "5"

	r, err := serviceLider.SendPlays(ctx, &pb.SendRequest{Player: playerID, Play: jugada, Stage: stage})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())/*

	//////////////////////

	// POZOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO

	/*var second string
	message := "solicitar"

	fmt.Println("ingresa la letra a para solicitar monto: ")
	fmt.Scanln(&second)

	//if true {} aqui parte el POZO
	conn2, err2 := grpc.Dial("10.6.43.43:8080", grpc.WithInsecure())

	if err2 != nil {
		panic("cannot connect with pozo " + err.Error())
	}
	serviceClient := pb.NewSquidGameServiceClient(conn2)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//aqui primer intento del consultar desde el servidor a otra entidad.
	r2, err3 := serviceClient.AmountCheck(ctx, &pb.AmountRequest{Message: message})
	if err3 != nil {
		log.Fatalf("no se pudo solicitar el monto: %v", err3)
	}
	log.Printf("Greeting: %s", r2.GetMessage())
	*/
	/////////////////////////////

}
