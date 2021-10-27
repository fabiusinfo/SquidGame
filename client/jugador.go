package main

import (
	"context"
	"fmt"
	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"

	//"math/rand"
	"strconv"
	"log"
	"time"
)

func main() {
	var action string
	playerNumber := "1"
	play := "2"
	actualStage:="none"
	codes1 := "none"
	codes2 := "none"
	codes3 := "none"
	alive := true
	started:=false
	var playerCodes[16]string
	//inscribimos los bots
	for i:=0 ; i<16 ; i++ {
		conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

		if err != nil {
			panic("cannot connect with server " + err.Error())
		}

		servicePlayer := pb.NewSquidGameServiceClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := servicePlayer.JoinGame(ctx, &pb.JoinRequest{Player: i})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("inscrito")
		//signed=r.GetSigned()
		playerCodes[i]=r.GetCodes1()
		//codes1 = r.GetCodes1()
		//codes2 = r.GetCodes2()
		//codes3 = r.GetCodes3()
		//actualStage=codes1
		fmt.Println("inscripción al SquidGame realizada con éxito.")
	}
	fmt.Println(codes1+codes2+codes3)
	for alive {
		fmt.Println("ID del jugador: " + playerNumber + " , Jugada: " + play + " , etapa: " + actualStage)
		fmt.Println("Activar jugador, join->unirse, send->enviar jugadas, amount->solicitar monto: ")
		fmt.Scanln(&action)
		if action == "send"{
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

		switch action{
			// unirse al juego del calamar
		case "join":
			if actualStage == "none" {
				r, err := servicePlayer.JoinGame(ctx, &pb.JoinRequest{Player: playerNumber})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				log.Printf("inscrito")
				//signed=r.GetSigned()
				codes1 = r.GetCodes1()
				codes2 = r.GetCodes2()
				codes3 = r.GetCodes3()
				actualStage=codes1
				fmt.Println("inscripción al SquidGame realizada con éxito.")
			} else {
				fmt.Println("ya estas inscrito.")
			}
			//enviar jugada realizada
		case "send":
			if actualStage != "none" && started == true{
				
				play, err2 := strconv.Atoi(play)
				playsend:=int32(play)
				r, err := servicePlayer.SendPlays(ctx, &pb.SendRequest{Player: playerNumber, Play: playsend, Stage: actualStage})
				if err != nil {
					log.Fatalf("fallo 1: %v", err)
				}
				if err2 != nil {
					log.Fatalf("fallo 2: %v", err2)
				}
				//log.Printf("Greeting: %s", r.GetMessage())
				actualStage=r.GetStage()
				alive = r.GetAlive()
				//started = r.GetStarted()
		} else{
			fmt.Println("todavía no comienza el SquidGame o aún no estás inscrito.")
		}

			//solicitar monto
		case "amount":
			message:= "solicito monto"
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

}
