package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSquidGameServiceServer
}

func existeError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func crearArchivo(path string) {
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

var delet int = 1

func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	//aqui implementar la escritura del archivo de texto
	var path = "DN_plays/jugador_" + in.GetPlayer() + "__ronda_" + in.GetStage() + ".txt"

	if delet == 1 {
		delet = 0
		nombreArchivo := path // El nombre o ruta absoluta del archivo
		err := os.Remove(nombreArchivo)
		if err != nil {
			fmt.Printf("Error eliminando archivo: %v\n", err)
		} else {
			fmt.Println(" ")
		}
	}

	crearArchivo(path)

	// añadir al texto
	b, errtxt := ioutil.ReadFile(path)

	if errtxt != nil {
		log.Fatal(errtxt)
	}

	b = append(b, []byte(in.GetPlay()+" \n")...)
	errtxt = ioutil.WriteFile(path, b, 0644)

	if errtxt != nil {
		log.Fatal(errtxt)
	}

	fmt.Println("Se recibe... Player: "+ in.GetPlayer() + " / Play:  " + in.GetPlay()+ " / Stage: "+ in.GetStage())
	return &pb.SendReply{Stage: "Amongus", Alive: true}, nil
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
		log.Printf("Paso por el fallo")
		panic("cannot initialize the server" + err.Error())
	}
}
