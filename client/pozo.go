package main

import (
	//"context"
	//"fmt"
	//"io/ioutil"
	//"math/rand"
	//"strconv"
	//"time"
	"log"
	//"net"
	
	amqp "github.com/rabbitmq/amqp091-go"

	//"fmt"
	//pb "github.com/fabiusinfo/SquidGame/proto"
	//"google.golang.org/grpc"

	
)

/*type server struct {
	pb.UnimplementedSquidGameServiceServer
}*/

/*func (s *server) AmountCheck(ctx context.Context, in *pb.AmountRequest) (*pb.AmountReply, error) {
	monto := "5000"
	return &pb.AmountReply{Monto: monto}, nil
}*/

// RabbitMQ dale recibe

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	go func(){
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

	conn, err := amqp.Dial("amqp://admin:test@localhost:5672/")
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
	// Resgistrar registro de muertes, registarr registro aweonao tonto culiao te veo te mato no weon noOOOOO
	var path = "./client/registro_de_muertes.txt"
	b, errtxt := ioutil.ReadFile(path)

	if errtxt != nil {
		log.Fatal(errtxt)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			cadena := string(d.Body)
			b = append(b, []byte(cadena+" \n")...)
			errtxt = ioutil.WriteFile(path, b, 0644)
			if errtxt != nil {
				log.Fatal(errtxt)
			}

			fmt.Println("Alguien murio")

		}
		
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
