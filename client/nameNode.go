package main

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"
)

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

type server struct {
	pb.UnimplementedSquidGameServiceServer
}

func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

//habilitar el puerto 8080 en la máquina 162        Javier: listoco, comando aplicado
//acá definir la función sendplays
func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	////// Kathy y Eloli deben implementar la consulta a los 3 datanodes de forma aleatoria para poder
	//enviar la jugada a cualquiera de los 3.
	var direction string

	rand.Seed(time.Now().UnixNano())
	id := rand.Int63n(3)

	if id == 0 {
		direction = "10.6.43.41"
	} else if id == 1 {
		direction = "10.6.43.43"
	} else {
		direction = "10.6.43.44"
	}
	fmt.Println("se reciben los siguientes parametros: player: " + in.GetPlayer() + " ; play:  " + in.GetPlay())
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

	var plays_check string

	// Leer jugadas de jugadores que jugaron el juego
	fmt.Println("--DEMO--")
	fmt.Println("check -> Ver jugadas ")
	fmt.Scanln(&plays_check)
	if plays_check == "check" {
		path := "registro.txt"
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
		s, e := Readln(r)
		for e == nil {
			linea := strings.Split(s, " ")
			jugador := linea[0]
			numerojugador := strings.Split(jugador, "_")
			fmt.Println(numerojugador)
			s, e = Readln(r)
		}
	}
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
