package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jauntyward/cbufd/cbuf"
)

type (
	//Point represents a time series point to be added to or taken from a cbuf
	Point struct {
		Series      string
		Measurement interface{}
	}

	//Request represents a request for a point
	Request struct {
		Series string
	}
)

var (
	buffers     map[string]cbuf.CircularBuffer
	errorLogger *log.Logger
)

func main() {
	buffers = make(map[string]cbuf.CircularBuffer)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/writePoint", WritePoint)

	errorLogger = log.New(os.Stdout,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	println("Starting CbufD")
	log.Fatal(http.ListenAndServe(":8080", router))

}

//WritePoint handles API requests to store a single time series sample
func WritePoint(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	point := new(Point)

	if err := json.Unmarshal(body, &point); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			errorLogger.Print(err)
		}
	} else {
		buffer, ok := buffers[point.Series]
		if !ok {
			buffers[point.Series] = *new(cbuf.CircularBuffer)
			buffer = buffers[point.Series]
		}
		buffer.Enqueue(point.Measurement)
	}
}

//GetPoint returns the oldest point from a cbuf series
func GetPoint(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	request := new(Request)

	if err := json.Unmarshal(body, &request); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			errorLogger.Print(err)
		}
	} else {
		buffer := buffers[request.Series]
		point := buffer.Dequeue()
		json.NewEncoder(w).Encode(point)
	}

}

//PeakPoint returns the oldest point from a cbuf series
func PeakPoint(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	request := new(Request)

	if err := json.Unmarshal(body, &request); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			errorLogger.Print(err)
		}
	} else {
		buffer := buffers[request.Series]
		point := buffer.Peak()
		json.NewEncoder(w).Encode(point)
	}

}
