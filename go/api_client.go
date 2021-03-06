/*
 * Swagger Clients server
 *
 * This is a sample server Clients server
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var repository = NewInMemoryRepo()

func AddClient(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	client := Client{}

	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println(err)
		http.Error(w, "Error decoding request object", http.StatusBadRequest)
		return
	}
	repository.Create(&client)
	res, _ := json.Marshal(client)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res)
}

func DeleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	clientId, err := strconv.Atoi(vars["clientId"])
	if err != nil {
		log.Println("Invalid client Id:", err)
		http.Error(w, "Invalid client Id", http.StatusBadRequest)
		return
	}
	repository.Delete(int64(clientId))

	w.WriteHeader(http.StatusNoContent)
	_, _ = w.Write([]byte{})
}

func GetClientById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	clientId, err := strconv.Atoi(vars["clientId"])
	if err != nil {
		log.Println("Invalid client Id:", err)
		http.Error(w, "Invalid client Id", http.StatusBadRequest)
		return
	}

	client := repository.FindById(int64(clientId))
	res, _ := json.Marshal(client)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func GetClients(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	clients := repository.FindAll()
	res, _ := json.Marshal(clients)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	client := Client{}
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		log.Println(err)
		http.Error(w, "Error decoding request object", http.StatusBadRequest)
		return
	}

	repository.Update(&client)

	res, _ := json.Marshal(client)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
