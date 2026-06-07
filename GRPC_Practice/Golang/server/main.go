package main

import (
	"io"
	"log"
	"net"

	pb "pingpong/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGameServiceServer
}

func (s *server) PlayPingPong(stream pb.GameService_PlayPingPongServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		reply := "Pong"
		if in.Text == "Pong" {
			reply = "Ping"
		}

		log.Printf("%s → %s", in.Text, reply)

		if err := stream.Send(&pb.Message{Text: reply}); err != nil {
			return err
		}
	}
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, &server{})
	log.Println("Listening on :50051")
	s.Serve(lis)
}
