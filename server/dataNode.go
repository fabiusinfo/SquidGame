////// Kathy y Eloli deben implementar la estructura del DataNode habilitandolos como servidor.
package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"os"
	
	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSquidGameServiceServer
}

func crearArchivo(var path string) {
	//Verifica que el archivo existe
	var _, err = os.Stat(path)
	//Crea el archivo si no existe
	if os.IsNotExist(err) {
	  var file, err = os.Create(path)
	  if existeError(err) {
		return
	  }
	  defer file.Close()
	}
  }

func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	//aqui implementar la escribicion jugador_1__ronda_1.txt
	var path = "DN_plays/jugador_"+in.GetPlayer()+"__ronda_"+in.GetStage()+".txt"

	crearArchivo(path)
	
	fmt.Println("yo lo recibí")
	return &pb.SendReply{Message: "El DataNode recibió las jugadas con éxito\n" + "El jugador " + in.GetPlayer() + " hizo una jugada " + in.GetPlay() + "en la etapa" + in.GetStage()}, nil
}



func main() {
	// nos convertimos en servidor (dataNode)
	listner, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic("cannot create tcp connection" + err.Error())
	}

	servDN := grpc.NewServer()
	pb.RegisterSquidGameServiceServer(servDN, &server{})

	//esto es lo que estaba al final, no sé donde ponerlo
	if err = servDN.Serve(listner); err != nil {
		log.Printf("paso por el fallo")
		panic("cannot initialize the server" + err.Error())
	}
}
