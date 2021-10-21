package main

import (
	"context"
	"fmt"
	"net"

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

//habilitar el puerto 8080 en la máquina 162        Javier: listoco, comando aplicado
//acá definir la función sendplays
func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	return &pb.SendReply{Message: "El jugador " + in.GetPlayer() + " hizo una jugada " + in.GetPlay() + "en la etapa" + in.GetStage()}, nil
}

func main() {
	//Acá habilitar el nameNode como servidor (con el líder)

	listner, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	servNN := grpc.NewServer()
	pb.RegisterSquidGameServiceServer(servNN, &server{})

	var first string
	message := "recibir_jugada"

	fmt.Println("aqui recibimos las jugadas del lider")
	fmt.Scanln(&first)

	// ???? se va o no
	if err = serv.Serve(listner); err != nil {
		log.Printf("paso por el fallo")
		panic("cannot initialize the server" + err.Error())
	}

	

	/*conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceNameNode := pb.NewSquidGameServiceClient(conn)

	/*res, err := serviceClient.Create(context.Background(), &pb.CreateWishListReq{
		WishList: &pb.WishList{
			Id:   generateID(),
			Name: "my wishlist",
		},
	}) /
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//aqui primer intento del hello world
	name := "nameNode"
	r, err := serviceClient.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	*/
}
