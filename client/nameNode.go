package main

import (
	"context"
	//"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"
)

func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

//habilitar el puerto 8080 en la máquina 162
//acá definir la función sendplays
func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {

	return &pb.SendReply{Message: "mensaje de la jugada por constuir equisde buenas noches"}, nil
}

func main() {
	//Acá habilitar el nameNode como servidor.

	conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewSquidGameServiceClient(conn)

	/*res, err := serviceClient.Create(context.Background(), &pb.CreateWishListReq{
		WishList: &pb.WishList{
			Id:   generateID(),
			Name: "my wishlist",
		},
	})*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//aqui primer intento del hello world
	name := "nameNode"
	r, err := serviceClient.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
