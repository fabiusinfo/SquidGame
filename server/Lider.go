package main

import (
	"context"
	"fmt"
	"log"
	"math"
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
var flaggy bool
var list_of_players []PlayerStruct

//listas stage 2
var group1 []PlayerStruct
var group2 []PlayerStruct
var groupaux []PlayerStruct

//listas stage 3
var group3 []PlayerStruct

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
	return &pb.JoinReply{Message: "inscrito con exito"}, nil
}

func (s *server) Started(ctx context.Context, in *pb.StartRequest) (*pb.StartReply, error) {
	return &pb.StartReply{Started: started}, nil
}

func (s *server) AskRound(ctx context.Context, in *pb.AskRequest) (*pb.AskReply, error) {
	return &pb.AskReply{Round: int32(actualRound)}, nil
}

func (s *server) DeadOrAlive(ctx context.Context, in *pb.DeadRequest) (*pb.DeadReply, error) {
	alive := true
	if in.GetStage() == "1rv" {
		for i := 0; i < 16; i++ {
			if list_of_players[i].id == in.GetPlayer() {
				alive = list_of_players[i].alive
			}
		}
	} else if in.GetStage() == "2tc" {
		for i := 0; i < len(group1); i++ {
			if group1[i].id == in.GetPlayer() {
				alive = group1[i].alive
			}
		}
		for i := 0; i < len(group2); i++ {
			if group2[i].id == in.GetPlayer() {
				alive = group2[i].alive
			}
		}
	} else {
		log.Printf("para el nivel 3 no es necesario")

	}

	return &pb.DeadReply{Dead: alive}, nil
}

func (s *server) SendPlays(ctx context.Context, in *pb.SendRequest) (*pb.SendReply, error) {
	alive := true
	flaggy = true
	if actualRound != 0 {
		if in.GetRound() == actualRound {

			//envío al nameNode
			conn, err := grpc.Dial("10.6.43.42:8080", grpc.WithInsecure())

			if err != nil {
				panic("cannot connect with server " + err.Error())
			}

			serviceLider := pb.NewSquidGameServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := serviceLider.SendPlays(ctx, &pb.SendRequest{Player: in.GetPlayer(), Play: in.GetPlay(), Stage: in.GetStage(), Round: in.GetRound()})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			//Envío al Pozo

			if started == true {
				pPlay, errpPlay := strconv.Atoi(in.GetPlay())
				if errpPlay != nil {
					log.Fatalf("could not greet: %v", errpPlay)
				}
				if actualStage == "1rv" {
					for i := 0; i < 16; i++ {
						if list_of_players[i].id == in.GetPlayer() {

							list_of_players[i].score += pPlay
						}
					}
				} else if actualStage == "2tc" {
					for i := 0; i < len(group1); i++ {
						if group1[i].id == in.GetPlayer() {

							group1[i].score += pPlay
						}
					}
					for i := 0; i < len(group2); i++ {
						if group2[i].id == in.GetPlayer() {

							group2[i].score += pPlay
						}
					}
				} else {
					for i := 0; i < len(group3); i++ {
						if group3[i].id == in.GetPlayer() {

							group3[i].score += pPlay
						}
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
						log.Printf("Ha muerto: %s ", body)
					}
				}
			} else {
				log.Printf("aún no comienza el nivel")

				return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound()}, nil
			}
			return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound() + 1}, nil
		} else {
			log.Printf("ya realizaste la jugada en esta ronda ")
			return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound()}, nil
		}

	} else {
		log.Printf("el lider todavía no comienza la ronda")
		return &pb.SendReply{Stage: actualStage, Alive: alive, Round: in.GetRound()}, nil
	}

}

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
	//var stage string
	var next string
	started = false
	actualStage = "1rv"
	totalPlayers = 0
	SquidGame := "none"
	flaggy = false
	finisher := true
	flag1 := false
	Winner := "none"

	for totalPlayers != 16 {
		fmt.Println("Escribe -start- para iniciar el SquidGame: ")
		fmt.Scanln(&SquidGame)
		if totalPlayers != 16 {
			fmt.Println("No hay suficientes jugadores para comenzar el SquidGame ")
		}

	}
	if totalPlayers == 16 {
		//se da inicio al juego
		fmt.Println("Escribe -start- para comenzar la etapa 1: ")
		fmt.Scanln(&start)

		if start == "start" {
			fmt.Println("Ha comenzado la etapa " + actualStage + ": Rojo y Verde")
		}
		started = true
		actualRound = 1
		for i := 0; i < 4; i++ {

			rand.Seed(time.Now().UnixNano())
			fmt.Println("Ronda " + strconv.Itoa(i+1))
			liderPlay = int(rand.Int63n(5))
			liderPlay = liderPlay + 6
			guardian := false
			for !guardian {
				fmt.Println("Jugada de lider: " + strconv.Itoa(liderPlay))
				if i == 3 {
					fmt.Println("Escribe -next- pasar al siguiente juego: ")
				} else {
					fmt.Println("Escribe -next- para la siguiente ronda: ")
				}
				fmt.Scanln(&next)
				if next == "next" {
					if flaggy == false {
						fmt.Println("Los jugadores todavía no realizan las jugadas ")
					} else {
						guardian = true
					}
				} else {
					fmt.Println("Ingresaste mal el comando ")
				}
			}
			actualRound += 1
			flaggy = false
			if i == 3 {
				actualStage = "2tc"

			}

		}
		started = false

		// Si siguen vivos pero tienen menos de 21 puntos, son eliminados
		for j := 0; j < 16; j++ {
			if (list_of_players[j].score < 21) && (list_of_players[j].alive == true) {
				list_of_players[j].alive = false
				puntaje := strconv.Itoa(int(list_of_players[j].score))
				fmt.Println("El jugador " + list_of_players[j].id + " fue eliminado: " + puntaje + "/21 :c") // 14/21
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
				//s := in.GetStage()
				//i_str := strconv.Itoa(int(i))

				body := "Jugador_" + list_of_players[j].id + " Ronda_1rv"

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
				//log.Printf("Ha muerto: %s ", body)
				//log.Printf(" [x] Sent %d ", body)

			}

		}

		// Hay que anunciar a los ganadores del nivel
		winnerCount := 0
		for i := 0; i < 16; i++ {
			list_of_players[i].score = 0
			if list_of_players[i].alive == true {
				winnerCount += 1
				fmt.Println("El jugador: " + list_of_players[i].id + " pasa al siguiente nivel!")
				Winner = list_of_players[i].id
			}
			//acá eliminamos al azar, al jugador sobrante.
		}
		if winnerCount == 1 {
			fmt.Println("El ganador del SquidGame es: " + Winner)
			finisher = false
		}

		if finisher {
			for winnerCount%2 == 1 {
				rand.Seed(time.Now().UnixNano())
				liderPlay = int(rand.Int63n(15))
				if list_of_players[liderPlay].alive == true {
					list_of_players[liderPlay].alive = false
					winnerCount -= 1
					fmt.Println("El jugador: " + list_of_players[liderPlay].id + " es eliminado automáticamente :c")
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

					body := "Jugador_" + list_of_players[liderPlay].id + " Ronda_" + actualStage

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
					log.Printf("Ha muerto: %s ", body)

				}
				///////////////////////////////////////////////////////////////////////////////

			}
			if finisher {
				winnerCount = 0
				for i := 0; i < 16; i++ {
					list_of_players[i].score = 0
					if list_of_players[i].alive == true {
						winnerCount += 1
						//fmt.Println("El jugador: " + list_of_players[i].id + " pasa al siguiente nivel!")
						Winner = list_of_players[i].id
					}
					//acá eliminamos al azar, al jugador sobrante.
				}
				if winnerCount == 1 {
					fmt.Println("El ganador del SquidGame es: " + Winner)
					finisher = false
				}
			}

		}

		if finisher {
			actualStage = "2tc"
			//Separar a los ganadores de la ronda pasada en 2 grupos
			changer := 0
			for i := 0; i < 16; i++ {
				if list_of_players[i].alive == true {
					if changer == 0 {
						group1 = append(group1, PlayerStruct{list_of_players[i].id, true, 0})
						fmt.Println("Se agrega al grupo 1: " + list_of_players[i].id)
						changer = 1
					} else {
						group2 = append(group2, PlayerStruct{list_of_players[i].id, true, 0})
						fmt.Println("Se agrega al grupo 2: " + list_of_players[i].id)
						changer = 0
					}

				}
			}
			started = false

			for !flag1 {
				fmt.Println("Escribe -start- para comenzar la etapa 2: ")
				actualRound = 1
				fmt.Scanln(&start)
				if start == "start" {
					started = true
					flag1 = true
					fmt.Println("Ha comenzado la etapa " + actualStage + ": Tirar la Cuerda")
				} else {
					fmt.Println("Ingresaste mal el comando")
				}
			}

			rand.Seed(time.Now().UnixNano())
			liderPlay = int(rand.Int63n(3))
			liderPlay = liderPlay + 1
			fmt.Println("Jugada de lider: " + strconv.Itoa(liderPlay))
			fmt.Println("Ingresa -start- cuando los jugadores ya hayan realizado sus jugadas: ")
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
			fmt.Println("Grupo1: " + strconv.Itoa(scoreGroup1))
			fmt.Println("Grupo2: " + strconv.Itoa(scoreGroup2))
			if scoreGroup1%2 == liderPlay%2 {
				passGroup1 = true
			}
			if scoreGroup2%2 == liderPlay%2 {
				passGroup2 = true
			}
			if passGroup1 == true && passGroup2 == true {
				fmt.Println("Ambos grupos avanzan! yey")
				winnerCount = len(group2) + len(group1)
				groupaux = append(group1, group2...)
			} else if passGroup1 == true && passGroup2 == false {
				fmt.Println("Avanza el Grupo 1")
				for i := 0; i < len(group2); i++ {
					group2[i].alive = false
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

					body := "Jugador_" + group2[i].id + " Ronda_2tc"

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
					log.Printf("Ha muerto: %s ", body)
				}
				winnerCount = len(group1)
				groupaux = group1
			} else if passGroup1 == false && passGroup2 == true {
				fmt.Println("Avanza el Grupo 2")
				for i := 0; i < len(group1); i++ {
					group1[i].alive = false
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

					body := "Jugador_" + group1[i].id + " Ronda_2tc"

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
					log.Printf("Ha muerto: %s ", body)
				}
				winnerCount = len(group2)
				groupaux = group2
			} else {
				rand.Seed(time.Now().UnixNano())
				liderPlay = int(rand.Int63n(1))
				if liderPlay == 0 {
					fmt.Println("Avanza el Grupo 1 por aletoriedad")
					for i := 0; i < len(group2); i++ {
						group2[i].alive = false
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

						body := "Jugador_" + group2[i].id + " Ronda_2tc"

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
						log.Printf("Ha muerto: %s ", body)

					}
					winnerCount = len(group1)
					groupaux = group1
				} else {
					fmt.Println("Avanza el Grupo 2 por aletoriedad")
					for i := 0; i < len(group1); i++ {
						group1[i].alive = false
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

						body := "Jugador_" + group1[i].id + " Ronda_2tc"

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
						log.Printf("Ha muerto: %s ", body)
					}
					winnerCount = len(group2)
					groupaux = group2
				}

			}
		}

		//aqui esta terminando la ronda 2
		if finisher {
			winnerCount = 0
			for i := 0; i < len(groupaux); i++ {
				groupaux[i].score = 0
				if groupaux[i].alive == true {
					winnerCount += 1
					fmt.Println("El jugador: " + groupaux[i].id + " avanza a la siguiente etapa!")
					Winner = groupaux[i].id
				}
			}

			if winnerCount == 1 {
				fmt.Println("El ganador del SquidGame es: " + Winner)
				finisher = false
			}
		}
		if finisher {
			//Se elimina si son impares
			for winnerCount%2 == 1 && winnerCount != 1 {
				rand.Seed(time.Now().UnixNano())
				liderPlay = int(rand.Int63n(int64(len(groupaux))))
				if groupaux[liderPlay].alive == true {
					groupaux[liderPlay].alive = false
					winnerCount -= 1
					fmt.Println("El jugador: " + groupaux[liderPlay].id + " es eliminado automáticamente")
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
					//s := in.GetStage()
					//i_str := strconv.Itoa(int(i))

					body := "Jugador_" + groupaux[liderPlay].id + " Ronda_" + actualStage

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
					log.Printf("Ha muerto: %s ", body)
					//log.Printf(" [x] Sent %d ", body)
				}
			}
			winnerCount = 0
			for i := 0; i < len(groupaux); i++ {
				if groupaux[i].alive == true {
					winnerCount += 1
					Winner = groupaux[i].id
				}
			}
			if winnerCount == 1 {
				fmt.Println("El ganador del SquidGame es: " + Winner)
				finisher = false
			}

		}

		if finisher {
			actualStage = "3tn"
			actualRound = 1
			started = false
			//for que recorra los vivos, haga parejas y entre ellos jueguen
			for i := 0; i < len(groupaux); i++ {
				if groupaux[i].alive == true {
					group3 = append(group3, PlayerStruct{groupaux[i].id, true, 0})
				}
			}
			flag1 = false
			for !flag1 {

				fmt.Println("Escribe -start- para comenzar la etapa 3: ")
				fmt.Scanln(&start)
				if start == "start" {
					started = true
					flag1 = true
					fmt.Println("Ha comenzado la etapa " + actualStage + ": Todo o Nada")
				} else {
					fmt.Println("Ingresaste mal el comando")
				}
			}
			liderPlay = int(rand.Int63n(9))
			liderPlay = liderPlay + 1
			fmt.Println("Jugada de lider: " + strconv.Itoa(liderPlay))
			for flag1 {
				fmt.Println("Escribe -finish- para mostrar a los ganadores del SquidGame y finalizar el proceso: ")
				fmt.Scanln(&start)
				if start == "finish" {
					flag1 = false
				} else {
					fmt.Println("Ingresaste mal el comando")
				}

			}

			// Jugadas
			//aqui esta el problema uwu
			if len(group3)%2 == 1 {
				fmt.Println("El ganador del SquidGame es: " + group3[0].id)
			} else {
				for i := 0; i < len(group3); i++ {
					fmt.Println(len(group3))
					if group3[i].score == group3[i+1].score {
						fmt.Println(group3[i].id + " es un ganador del Squid Game \n")
						fmt.Println(group3[i+1].id + " es un ganador del Squid Game \n")
					} else if int(math.Abs(float64(liderPlay)-float64(group3[i].score))) == int(math.Abs(float64(liderPlay)-float64(group3[i+1].score))) { // si el calculo da el mismo resultado pues ambos ganan
						fmt.Println(group3[i].id + " es un ganador del Squid Game \n")
						fmt.Println(group3[i+1].id + " es un ganador del Squid Game \n")
					} else if int(math.Abs(float64(liderPlay)-float64(group3[i].score))) < int(math.Abs(float64(liderPlay)-float64(group3[i+1].score))) {
						fmt.Println(group3[i].id + " es un ganador del Squid Game \n")
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

						body := "Jugador_" + group3[i+1].id + " Ronda_3tn"

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
						log.Printf("Ha muerto: %s ", body)
					} else {
						fmt.Println(group3[i+1].id + " es un ganador del Squid Game \n")
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

						body := "Jugador_" + group3[i].id + " Ronda_3tn"

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
						log.Printf("Ha muerto: %s ", body)
					}

					i++
				}
			}

			actualStage = "4end"

		}
	}

}
