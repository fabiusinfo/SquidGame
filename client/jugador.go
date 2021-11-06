package main

import (
	"context"
	"fmt"

	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"

	"log"
	"math/rand"
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
	next := "none"
	started := false
	flag1 := false

	var list_of_players []PlayerStruct

	//inscripción
	for !flag1 {
		fmt.Println("Escribe -join- para inscribirse en el SquidGame ")
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
		log.Printf(r.GetMessage())
		actualStage = "1rv"
		started = true
	} else {
		fmt.Println("Ya estas inscrito")
	}

	//inscribimos los bots

	for i := 1; i < 16; i++ {

		list_of_players = append(list_of_players, PlayerStruct{strconv.Itoa(i + 1), true, 1, 0})

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
		log.Printf(r.GetMessage())

		list_of_players[i].id = strconv.Itoa(i + 1)
		list_of_players[i].alive = true
	}
	//Aquí finaliza la inscripción

	started = false
	flag1 = false
	for !flag1 {
		fmt.Println("Ingresa -next- para comenzar el nivel 1: Rojo y Verde")
		fmt.Scanln(&next)
		if next == "next" {
			conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

			if err != nil {
				panic("cannot connect with server " + err.Error())
			}

			servicePlayer := pb.NewSquidGameServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err := servicePlayer.Started(ctx, &pb.StartRequest{Message: "solicito ingresar al nivel 1"})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			started = r.GetStarted()
			if started == true {
				flag1 = true
				break
			}

		} else {
			fmt.Println("Ingresaste mal el comando")
		}
	}

	//Aquí realizar jugada o checkAmount nivel 1
	contStage := 1
	actualStage = "1rv"
	LiderRound := 0
	flag1 = false
	for !flag1 {
		if contStage == 5 {
			flag1 = true
			break
		}

		fmt.Println("Rojo y Verde: -send- Enviar jugada / -check- Solicitar monto")
		fmt.Scanln(&action)

		if action == "send" {

			conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

			if err != nil {
				panic("cannot connect with server " + err.Error())
			}

			servicePlayer := pb.NewSquidGameServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := servicePlayer.AskRound(ctx, &pb.AskRequest{Message: "porfa dime la ronda del Lider"})
			if err != nil {
				log.Fatalf("fallo 1: %v", err)
			}
			LiderRound = int(r.GetRound())

			if LiderRound == contStage {

				if list_of_players[0].alive == true {
					if list_of_players[0].score < 21 {
						fmt.Println("Escribe un número del 1 al 10: ")
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
						actualStage = r.GetStage()
						list_of_players[0].round = r.GetRound()

						list_of_players[0].alive = r.GetAlive() // el jugador debe estar en la posicion 0 de la lista
						play_int, err32 := strconv.Atoi(play)
						if err32 != nil {
							log.Fatalf("fallo 32: %v", err32)
						}
						list_of_players[0].score = list_of_players[0].score + int32(play_int)

					} else {
						fmt.Println("Lograste sumar 21, estas salvado!")
					}
				} else {
					fmt.Println("El jugador muere")
				}
				// sección bots

				//Este pedazo de código es para las jugadas de los bots
				for i := 1; i < 16; i++ {

					if list_of_players[i].alive == true {
						if list_of_players[i].score < 21 {
							//fmt.Println(strconv.Itoa(i))
							botPlayer := list_of_players[i].id
							conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

							if err != nil {
								panic("cannot connect with server " + err.Error())
							}

							servicePlayer := pb.NewSquidGameServiceClient(conn)

							ctx, cancel := context.WithTimeout(context.Background(), time.Second)
							defer cancel()
							//jugada aleatoria
							rand.Seed(time.Now().UnixNano())
							playsend := rand.Int63n(10) + 1
							playsend_str := strconv.Itoa(int(playsend))
							r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: botPlayer, Play: playsend_str, Stage: actualStage, Round: list_of_players[i].round})
							if err != nil {
								log.Fatalf("fallo 1: %v", err)
							}
							actualStage = r.GetStage()
							list_of_players[i].round = r.GetRound()
							list_of_players[i].alive = r.GetAlive()
							list_of_players[i].score = list_of_players[i].score + int32(playsend)

						} else {
							fmt.Println("Lograste sumar 21, estas salvado!")
						}
					}

				}
				contStage += 1
			} else {
				fmt.Println("El lider todavía no inicia la siguiente ronda")
			}

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
				log.Fatalf("No se pudo solicitar el monto: %v", err)
			}
			log.Printf("Monto actual: %s", r.GetMonto())

		} else {
			fmt.Println("Ingresaste mal el comando o el lider todavía no comienza la ronda")
		}

	}

	started = false
	flag1 = false
	for !flag1 {
		fmt.Println("Ingresa -next- para comenzar el juego 2: Tirar la Cuerda")
		fmt.Scanln(&next)
		if next == "next" {
			conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

			if err != nil {
				panic("cannot connect with server " + err.Error())
			}

			servicePlayer := pb.NewSquidGameServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err := servicePlayer.Started(ctx, &pb.StartRequest{Message: "solicito ingresar al nivel 2"})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			started = r.GetStarted()
			if started == true {
				flag1 = true
				break
			}

		} else {
			fmt.Println("Ingresaste mal el comando")
		}
	}

	for i := 0; i < 16; i++ {
		list_of_players[i].score = 0
		list_of_players[i].round = 1
		if list_of_players[i].alive == true {
			conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

			if err != nil {
				panic("cannot connect with server " + err.Error())
			}

			servicePlayer := pb.NewSquidGameServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := servicePlayer.DeadOrAlive(ctx, &pb.DeadRequest{Player: list_of_players[i].id, Stage: actualStage})
			if err != nil {
				log.Fatalf("fallo 1: %v", err)
			}
			list_of_players[i].alive = r.GetDead()
		}
	}

	//Aquí realizar jugada o checkAmount nivel 2
	contStage = 1
	actualStage = "2tc"
	flag1 = false
	for !flag1 {
		if 1 >= len(list_of_players) {
			flag1 = false
			break
		}
		if contStage == 2 {
			flag1 = true
			break
		}
		fmt.Println("Tirar la Cuerda: -send- Enviar jugada / -check- Solicitar monto")
		fmt.Scanln(&action)
		if action == "send" {
			contStage += 1
			if list_of_players[0].alive == true {

				fmt.Println("Escribe un número del 1 al 4: ")
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
				actualStage = r.GetStage()
				list_of_players[0].round = r.GetRound()
				list_of_players[0].alive = r.GetAlive() // el jugador debe estar en la posicion 0 de la lista
				play_int, err32 := strconv.Atoi(play)
				if err32 != nil {
					log.Fatalf("fallo 32: %v", err32)
				}
				list_of_players[0].score = list_of_players[0].score + int32(play_int)

			} else {
				fmt.Println("El jugador muere")
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
					//jugada aleatoria
					rand.Seed(time.Now().UnixNano())
					playsend := rand.Int63n(3) + 1
					playsend_str := strconv.Itoa(int(playsend))
					r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: botPlayer, Play: playsend_str, Stage: actualStage, Round: list_of_players[i].round})
					if err != nil {
						log.Fatalf("fallo 1: %v", err)
					}
					actualStage = r.GetStage()
					list_of_players[i].round = r.GetRound()
					list_of_players[i].alive = r.GetAlive()
					list_of_players[i].score = list_of_players[i].score + int32(playsend)

				}
			}

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
			log.Printf("Monto actual: %s", r.GetMonto())

		} else {
			fmt.Println("Ingresaste mal el comando")
		}
	}

	started = false
	flag1 = false
	for !flag1 {
		fmt.Println("Ingresa -next- para comenzar el juego 3: Todo o Nada")
		fmt.Scanln(&next)
		if next == "next" {
			conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

			if err != nil {
				panic("cannot connect with server " + err.Error())
			}

			servicePlayer := pb.NewSquidGameServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err := servicePlayer.Started(ctx, &pb.StartRequest{Message: "solicito ingresar al nivel 3"})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			started = r.GetStarted()
			if started == true {
				flag1 = true
				break
			}

		} else {
			fmt.Println("Ingresaste mal el comando")
		}
	}

	for i := 0; i < 16; i++ {
		list_of_players[i].score = 0
		list_of_players[i].round = 1
		if list_of_players[i].alive == true {
			conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

			if err != nil {
				panic("cannot connect with server " + err.Error())
			}

			servicePlayer := pb.NewSquidGameServiceClient(conn)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := servicePlayer.DeadOrAlive(ctx, &pb.DeadRequest{Player: list_of_players[i].id, Stage: actualStage})
			if err != nil {
				log.Fatalf("fallo 1: %v", err)
			}
			list_of_players[i].alive = r.GetDead()
		}
	}

	//Aquí realizar jugada o checkAmount nivel 3
	contStage = 1
	actualStage = "3tn"
	flag1 = false
	for !flag1 {
		if 1 >= len(list_of_players) {
			flag1 = false
			break
		}

		if contStage == 2 {
			flag1 = true
			break
		}
		fmt.Println("Todo o Nada: -send- Enviar jugada / -check- Solicitar monto")
		fmt.Scanln(&action)
		if action == "send" {
			contStage += 1
			if list_of_players[0].alive == true {

				fmt.Println("Escribe un número del 1 al 10: ")
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
				actualStage = r.GetStage()
				list_of_players[0].round = r.GetRound()
				list_of_players[0].alive = r.GetAlive() // el jugador debe estar en la posicion 0 de la lista
				play_int, err32 := strconv.Atoi(play)
				if err32 != nil {
					log.Fatalf("fallo 32: %v", err32)
				}
				list_of_players[0].score = list_of_players[0].score + int32(play_int)

			} else {
				fmt.Println("El jugador muere")
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
					//jugada aleatoria
					rand.Seed(time.Now().UnixNano())
					playsend := rand.Int63n(9) + 1
					playsend_str := strconv.Itoa(int(playsend))
					r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: botPlayer, Play: playsend_str, Stage: actualStage, Round: list_of_players[i].round})
					if err != nil {
						log.Fatalf("fallo 1: %v", err)
					}
					actualStage = r.GetStage()
					list_of_players[i].round = r.GetRound()
					list_of_players[i].alive = r.GetAlive()
					list_of_players[i].score = list_of_players[i].score + int32(playsend)

				}
			}

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
				log.Fatalf("No se pudo solicitar el monto: %v", err)
			}
			log.Printf("Monto actual: %s", r.GetMonto())

		} else {
			fmt.Println("Ingresaste mal el comando")
		}
	}

	for flag1 {
		fmt.Println("Escribe -finish- para finalizar el proceso jugador ")
		fmt.Scanln(&action)
		if action == "finish" {
			flag1 = true
		} else {
			fmt.Println("Ingresaste mal el comando")
		}

	}
}
