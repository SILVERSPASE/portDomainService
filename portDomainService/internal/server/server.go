package server

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/silverspase/portService/portDomainService/internal/protocol/grpc"
	proto "github.com/silverspase/portService/portDomainService/internal/service/v1"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

// Config is configuration for Server
type Config struct {
	GRPCPort string
	mongoURI string
	dbName   string
}

// RunServer runs gRPC server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.mongoURI, "mongoURI", "", "mongodb URI")
	flag.StringVar(&cfg.dbName, "dbName", "", "mongodb URI")

	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(cfg.dbName)

	v1API := proto.NewPortServiceServer(db)
	return grpc.RunServer(ctx, v1API, cfg.GRPCPort)
}

