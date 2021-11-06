package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	//"math/rand"
	//"strconv"
	//"time"
	"log"
	"net"

	amqp "github.com/rabbitmq/amqp091-go"

	//"fmt"
	pb "github.com/fabiusinfo/SquidGame/proto"
	"google.golang.org/grpc"
)

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

type server struct {
	pb.UnimplementedSquidGameServiceServer
}

func (s *server) AmountCheck(ctx context.Context, in *pb.AmountRequest) (*pb.AmountReply, error) {
	monto := strconv.Itoa(monto_actual)
	return &pb.AmountReply{Monto: monto}, nil
}

// RabbitMQ dale recibe

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

var monto_actual int = 0

func main() {

	//Se elimina y se vuelve a crear registro_de_muertes al iniciar el pozo
	if delet == 1 {
		delet = 0
		nombreArchivo := "registro_de_muertes.txt" // El nombre o ruta absoluta del archivo
		err := os.Remove(nombreArchivo)
		if err != nil {
			fmt.Printf("\n")
		} else {
			fmt.Println("Se ha creado registro_de_muertes.txt")
		}
		crearArchivo(nombreArchivo)
	}

	go func() {
		listner, err := net.Listen("tcp", ":8080")
		//conn, err := grpc.Dial("10.6.43.41:8080", grpc.WithInsecure())

		if err != nil {
			panic("cannot connect with server " + err.Error())
		}

		serv := grpc.NewServer()
		pb.RegisterSquidGameServiceServer(serv, &server{})
		if err = serv.Serve(listner); err != nil {
			panic("cannot initialize the server" + err.Error())

		}
	}()

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

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)
	var path = "./Pozo/registro_de_muertes.txt"
	b, errtxt := ioutil.ReadFile(path)

	if errtxt != nil {
		log.Fatal(errtxt)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			cadena := string(d.Body)
			monto_actual = monto_actual + 100000000
			monto_str := strconv.Itoa(monto_actual)
			b = append(b, []byte(cadena+" "+monto_str+" \n")...)
			errtxt = ioutil.WriteFile(path, b, 0644)
			if errtxt != nil {
				log.Fatal(errtxt)
			}

		}

	}()

	log.Printf(" [*] Pozo esperando mensajes. Para salir -CTRL+C-")
	<-forever

}
