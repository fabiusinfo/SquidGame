////// Kathy y Eloli deben implementar la estructura del DataNode habilitandolos como servidor.
package main

import (
	"bufio"
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

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
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

func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	//aqui implementar la escribicion jugador_1__ronda_1.txt
	var path = "DN_plays/jugador_" + in.GetPlayer() + "__ronda_" + in.GetStage() + ".txt"

	crearArchivo(path)

	// añadir al texto
	b, errtxt := ioutil.ReadFile(path)

	if errtxt != nil {
		log.Fatal(errtxt)
	}

	b = append(b, []byte(in.GetPlay()+" \n")...) //hay que conseguirse la jugada, quizas en el mercado negro hay
	errtxt = ioutil.WriteFile(path, b, 0644)

	if errtxt != nil {
		log.Fatal(errtxt)
	}

	fmt.Println("yo lo recibí")
	return &pb.SendReply{Stage: "Amongus", Alive: true}, nil
}

//CONSULTA preguntar sobre todas las jugadas en todas las rondas de un determinado jugador
func (s *server) AllPlaysOf(ctx context.Context, in *pb.AllplaysRequest) (*pb.AllplaysReply, error) {
	// Leer jugadas de jugadores que jugaron el juego
	fmt.Println("holita, si entre uwu")
	path := in.GetPlayer() //aqui recibe el nombre del archivo
	plays2 := ""
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(file)
	ese, e := Readln(r)
	for e == nil {
		plays2 += string(ese)
	}
	return &pb.AllplaysReply{Plays: plays2}, nil
}

func main() {
	// nos convertimos en servidor (dataNode)
	go func() {
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
}
