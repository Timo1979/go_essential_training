package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

type DbData struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type Status struct {
	Message string `json:"message"`
}

var (
	Data  = map[string]interface{}{}
	mutex sync.Mutex
)

func setSynced(data *DbData) error {
	// fmt.Printf("Data is %+v (%T)\n", Data, Data)
	// fmt.Printf("data to store: %+v\n", data)
	// fmt.Printf("data.Key %v data.Value %v\n", data.Key, data.Value)

	mutex.Lock()
	defer mutex.Unlock()
	Data[data.Key] = data.Value
	return nil
}

func getSynced(key string) (*DbData, error) {
	// fmt.Printf("Data is %+v (%T)\n", Data, Data)
	mutex.Lock()
	value, ok := Data[key]
	mutex.Unlock()
	if !ok {
		return nil, fmt.Errorf("not found")
	}
	return &DbData{key, value}, nil
}

func writeResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		return err
	}
	return nil
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != "POST" {
		http.Error(w, fmt.Sprintf("Метод должен быть POST"), http.StatusBadRequest)
		return
	}
	dec := json.NewDecoder(r.Body)
	data := &DbData{}
	if err := dec.Decode(data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// writeResponse(w, Status{"value set"})
	writeResponse(w, &Status{"записано!"})
	setSynced(data)

}

func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		http.Error(w, fmt.Sprintf("Метод должен быть GET"), http.StatusBadRequest)
		return
	}
	splitted := strings.Split(r.RequestURI, "/")
	if (len(splitted) < 3) || (splitted[2] == "") {
		http.Error(w, fmt.Sprintf("Не передан ключ"), http.StatusBadRequest)
		return
	}
	data, err := getSynced(splitted[2])
	if err != nil {
		http.Error(w, fmt.Sprintf("не найдено"), http.StatusNotFound)
		return
	}
	writeResponse(w, data)
}

func main() {
	http.HandleFunc("/db", postHandler)
	http.HandleFunc("/db/", getHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
