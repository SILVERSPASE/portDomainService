package api

import (
	"fmt"

	"google.golang.org/genproto/googleapis/type/latlng"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	proto "github.com/silverspase/portService/clientAPI/internal/api/v1"
)

const apiVer = "v1"

func (s *APIServer) Routes() {
	s.Router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,          // Log API request calls
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing server
	)
	s.Router.Get("/{portID}", s.GetAPort)
	s.Router.Delete("/{portID}", s.DeletePort)
	s.Router.Post("/json", s.LoadFromJSON)
	s.Router.Get("/", s.GetAllPorts)

	s.Router.Route(fmt.Sprintf("/%s", s.API), func(r chi.Router) {
		r.Mount("/api/ports", s.Router)
	})
}

func NewAPIServer(portService proto.PortServiceClient) *APIServer {
	s := &APIServer{
		Service: portService,
		Router:  chi.NewRouter(),
		API:     apiVer,
	}
	s.Routes()
	return s
}

type APIServer struct {
	Service proto.PortServiceClient
	Router  *chi.Mux
	API     string
}

type Port struct {
	portID      string
	name        string
	city        string
	country     string
	alias       []string
	regions     string
	coordinates latlng.LatLng
	province    string
	timezone    string
	unlocs      []string
	code        string
}
