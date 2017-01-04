package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/acidleroy/elevation/message"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
)

// Function that starts the elevation service endpoint. The only valid
// url at this point is a POST method to "/elevations/" endpoint
func Start() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/elevations/", GetElevationsEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// The main page.
//
// This is for a simple page to validate that the DEM server is running
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the DEM service!")
}

func GetElevationsEndpoint(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received a request for elevations!")
	log.Printf("The type of the message is: %s", r.Method)
	dataBuffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("There was an error reading the databuffer!", err)
	}

	log.Printf("The size of the body is: %d", r.ContentLength)

	elevationsRetrieved := &message.ElevationQuery{}
	err = proto.Unmarshal(dataBuffer, elevationsRetrieved)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Printf("Asking for %d elevations.", len(elevationsRetrieved.GetQueryList()))
	log.Println("Asking for the following Latitudes and Longitudes: ")
	for _, latlonalt := range elevationsRetrieved.GetQueryList() {
		log.Printf("Latitude = %f, Longitude = %f", latlonalt.Latitude, latlonalt.Longitude)
		latlonalt.Altitude = 666 // assign dummy value for now.
	}

	// Marshall new data
	data, err := proto.Marshal(elevationsRetrieved)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	//	fmt.Fprint(w, data)
	w.Write(data)
}
