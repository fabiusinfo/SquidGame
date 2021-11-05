package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"
)

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

type server struct {
	pb.UnimplementedSquidGameServiceServer
}

func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	//enviar la jugada a cualquiera de los 3.
	var direction string

	if delet == 1 {
		delet = 0
		nombreArchivo := "registro.txt" // El nombre o ruta absoluta del archivo
		err := os.Remove(nombreArchivo)
		if err != nil {
			fmt.Printf("Error eliminando archivo: %v\n", err)
		} else {
			fmt.Println("Eliminado correctamente")
		}
	}

	rand.Seed(time.Now().UnixNano())
	id := rand.Int63n(3)
	if id == 0 {
		direction = "10.6.43.41" // maquina 1
	} else if id == 1 {
		direction = "10.6.43.43" // maquina 3
	} else {
		direction = "10.6.43.44" // maquina 4
	}
	fmt.Println("Se reciben los siguientes parametros: Player: " + in.GetPlayer() + " ; Play:  " + in.GetPlay())
	conn, err := grpc.Dial(direction+":9000", grpc.WithInsecure())
	serviceNN := pb.NewSquidGameServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := serviceNN.SendPlays(ctx, &pb.SendRequest{Player: in.GetPlayer(), Play: in.GetPlay(), Stage: in.GetStage()})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetStage())
	// añadir al texto
	/*
		nombreArchivo := "registro.txt" // El nombre o ruta absoluta del archivo
		crearArchivo(nombreArchivo)
	*/

	b, errtxt := ioutil.ReadFile("registro.txt")

	if errtxt != nil {
		log.Fatal(errtxt)
	}

	b = append(b, []byte("Jugador_"+in.GetPlayer()+" Ronda_"+in.GetStage()+" "+direction+"\n")...)
	errtxt = ioutil.WriteFile("registro.txt", b, 0644)

	if errtxt != nil {
		log.Fatal(errtxt)
	}

	return &pb.SendReply{Stage: "www", Alive: true}, nil
}

func main() {

	// nos convertimos en servidor (NameNode)
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

	var first string

	fmt.Println("aqui recibimos las jugadas del lider")
	fmt.Scanln(&first)
}
