package main

import (
	"context"
	"fmt"
	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"

	//"math/rand"
	//"strconv"
	"log"
	"time"
)

func main() {
	var first string
	playerNumber := "1"
	play := "2"
	var stage string
	signed:=false
	//state:="2"
	fmt.Println("ID del jugador: " + playerNumber + " , Jugada: " + play + " , etapa: " + stage)
	fmt.Println("Activar jugador, join->unirse, send->enviar jugadas, amount->solicitar monto: ")
	fmt.Scanln(&first)

	conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	servicePlayer := pb.NewSquidGameServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch first{

	case "join":
		r, err := servicePlayer.JoinGame(ctx, &pb.JoinRequest{Player: playerNumber})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("inscrito: %s", r.GetSigned())
		signed=r.GetSigned()
	case "send":
		r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: playerNumber, Play: play, Stage: stage})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	case "amount":
		message:= "solicito monto"
		r, err := servicePlayer.AmountCheck(ctx, &pb.AmountRequest{Message: message})
		if err != nil {
			log.Fatalf("no se pudo solicitar el monto: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMonto())
	
	default:
		fmt.Println("ingresaste un mal comando.")

	}
	/*if first == "join"{
		r, err := servicePlayer.JoinGame(ctx, &pb.JoinRequest{Player: playerNumber, State: state})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())

	} else if first == "send"{

		r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: playerNumber, Play: play, Stage: stage})
		if err2 != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r2.GetMessage())

	} else if first == "amount"{
		r, err := servicePlayer.AmountCheck(ctx, &pb.AmountRequest{Message: message})
		if err3 != nil {
			log.Fatalf("no se pudo solicitar el monto: %v", err3)
		}
		log.Printf("Greeting: %s", r2.GetMessage())

	}
	else {
		fmt.Println("ingresaste un mal comando.")
	}*/

	
	
	

}
