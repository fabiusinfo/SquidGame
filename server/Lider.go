package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	pb "github.com/fabiusinfo/SquidGame/proto"
	amqp "github.com/rabbitmq/amqp091-go"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSquidGameServiceServer
}

type PlayerStruct struct {
	id    string
	alive bool
	score int
}

var liderPlay int
var actualStage string
var actualRound int32
var started bool
var list_of_players []PlayerStruct

//listas stage 2
var group1 []PlayerStruct
var group2 []PlayerStruct

//var players [16]string
var totalPlayers int

// Error para el Rabbit
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (s *server) JoinGame(ctx context.Context, in *pb.JoinRequest) (*pb.JoinReply, error) {
	//players[in.GetPlayer()] = "alive"
	totalPlayers += 1
	list_of_players = append(list_of_players, PlayerStruct{in.GetPlayer(), true, 0})
	return &pb.JoinReply{Codes1: "1rv", Codes2: "2tc", Codes3: "3tn"}, nil
}

func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	alive := true
	if actualRound != 0 {
		if in.GetRound() == actualRound {

			//envío al nameNode

			/*

				conn, err := grpc.Dial("10.6.43.42:8080", grpc.WithInsecure())

				if err != nil {
					panic("cannot connect with server " + err.Error())
				}

				serviceLider := pb.NewSquidGameServiceClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				r, err := serviceLider.SendPlays(ctx, &pb.SendRequest{Player: in.GetPlayer(), Play: in.GetPlay(), Stage: in.GetStage(), Round:in.GetRound()})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
			*/

			//Envío al Pozo

			if started == true {
				pPlay, errpPlay := strconv.Atoi(in.GetPlay())
				if errpPlay != nil {
					log.Fatalf("could not greet: %v", errpPlay)
				}

				for i := 0; i < 16; i++ {
					if list_of_players[i].id == in.GetPlayer() {

						list_of_players[i].score += pPlay
					}
				}

				if actualStage == "1rv" {

					if pPlay > liderPlay {
						alive = false
						for i := 0; i < 16; i++ {
							if list_of_players[i].id == in.GetPlayer() {
								list_of_players[i].alive = false
							}
						}
						conn, err := amqp.Dial("amqp://admin:test@10.6.43.41:5672/")
						failOnError(err, "Failed to connect to RabbitMQ")
						defer conn.Close()

						ch, err := conn.Channel()
						failOnError(err, "Failed to open a channel")
						defer ch.Close()

						q, err := ch.QueueDeclare(
							"hello", // name
							false,   // durable
							false,   // delete when unused
							false,   // exclusive
							false,   // no-wait
							nil,     // arguments
						)
						failOnError(err, "Failed to declare a queue")

						i := in.GetPlayer()
						s := in.GetStage()
						//i_str := strconv.Itoa(int(i))

						body := "Jugador_" + i + " Ronda_" + s

						err = ch.Publish(
							"",     // exchange
							q.Name, // routing key
							false,  // mandatory
							false,  // immediate
							amqp.Publishing{
								ContentType: "text/plain",
								Body:        []byte(body),
							})
						failOnError(err, "Failed to publish a message")
						log.Printf(" ha muerdo: %d ", body)
						//log.Printf(" [x] Sent %d ", body)
					}
				} else if actualStage == "2tc" {
					log.Printf("la jugada fue realizada en 2tc")
				}
			} else {
				log.Printf("aún no comienza el nivel")
			}
			//log.Printf("Greeting: %s", r.GetStage())

		}
		return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound() + 1}, nil
	} else {
		log.Printf("ya realizaste la jugada en esta ronda")
		return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound()}, nil
	}

}

//"El jugador " + in.GetPlayer() + " hizo una jugada " + in.GetPlay() + "en la etapa" + in.GetStage()

func (s *server) AmountCheck(ctx context.Context, in *pb.AmountRequest) (*pb.AmountReply, error) {
	message := "solicito monto"
	//conexión con el pozo
	conn, err := grpc.Dial("10.6.43.43:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with pozo " + err.Error())
	}
	serviceClient := pb.NewSquidGameServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//aqui primer intento del consultar desde el servidor a otra entidad.
	r, err := serviceClient.AmountCheck(ctx, &pb.AmountRequest{Message: message})
	if err != nil {
		log.Fatalf("no se pudo solicitar el monto: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMonto())
	return &pb.AmountReply{Monto: r.GetMonto()}, nil
}

func main() {
	//códigos Etapas
	//1rv
	//2tc
	//3tn

	go func() {
		// nos convertimos en servidor (LIDER)
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

	}()

	///////////// Interfaz
	var start string
	var stage string
	var next string
	started = false
	actualStage = "1rv"
	totalPlayers = 0
	SquidGame := "none"
	//var ronda_actual int32
	//ronda_actual = 0
	for totalPlayers != 16 {
		fmt.Println("escribe start para iniciar el SquidGame: ")
		fmt.Scanln(&SquidGame)
		if totalPlayers != 16 {
			fmt.Println("no hay suficientes jugadores para comenzar el SquidGame ")
		}

	}
	if totalPlayers == 16 {
		//se da inicio al juego
		fmt.Println("escribe start para comenzar la etapa 1: ")
		fmt.Scanln(&start)
		if start == "start" {
			fmt.Println("Ha comenzado la etapa: " + actualStage)
		}
		started = true
		actualRound = 0
		for i := 0; i < 4; i++ {

			actualRound += 1
			rand.Seed(time.Now().UnixNano())
			fmt.Println("ronda " + strconv.Itoa(i+1))
			liderPlay = int(rand.Int63n(5))
			liderPlay = liderPlay + 6
			fmt.Println("jugada de lider: " + strconv.Itoa(liderPlay))
			fmt.Println("escribe cualquier letra para la siguiente ronda: ")
			fmt.Scanln(&next)
		}

		// Si tienen menos de 21 puntos, hay tabla
		for i := 0; i < 16; i++ {
			if list_of_players[i].score < 21 {
				list_of_players[i].alive = false
				fmt.Println("el jugador: " + list_of_players[i].id + " fue eliminado por no alcanzar puntaje requerido")
			}

		}

		// Hay que anunciar a los ganadores del nivel con un print o algo asi
		winnerCount := 0
		for i := 0; i < 16; i++ {
			list_of_players[i].score = 0
			if list_of_players[i].alive == true {
				winnerCount += 1
				fmt.Println("el jugador: " + list_of_players[i].id + " pasa al siguiente nivel")
			}
			//acá eliminamos al jugador sobrante.
		}
		for winnerCount%2 == 1 {
			rand.Seed(time.Now().UnixNano())
			liderPlay = int(rand.Int63n(15))
			if list_of_players[liderPlay].alive == true {
				list_of_players[liderPlay].alive = false
				winnerCount -= 1
				fmt.Println("el jugador: " + list_of_players[liderPlay].id + " es eliminado automáticamente")
			}
		}

		actualStage = "2tc"
		//esto separa a los ganadores de la ronda pasada en 2 grupos
		changer := 0
		for i := 0; i < 16; i++ {
			if list_of_players[i].alive == true {
				if changer == 0 {
					group1 = append(group1, PlayerStruct{list_of_players[i].id, true, 0})
					changer = 1
				} else {
					group2 = append(group2, PlayerStruct{list_of_players[i].id, true, 0})
					changer = 0
				}

			}
		}

		fmt.Println("escribe start para comenzar la etapa 2: ")
		fmt.Scanln(&start)
		if start == "start" {
			fmt.Println("Ha comenzado la etapa: " + actualStage)
		}
		liderPlay = int(rand.Int63n(3))
		liderPlay = liderPlay + 1
		fmt.Println("jugada de lider: " + strconv.Itoa(liderPlay))
		fmt.Println("ingresa start cuando los jugadores ya hayan realizado sus jugadas: ")
		fmt.Scanln(&start)
		scoreGroup1 := 0
		scoreGroup2 := 0
		passGroup1 := false
		passGroup2 := false
		for i := 0; i < len(group1); i++ {
			scoreGroup1 += group1[i].score
		}
		for i := 0; i < len(group2); i++ {
			scoreGroup2 += group2[i].score
		}

		if scoreGroup1%2 == liderPlay%2 {
			fmt.Println("pasa grupo 1")
			passGroup1 = true

		}
		if scoreGroup2%2 == liderPlay%2 {
			fmt.Println("pasa grupo 2")
			passGroup2 = true
		}
		if passGroup1 == true && passGroup2 == true {
			fmt.Println("ambos equipos pasan")
		} else if passGroup1 == true && passGroup2 == false {
			fmt.Println("pasa el equipo 1")
		} else if passGroup1 == false && passGroup2 == true {
			fmt.Println("pasa el equipo 2")
		} else {
			fmt.Println("aqui hay que escoger al azar uno de los 2 equipos")
		}

		fmt.Println("se ha muerto ste men: 2")
		fmt.Println("los jugadores vivos que pasan a la siguiente ronda son 16")
		fmt.Println("los ganadores de la ronda son 1,2,3 ")
		actualStage = "3tn"

		fmt.Println("escribe start para comenzar la etapa 3: ")
		fmt.Scanln(&start)
		if start == "start" {
			fmt.Println("Ha comenzado la etapa: " + stage)
		}
		liderPlay = int(rand.Int63n(9))
		liderPlay = liderPlay + 1
		fmt.Println("jugada de lider: " + strconv.Itoa(liderPlay))
		fmt.Println("se ha muerto ste men: 2")
		fmt.Println("los jugadores vivos que pasan a la siguiente ronda son 16")
		fmt.Println("los ganadores de la ronda son 1,2,3 ")
		stage = "4end"
	}

	//enviar

	// NAMENOOOOOOOOOOOOOOOOOOOOOODEEEEEEEEEEEEEEE
	/*conn, err := grpc.Dial("10.6.43.42:8080", grpc.WithInsecure())

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceLider := pb.NewSquidGameServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	playerID := "1"
	stage := "1"
	jugada := "5"

	r, err := serviceLider.SendPlays(ctx, &pb.SendRequest{Player: playerID, Play: jugada, Stage: stage})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())/*

	//////////////////////

	// POZOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO

	/*var second string
	message := "solicitar"

	fmt.Println("ingresa la letra a para solicitar monto: ")
	fmt.Scanln(&second)

	//if true {} aqui parte el POZO
	conn2, err2 := grpc.Dial("10.6.43.43:8080", grpc.WithInsecure())

	if err2 != nil {
		panic("cannot connect with pozo " + err.Error())
	}
	serviceClient := pb.NewSquidGameServiceClient(conn2)
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//aqui primer intento del consultar desde el servidor a otra entidad.
	r2, err3 := serviceClient.AmountCheck(ctx, &pb.AmountRequest{Message: message})
	if err3 != nil {
		log.Fatalf("no se pudo solicitar el monto: %v", err3)
	}
	log.Printf("Greeting: %s", r2.GetMessage())
	*/
	/////////////////////////////

}
