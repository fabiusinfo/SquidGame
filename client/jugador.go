package main

import (
	"context"
	//"fmt"
	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"

	//"math/rand"
	//"strconv"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	servicePlayer := pb.NewSquidGameServiceClient(conn)

	/*res, err := serviceClient.Create(context.Background(), &pb.CreateWishListReq{
		WishList: &pb.WishList{
			Id:   generateID(),
			Name: "my wishlist",
		},
	})*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	playerNumber := "1"
	state := "2"
	r, err := servicePlayer.JoinGame(ctx, &pb.JoinRequest{Player: playerNumber, State: state})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	play := "5"
	r2, err2 := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: playerNumber, Play: play})
	if err2 != nil {
		log.Fatalf("could not greet: %v", err2)
	}
	log.Printf("Greeting: %s", r2.GetMessage())

}
