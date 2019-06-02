package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	proto "github.com/silverspase/portService/clientAPI/internal/api/v1"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (s *APIServer) GetAPort(w http.ResponseWriter, r *http.Request) {
	portID := chi.URLParam(r, "portID")
	port, err := s.Service.Read(context.TODO(), &proto.ReadRequest{Api: s.API, PortID: portID})
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, port)
}

func (s *APIServer) DeletePort(w http.ResponseWriter, r *http.Request) {
	portID := chi.URLParam(r, "portID")
	id, err := s.Service.Delete(context.TODO(), &proto.DeleteRequest{Api: s.API, PortID: portID})
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, id)
}

func (s *APIServer) AddPort(w http.ResponseWriter, r *http.Request) { //TODO get port from request
	port := &proto.Port{
		PortID: "test",
		Name:   "test",
	}
	id, err := s.Service.Create(context.TODO(), &proto.CreateRequest{Api: s.API, Port: port})
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, id)
}

func (s *APIServer) LoadFromJSON(w http.ResponseWriter, r *http.Request) {
	var portsArr []*proto.Port

	input, err := ioutil.ReadFile("/Users/ashch/go/src/github.com/silverspase/portDomainService/clientAPI/server/ports.json")
	if err != nil {
		log.Fatal(err)
	}
	var ports map[string]proto.Port

	err = json.Unmarshal([]byte(input), &ports)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ports)

	for k, val := range ports {
		val.PortID = k
		portsArr = append(portsArr, &val)
	}

	res, err := s.Service.LoadFromJSON(context.TODO(), &proto.LoadFromJSONRequest{
		Api:   s.API,
		Ports: portsArr,
	})
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, res)

}

func (s *APIServer) GetAllPorts(w http.ResponseWriter, r *http.Request) {
	ports, err := s.Service.ReadAll(context.TODO(), &proto.ReadAllRequest{Api: s.API})
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, ports)
}
