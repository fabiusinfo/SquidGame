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

type server struct {
	pb.UnimplementedSquidGameServiceServer
}

func generateID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

var delet int = 1
//habilitar el puerto 8080 en la máquina 162        Javier: listoco, comando aplicado
//acá definir la función sendplays
func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	////// Kathy y Eloli deben implementar la consulta a los 3 datanodes de forma aleatoria para poder
	//enviar la jugada a cualquiera de los 3.
	var direction string

	//var plays_check string
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
	nombreArchivo := "registro.txt" // El nombre o ruta absoluta del archivo
	crearArchivo(nombreArchivo)

	b, errtxt := ioutil.ReadFile(nombreArchivo)

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

//CONSULTA preguntar sobre todas las jugadas en todas las rondas de un determinado jugador
func (s *server) AllPlaysOf(ctx context.Context, in *pb.AllplaysRequest) (*pb.AllplaysReply, error) {
	// Leer jugadas de jugadores que jugaron el juego
	player := in.GetPlayer()
	plays := "Jugadas de "+player+"\n"
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
			num_jugador := linea[0]
			num_ronda := linea[1]
			ip_maquina := linea[2]
			//Ruta := "DN_plays/jugador_" + num_jugador + "__ronda_" + num_ronda + "rv.txt"
			numerojugador := strings.Split(num_jugador, "_") // [Jugador n]
			numeroronda := strings.Split(num_ronda, "_")     // [Ronda 1rv]
			if player == numerojugador[1]{
				plays += "Ronda: "+numeroronda[1]+"\n"
				//ahora vamos a conectarnos con el datanode que tiene el archivo
				conn, err := grpc.Dial(ip_maquina+":8080", grpc.WithInsecure())
				if err != nil {
					panic("cannot connect with server " + err.Error())
				}
				servicePlayer := pb.NewSquidGameServiceClient(conn)
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				arch := "jugador_"+numerojugador[1]+"__ronda_"+num_ronda[1]+".txt"
				r, err := servicePlayer.AllPlaysOf(ctx, &pb.AllplaysRequest{Player: arch})
				plays += r.GetPlays() //agregamos las jugadas a nuestro super string
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}

			}
			//"El jugador: " + numerojugador[1] + " tiene una jugada de la ronda: " + numeroronda[1] + " en la ip: " + ip_maquina
			fmt.Println(plays)
			//s, e = Readln(r)
		}
	}
	
	return &pb.AllplaysReply{Plays: plays}, nil
}

func main() {
	
	/*
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
				num_jugador := linea[0]
				num_ronda := linea[1]
				ip_maquina := linea[2]
				//Ruta := "DN_plays/jugador_" + num_jugador + "__ronda_" + num_ronda + "rv.txt"
				numerojugador := strings.Split(num_jugador, "_") // [Jugador n]
				numeroronda := strings.Split(num_ronda, "_")     // [Ronda 1rv]
				fmt.Println("El jugador: " + numerojugador[1] + " tiene una jugada de la ronda: " + numeroronda[1] + " en la ip: " + ip_maquina)
				s, e = Readln(r)
			}
		} */
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
