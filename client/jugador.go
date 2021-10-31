package main

import (
	"context"
	"fmt"

	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"

	"math/rand"
	"log"
	"strconv"
	"time"
)

type PlayerStruct struct {
	id    string
	alive bool
	round int32
	score int32
}

func main() {
	var action string
	playerNumber := "1"
	play := "2"
	actualStage := "none"
	codes1 := "none"
	codes2 := "none"
	codes3 := "none"
	alive := true
	started := false
	flag1 := false
	//	var playersAlive [16]bool

	var list_of_players []PlayerStruct


	//inscripción
	for !flag1 {
		fmt.Println("escribe join para inscribirse en el SquidGame: ")
		fmt.Scanln(&action)
		if action == "join" {
			flag1 = true
		}
	}

	if actualStage == "none" {
		conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())
		if err != nil {
			panic("cannot connect with server " + err.Error())
		}
		servicePlayer := pb.NewSquidGameServiceClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := servicePlayer.JoinGame(ctx, &pb.JoinRequest{Player: playerNumber})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		list_of_players = append(list_of_players, PlayerStruct{playerNumber, true, 1, 0})
		log.Printf("inscrito")
		//signed=r.GetSigned()
		codes1 = r.GetCodes1()
		codes2 = r.GetCodes2()
		codes3 = r.GetCodes3()
		actualStage = codes1
		started = true
		if started == true {
			fmt.Println("inscripción al SquidGame realizada con éxito.")
		}
	} else {
		fmt.Println("ya estas inscrito.")
	}

	//inscribimos los bots
	
	for i := 0; i < 15; i++ {

		list_of_players = append(list_of_players, PlayerStruct{strconv.Itoa(i + 2), true, 1, 0})

		conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

		if err != nil {
			panic("cannot connect with server " + err.Error())
		}

		servicePlayer := pb.NewSquidGameServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := servicePlayer.JoinGame(ctx, &pb.JoinRequest{Player: list_of_players[i].id})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("inscrito")
		//signed=r.GetSigned()
		list_of_players[i].id = strconv.Itoa(i + 2)
		list_of_players[i].alive = true
		//codes1 = r.GetCodes1()
		codes2 = r.GetCodes2()
		//codes3 = r.GetCodes3()
		//actualStage=codes1
		fmt.Println("inscripción al SquidGame realizada con éxito.")
	}
	fmt.Println(codes1 + codes2 + codes3)
	//Aquí finaliza la inscripción

	//Aquí realizar jugada o checkAmount nivel 1 ¿?
	for alive {
		flag1 = false
		for !flag1 {
			fmt.Println("escribe send -> enviar jugada, check -> solicitar monto: ")
			fmt.Scanln(&action)
			if action == "send" {
				fmt.Println("escribe un número del 1 al 10: ")
				fmt.Scanln(&play)

				conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

				if err != nil {
					panic("cannot connect with server " + err.Error())
				}

				servicePlayer := pb.NewSquidGameServiceClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: playerNumber, Play: play, Stage: actualStage, Round: list_of_players[0].round, Score: list_of_players[0].score})
				if err != nil {
					log.Fatalf("fallo 1: %v", err)
				}
				/*
					if err2 != nil {
						log.Fatalf("fallo 2: %v", err2)
					} */

				//log.Printf("Greeting: %s", r.GetMessage())
				actualStage = r.GetStage()
				list_of_players[0].round = r.GetRound()
				list_of_players[0].alive = r.GetAlive() // el jugador debe estar en la posicion 15 de la lista
				play_int, err32 := strconv.Atoi(play)
				if err32 != nil {
					log.Fatalf("fallo 32: %v", err32)
				}
				list_of_players[0].score = list_of_players[0].score + int32(play_int)
				//started = r.GetStarted()

			} else if action == "check" {
				message := "solicito monto"

				conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

				if err != nil {
					panic("cannot connect with server " + err.Error())
				}

				servicePlayer := pb.NewSquidGameServiceClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()
				r, err := servicePlayer.AmountCheck(ctx, &pb.AmountRequest{Message: message})
				if err != nil {
					log.Fatalf("no se pudo solicitar el monto: %v", err)
				}
				log.Printf("Greeting: %s", r.GetMonto())

			} else {
				fmt.Println("ingresaste mal el comando")
			}

				// sección bots

		//Este pedazo de código es para las jugadas de los bots
		for i := 1; i < 16; i++ {
			if list_of_players[i].alive == true {
				fmt.Println(strconv.Itoa(i))
				botPlayer := list_of_players[i].id
				conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

				if err != nil {
					panic("cannot connect with server " + err.Error())
				}

				servicePlayer := pb.NewSquidGameServiceClient(conn)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				defer cancel()

				//play, err2 := strconv.Atoi(play)
				//playsend:=int32(play)
				//jugada aleatoria
				rand.Seed(time.Now().UnixNano())
				playsend := rand.Int63n(10) + 1
				playsend_str := strconv.Itoa(int(playsend))
				r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: botPlayer, Play: playsend_str, Stage: actualStage, Round:list_of_players[i].round})
				if err != nil {
					log.Fatalf("fallo 1: %v", err)
				}
				//if err2 != nil {
				//	log.Fatalf("fallo 2: %v", err2)
				//}
				//log.Printf("Greeting: %s", r.GetMessage())
				actualStage = r.GetStage()
				list_of_players[i].round = r.GetRound()
				list_of_players[i].alive = r.GetAlive()
				//started = r.GetStarted()
			}

		}
		}
	

	}
	fmt.Println("me muero (explota)")

}

//Aquí realizar jugada o checkAmount nivel 2 ¿?

//Aquí realizar jugada o checkAmount nivel 3 ¿?

/*
	for alive {
		//fmt.Println("ID del jugador: " + playerNumber + " , Jugada: " + play + " , etapa: " + actualStage)
		fmt.Println("Activar jugador, join->unirse, send->enviar jugadas, amount->solicitar monto: ")
		fmt.Scanln(&action)
		if action == "send" {
			fmt.Println("realiza jugada: ")
			fmt.Scanln(&play)
		}

		conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

		if err != nil {
			panic("cannot connect with server " + err.Error())
		}

		servicePlayer := pb.NewSquidGameServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		switch action {
		// unirse al juego del calamar
		case "join":
			if actualStage == "none" {
				r, err := servicePlayer.JoinGame(ctx, &pb.JoinRequest{Player: playerNumber})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				list_of_players = append(list_of_players, PlayerStruct{playerNumber, true, 1})
				log.Printf("inscrito")
				//signed=r.GetSigned()
				codes1 = r.GetCodes1()
				codes2 = r.GetCodes2()
				codes3 = r.GetCodes3()
				actualStage = codes1
				started = true
				fmt.Println("inscripción al SquidGame realizada con éxito.")
			} else {
				fmt.Println("ya estas inscrito.")
			}
			//enviar jugada realizada
		case "send":
			if actualStage != "none" && started == true {


				r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: playerNumber, Play: play, Stage: actualStage, Round:list_of_players[15].round})
				if err != nil {
					log.Fatalf("fallo 1: %v", err)
				}
				/*
					if err2 != nil {
						log.Fatalf("fallo 2: %v", err2)
					} */

//log.Printf("Greeting: %s", r.GetMessage()) */

/*
				actualStage = r.GetStage()
				list_of_players[15].round = r.GetRound()
				list_of_players[15].alive = r.GetAlive() // el jugador debe estar en la posicion 15 de la lista
				//started = r.GetStarted()

				//Este pedazo de código es para las jugadas de los bots

				for i := 0; i < 15; i++ {
					if list_of_players[i].alive == true {
						fmt.Println(strconv.Itoa(i))
						botPlayer := list_of_players[i].id
						conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

						if err != nil {
							panic("cannot connect with server " + err.Error())
						}

						servicePlayer := pb.NewSquidGameServiceClient(conn)

						ctx, cancel := context.WithTimeout(context.Background(), time.Second)
						defer cancel()

						//play, err2 := strconv.Atoi(play)
						//playsend:=int32(play)
						//jugada aleatoria
						rand.Seed(time.Now().UnixNano())
						playsend := rand.Int63n(10) + 1
						playsend_str := strconv.Itoa(int(playsend))
						r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: botPlayer, Play: playsend_str, Stage: actualStage, Round:list_of_players[i].round})
						if err != nil {
							log.Fatalf("fallo 1: %v", err)
						}
						//if err2 != nil {
						//	log.Fatalf("fallo 2: %v", err2)
						//}
						//log.Printf("Greeting: %s", r.GetMessage())
						actualStage = r.GetStage()
						list_of_players[i].round = r.GetRound()
						list_of_players[i].alive = r.GetAlive()
						//started = r.GetStarted()
					}

				}

			} else {
				fmt.Println("todavía no comienza el SquidGame o aún no estás inscrito.")
			}

			//solicitar monto
		case "amount":
			message := "solicito monto"
			r, err := servicePlayer.AmountCheck(ctx, &pb.AmountRequest{Message: message})
			if err != nil {
				log.Fatalf("no se pudo solicitar el monto: %v", err)
			}
			log.Printf("Greeting: %s", r.GetMonto())

		default:
			fmt.Println("ingresaste un mal comando.")

		}

	}
	if alive == false {

	}
	fmt.Println("me muero (explota)")

} */
