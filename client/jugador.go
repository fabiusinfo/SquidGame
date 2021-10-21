package main

import (
	"context"
	"fmt"
	//"fmt"
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
	stage := "3"
	//state:="2"
	fmt.Println("Activar jugador: ")
	fmt.Scanln(&first)

	conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	servicePlayer := pb.NewSquidGameServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// pa despues
	/*r, err := servicePlayer.JoinGame(ctx, &pb.JoinRequest{Player: playerNumber, State: state})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())*/
	r2, err2 := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: playerNumber, Play: play, Stage: stage})
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	log.Printf("Greeting: %s", r2.GetMessage())

}
