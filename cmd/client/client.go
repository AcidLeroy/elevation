package main

import (
	"log"

	"github.com/acidleroy/elevation/client"
	"github.com/acidleroy/elevation/message"
)

func main() {
	conn := client.NewHttpClient("localhost", 8080)
	query := message.ElevationQuery{}
	list := &query.QueryList
	*list = append(*list, &message.ElevationQuery_LatLonAlt{Latitude: 1.0, Longitude: 2.0, Altitude: 0.0})
	*list = append(*list, &message.ElevationQuery_LatLonAlt{Latitude: 11.0, Longitude: 22.0, Altitude: 0.0})
	log.Printf("Querying for the following values: ")
	testList := query.GetQueryList()
	for _, element := range testList {
		log.Printf("Latitude = %f, Longitude = %f", element.Latitude, element.Longitude)
	}

	log.Println("Calling RetrieveElevations...")
	results := conn.RetrieveElevations(&query)

	for _, element := range results.GetQueryList() {
		log.Printf("Latitude = %f, Longitude = %f, Altitude = %f", element.Latitude, element.Longitude, element.Altitude)
	}

	log.Println("Done!")
}
