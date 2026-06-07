package main

import (
	"context"
	"log"
	"time"

	pb "pingpong/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	stream, _ := pb.NewGameServiceClient(conn).PlayPingPong(context.Background())

	// Wait 2 sec then Go sends the first Ping to Node server
	time.Sleep(2 * time.Second)
	stream.Send(&pb.Message{Text: "Ping"})

	for {
		in, err := stream.Recv()
		if err != nil {
			return
		}

		reply := "Pong"
		if in.Text == "Pong" {
			reply = "Ping"
		}

		log.Printf("%s → %s", in.Text, reply)
		time.Sleep(2 * time.Second)
		stream.Send(&pb.Message{Text: reply})
	}
}
