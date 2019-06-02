package main

import (
	"flag"
	"log"
	"net/http"

	v1 "github.com/silverspase/portService/clientAPI/internal/api/v1"
	api "github.com/silverspase/portService/clientAPI/internal/service"
	"google.golang.org/grpc"
)

func main() {

	address := flag.String("server", "", "gRPC server in format host:port")
	port := flag.String("port", "", "http server port in format :port")
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	portService := v1.NewPortServiceClient(conn)

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	s := api.NewAPIServer(portService)

	log.Printf("Starting server")
	log.Fatal(http.ListenAndServe(*port, s.Router))
}
