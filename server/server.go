package main

import (
	"log"
	"net"
	pb "github.com/Mubinabd/RestaurantService/genproto"
	"github.com/Mubinabd/RestaurantService/service"
	"github.com/Mubinabd/RestaurantService/storage/postgres"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "github.com/lib/pq"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Println("error while connecting to postgres: ", err)
	}

	liss, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRestaurantServiceServer(s, service.NewRestaurantService(db))

	reflection.Register(s)
	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}