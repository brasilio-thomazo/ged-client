package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"br.dev.optimus/hermes/database"
	"br.dev.optimus/hermes/pb"
	"br.dev.optimus/hermes/service"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "server port")

func main() {
	flag.Parse()
	writerDBHost := os.Getenv("DB_WRITER_HOST")
	writerDBPort := os.Getenv("DB_WRITER_PORT")
	readerDBHost := os.Getenv("DB_READER_HOST")
	readerDBPort := os.Getenv("DB_READER_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	serverPort := os.Getenv("GRPC_PORT")
	writerDSN := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable", writerDBHost, writerDBPort, username, password, dbname)
	readerDSN := fmt.Sprintf("host=%s port=%s user=%s password='%s' dbname=%s sslmode=disable", readerDBHost, readerDBPort, username, password, dbname)

	reader, err := database.NewReader(readerDSN)
	if err != nil {
		log.Fatalf("connect to database reader server error %v", err)
	}

	writer, err := database.NewWriter(writerDSN)
	if err != nil {
		log.Fatalf("connect to database writer server error %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		log.Fatalf("listen %d error [%v]", *port, err)
	}

	s := grpc.NewServer()
	srv := service.NewHermesService(reader, writer)
	pb.RegisterHermesServer(s, srv)

	log.Printf("server listening %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("server start falied %v", err)
	}
}
