package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"

	v1 "github.com/silverspase/portService/clientAPI/internal/api/v1"
)

func TestGetAPort(t *testing.T) {
	s, rr, req := mockServer(t, "GET", "/AEAUH")
	s.Router.ServeHTTP(rr, req)
	expected, _ := MockPortServiceClient{}.Read(context.TODO(), &v1.ReadRequest{Api: "v1", PortID: "AEAUH"}, nil)
	res, _ := json.Marshal(expected)
	assert.Equal(t, string(res)+"\n", rr.Body.String())
}

func TestDeletePort(t *testing.T) {
	s, rr, req := mockServer(t, "DELETE", "/AEAUH")
	s.Router.ServeHTTP(rr, req)
	expected, _ := MockPortServiceClient{}.Delete(context.TODO(), &v1.DeleteRequest{Api: "v1", PortID: "AEAUH"}, nil)
	res, _ := json.Marshal(expected)
	assert.Equal(t, string(res)+"\n", rr.Body.String())
}

func TestReadAll(t *testing.T) {
	s, rr, req := mockServer(t, "GET", "/")
	s.Router.ServeHTTP(rr, req)
	expected, _ := MockPortServiceClient{}.ReadAll(context.TODO(), &v1.ReadAllRequest{Api: "v1"}, nil)
	res, _ := json.Marshal(expected)
	assert.Equal(t, string(res)+"\n", rr.Body.String())
}

func mockServer(t *testing.T, method, url string) (*APIServer, *httptest.ResponseRecorder, *http.Request) {
	c := MockPortServiceClient{}
	s := NewAPIServer(c)
	rr := httptest.NewRecorder()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Fatal(err)
	}
	return s, rr, req
}

//MockPortServiceClient is used to mock GRPC calls
type MockPortServiceClient struct{}

func (c MockPortServiceClient) Create(ctx context.Context, in *v1.CreateRequest, opts ...grpc.CallOption) (*v1.CreateResponse, error) {
	return &v1.CreateResponse{
		Api:    "",
		PortID: "",
	}, nil
}

func (c MockPortServiceClient) LoadFromJSON(ctx context.Context, in *v1.LoadFromJSONRequest, opts ...grpc.CallOption) (*v1.LoadFromJSONResponse, error) {
	return &v1.LoadFromJSONResponse{
		Api:     "",
		Created: 0,
		Updated: 0,
		Skipped: 0,
		Records: 0,
	}, nil
}

func (c MockPortServiceClient) Read(ctx context.Context, in *v1.ReadRequest, opts ...grpc.CallOption) (*v1.ReadResponse, error) {
	return &v1.ReadResponse{
		Api: "v1",
		Port: &v1.Port{
			PortID: in.PortID,
			Name:   "Santa Monika",
		},
	}, nil
}

func (c MockPortServiceClient) Delete(ctx context.Context, in *v1.DeleteRequest, opts ...grpc.CallOption) (*v1.DeleteResponse, error) {
	return &v1.DeleteResponse{
		Api:    "v1",
		PortID: in.PortID,
	}, nil
}

func (c MockPortServiceClient) ReadAll(ctx context.Context, in *v1.ReadAllRequest, opts ...grpc.CallOption) (*v1.ReadAllResponse, error) {
	return &v1.ReadAllResponse{
		Api: "",
		Ports: []*v1.Port{{
			PortID: "AEAUH",
			Name:   "Abu Dhabi",
		}, {
			PortID: "AEDXB",
			Name:   "Dubai",
		}},
	}, nil
}
