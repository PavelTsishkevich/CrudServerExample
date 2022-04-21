package main

import (
	swagger "CrudServerExample/go"
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type handleResult struct {
	method     string
	path       string
	body       []byte
	statusCode int
	handler    func(http.ResponseWriter,
		*http.Request)
	result string
}

var handleResults = []handleResult{
	{"GET", "/v2/", nil, 200, swagger.Index, "Hello World!"},
	{"POST", "/v2/clients", []byte(`{"id": "10","username": "TestLogin", "firstName":"TestName", "lastName":"TestLastName", "email":"testEmail", "phone":"testPhone"}`), 400, swagger.AddClient, "Error decoding request object\n"},
	{"POST", "/v2/clients", []byte(`{"id": 10,"username": 10, "firstName":"TestName", "lastName":"TestLastName", "email":"testEmail", "phone":"testPhone"}`), 400, swagger.AddClient, "Error decoding request object\n"},
	{"POST", "/v2/clients", []byte(`{"id": 10,"username": "TestLogin", "firstName":10, "lastName":"TestLastName", "email":"testEmail", "phone":"testPhone"}`), 400, swagger.AddClient, "Error decoding request object\n"},
	{"POST", "/v2/clients", []byte(`{"id": 10,"username": "TestLogin", "firstName":"TestName", "lastName":10, "email":"testEmail", "phone":"testPhone"}`), 400, swagger.AddClient, "Error decoding request object\n"},
	{"POST", "/v2/clients", []byte(`{"id": 10,"username": "TestLogin", "firstName":"TestName", "lastName":"TestLastName", "email":10, "phone":"testPhone"}`), 400, swagger.AddClient, "Error decoding request object\n"},
	{"POST", "/v2/clients", []byte(`{"id": 10,"username": "TestLogin", "firstName":"TestName", "lastName":"TestLastName", "email":"testEmail", "phone":10}`), 400, swagger.AddClient, "Error decoding request object\n"},

	{"POST", "/v2/clients", []byte(`{"id": 10,"username": "TestLogin", "firstName":"TestName", "lastName":"TestLastName", "email":"testEmail", "phone":"testPhone"}`), 201, swagger.AddClient, "{\"id\":1,\"username\":\"TestLogin\",\"firstName\":\"TestName\",\"lastName\":\"TestLastName\",\"email\":\"testEmail\",\"phone\":\"testPhone\"}"},
	{"GET", "/v2/clients/1", nil, 200, swagger.GetClientById, "{\"id\":1,\"username\":\"TestLogin\",\"firstName\":\"TestName\",\"lastName\":\"TestLastName\",\"email\":\"testEmail\",\"phone\":\"testPhone\"}"},
	{"GET", "/v2/clients/a", nil, 400, swagger.GetClientById, "Invalid client Id\n"},

	{"GET", "/v2/clients", nil, 200, swagger.GetClients, "[{\"id\":1,\"username\":\"TestLogin\",\"firstName\":\"TestName\",\"lastName\":\"TestLastName\",\"email\":\"testEmail\",\"phone\":\"testPhone\"}]"},
	{"PUT", "/v2/clients", []byte(`{"id": 1,"username": "TestLogin2", "firstName":"TestName2", "lastName":"TestLastName2", "email":"testEmail2", "phone":"testPhone2"}`), 200, swagger.UpdateClient, "{\"id\":1,\"username\":\"TestLogin2\",\"firstName\":\"TestName2\",\"lastName\":\"TestLastName2\",\"email\":\"testEmail2\",\"phone\":\"testPhone2\"}"},
	{"PUT", "/v2/clients", []byte(`{"id": "10","username": "TestLogin", "firstName":"TestName", "lastName":"TestLastName", "email":"testEmail", "phone":"testPhone"}`), 400, swagger.UpdateClient, "Error decoding request object\n"},
	{"PUT", "/v2/clients", []byte(`{"id": 10,"username": 10, "firstName":"TestName", "lastName":"TestLastName", "email":"testEmail", "phone":"testPhone"}`), 400, swagger.UpdateClient, "Error decoding request object\n"},
	{"PUT", "/v2/clients", []byte(`{"id": 10,"username": "TestLogin", "firstName":10, "lastName":"TestLastName", "email":"testEmail", "phone":"testPhone"}`), 400, swagger.UpdateClient, "Error decoding request object\n"},
	{"PUT", "/v2/clients", []byte(`{"id": 10,"username": "TestLogin", "firstName":"TestName", "lastName":10, "email":"testEmail", "phone":"testPhone"}`), 400, swagger.UpdateClient, "Error decoding request object\n"},
	{"PUT", "/v2/clients", []byte(`{"id": 10,"username": "TestLogin", "firstName":"TestName", "lastName":"TestLastName", "email":10, "phone":"testPhone"}`), 400, swagger.UpdateClient, "Error decoding request object\n"},

	{"DELETE", "/v2/clients/a", nil, 400, swagger.DeleteClient, "Invalid client Id\n"},
	{"DELETE", "/v2/clients/1", nil, 204, swagger.DeleteClient, ""},
}

func TestHandlers(t *testing.T) {
	for _, res := range handleResults {
		req, err := http.NewRequest(res.method, res.path, bytes.NewBuffer(res.body))
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := swagger.NewRouter()
		router.HandleFunc(res.path, res.handler)
		router.ServeHTTP(rr, req)

		if rr.Code != res.statusCode {
			t.Errorf("handler should have failed on routeVariable: got %v want %v",
				rr.Code, res.statusCode)
		}
		body, _ := ioutil.ReadAll(rr.Body)

		if output := string(body); output != res.result {
			t.Errorf("handler should return another result: got %v want %v",
				output, res.result)
		}
	}

}
