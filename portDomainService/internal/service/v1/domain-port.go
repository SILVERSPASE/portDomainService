package v1

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	proto "github.com/silverspase/portService/portDomainService/internal/api/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion          = "v1"
	noDocumentsInResult = "mongo: no documents in result"
)

// portServiceServer is implementation of proto.PortServiceServer proto interface
type portServiceServer struct {
	db *mongo.Database
}

// NewPortServiceServer creates proto.PortServiceServer
func NewPortServiceServer(db *mongo.Database) proto.PortServiceServer {
	return &portServiceServer{db: db}
}

// Create new port
func (s *portServiceServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	collection := s.db.Collection("ports")

	opts := options.UpdateOptions{}
	opts.Upsert = newTrue()

	res, err := collection.UpdateOne(ctx, bson.M{"portid": req.Port.PortID}, bson.M{"$set": req.Port}, &opts)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return &proto.CreateResponse{
		Api:    apiVersion,
		PortID: req.Port.PortID,
	}, nil
}

// Read all ports
func (s *portServiceServer) Read(ctx context.Context, req *proto.ReadRequest) (*proto.ReadResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	collection := s.db.Collection("ports")
	var port proto.Port
	err := collection.FindOne(ctx, bson.M{"portid": req.PortID}).Decode(&port)
	if err != nil && err.Error() != noDocumentsInResult {
		return nil, err
	}
	return &proto.ReadResponse{
		Api:  apiVersion,
		Port: &port,
	}, nil
}

// Update port
func (s *portServiceServer) LoadFromJSON(ctx context.Context, req *proto.LoadFromJSONRequest) (*proto.LoadFromJSONResponse, error) {
	//collection := s.db.Collection("ports")
	//opts := options.UpdateOptions{Upsert: newTrue()}
	//_, err := collection.UpdateMany(ctx, bson.D{},  bson.M{"$set": req.Ports}, &opts)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//return &proto.LoadFromJSONResponse{
	//	Api:         apiVersion,
	//	LoadedCount: "12",
	//}, nil

	var records, created, updated, skipped int64
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	collection := s.db.Collection("ports")

	opts := options.UpdateOptions{}
	opts.Upsert = newTrue()



	for key, val := range req.Ports {
		onePort := *val
		onePort.PortID = key
		res, err := collection.UpdateOne(ctx, bson.M{"portid": onePort.PortID}, bson.M{"$set": onePort}, &opts)
		if err != nil {
			return nil, err
		}
		records++
		created += res.UpsertedCount
		updated += res.ModifiedCount
		if res.UpsertedCount + res.ModifiedCount == 0 {
			skipped += res.MatchedCount
		}

	}

	return &proto.LoadFromJSONResponse{
		Api:    apiVersion,
		Records: records,
		Created: created,
		Updated: updated,
		Skipped: skipped,
	}, nil

}

// Delete port
func (s *portServiceServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	// check if the API version requested by client is supported by server
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	collection := s.db.Collection("ports")
	var v interface{}
	err := collection.FindOneAndDelete(ctx, bson.M{"portid": req.PortID}).Decode(&v)
	if err != nil {
		if err.Error() == noDocumentsInResult {
			return &proto.DeleteResponse{
				Api:    apiVersion,
				PortID: noDocumentsInResult,
			}, nil
		} else {
			return nil, err
		}
	}
	log.Println("Deleted a single document")
	return &proto.DeleteResponse{
		Api:    apiVersion,
		PortID: req.PortID,
	}, nil
}

// Read all ports
func (s *portServiceServer) ReadAll(ctx context.Context, req *proto.ReadAllRequest) (*proto.ReadAllResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	collection := s.db.Collection("ports")

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var ports []*proto.Port
	for cur.Next(ctx) {
		var elem proto.Port
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		ports = append(ports, &elem)
	}
	return &proto.ReadAllResponse{
		Api:   apiVersion,
		Ports: ports,
	}, nil
}

// checkAPI checks if the API version requested by client is supported by server
func (s *portServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func newTrue() *bool {
	b := true
	return &b
}
