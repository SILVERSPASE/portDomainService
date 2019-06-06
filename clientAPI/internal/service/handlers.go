package api

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	proto "github.com/silverspase/portService/clientAPI/internal/api/v1"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (s *APIServer) GetAPort(w http.ResponseWriter, r *http.Request) {
	portID := chi.URLParam(r, "portID")
	if portID == "" {
		render.JSON(w, r, "specify port id")
		return
	}
	port, err := s.Service.Read(context.TODO(), &proto.ReadRequest{Api: s.API, PortID: portID})
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, port)
	return
}

func (s *APIServer) DeletePort(w http.ResponseWriter, r *http.Request) {
	portID := chi.URLParam(r, "portID")
	id, err := s.Service.Delete(context.TODO(), &proto.DeleteRequest{Api: s.API, PortID: portID})
	if err != nil {
		log.Fatal(err)
	}
	render.JSON(w, r, id)
}


func (s *APIServer) LoadFromJSON(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	// ParseMultipartForm will keep fileMaxSize in memory and rest in disk trying to be not greedy by memory.
	if err := r.ParseMultipartForm(1*1024*1024); err != nil {
		render.JSON(w, r, err)
	}
	defer r.Body.Close()

	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		render.JSON(w, r, err)
	}

	var ports map[string]*proto.Port

	err = json.Unmarshal(buf.Bytes(), &ports)
	if err != nil {
		log.Fatal(err)
	}

	res, err := s.Service.LoadFromJSON(context.TODO(), &proto.LoadFromJSONRequest{
		Api:   s.API,
		Ports: ports,
	})
	if err != nil {
		log.Println(err)
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
