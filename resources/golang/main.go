package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"br.dev.optimus/duat/database"
	"br.dev.optimus/duat/pb"
	"br.dev.optimus/duat/service"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "server port")

func main() {
	flag.Parse()
	host := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	serverPort := os.Getenv("SERVER_PORT")
	if len(serverPort) < 5 {
		serverPort = fmt.Sprintf("%d", *port)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable", host, dbport, username, password, dbname)

	reader, err := database.NewReader(dsn)
	if err != nil {
		log.Fatalf("connect to database reader server error %v", err)
	}

	writer, err := database.NewWriter(dsn)
	if err != nil {
		log.Fatalf("connect to database writer server error %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		log.Fatalf("listen %d error [%v]", *port, err)
	}

	s := grpc.NewServer()
	serviceDocument := service.NewDocumentService(reader, writer)
	serviceDocumetImage := service.NewDocumentImageService(reader, writer)
	pb.RegisterDocumentServiceServer(s, serviceDocument)
	pb.RegisterDocumentImageServiceServer(s, serviceDocumetImage)

	log.Printf("server listening %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server start falied %v", err)
	}
}
