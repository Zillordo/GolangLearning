package main

import (
	".."
	"../../../database"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	op := flag.String("op", "s", "s for server, and c for client ")
	flag.Parse()
	switch strings.ToLower(*op) {
	case "s":
		runGRPCServer()
	case "c":
		runGRPCClient()
	}
}

func runGRPCServer() {
	grpclog.Infoln("Starting GRPC Server")
	lis, err := net.Listen("tcp", ":8282")
	grpclog.Infoln("Listening onf 127:0.0.1:8282")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	dinoServer, err := dinogrpc.NewDinoGrpcServer(database.MONGODB, "mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	dinogrpc.RegisterDinoServiceServer(grpcServer, dinoServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Faild to serve: %s", err)
	}
}

func runGRPCClient() {
	conn, err := grpc.Dial("127.0.0.1:8282", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := dinogrpc.NewDinoServiceClient(conn)
	input := ""
	fmt.Println("All animals?(y/n)")
	_, _ = fmt.Scanln(&input)
	if strings.ToLower(input) == "y" {
		animals, err := client.GetAllAnimals(context.Background(), &dinogrpc.Request{})
		if err != nil {
			log.Fatal(err)
		}

		for true {
			animal, err := animals.Recv()
			if err != io.EOF {
				break
			}
			if err != nil {
				grpclog.Fatal(err)
			}
			grpclog.Infoln(animal)
		}
		return
	}
	fmt.Println("Nickname?")
	_, _ = fmt.Scanln(&input)
	a, err := client.GetAnimal(context.Background(), &dinogrpc.Request{Nickname: input})
	if err != nil {
		log.Fatal(err)
	}
	grpclog.Infoln(a)
}
