package client

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/acidleroy/elevation/message"
	"github.com/golang/protobuf/proto"
)

type Client interface {
	BaseEndpoint() string
	RetrieveElevations(e *message.ElevationQuery) error
}

type HttpClient struct {
	address string
	port    uint32
}

func (c *HttpClient) BaseEndpoint() string {
	var buffer bytes.Buffer
	buffer.WriteString("http://")
	buffer.WriteString(c.address)
	buffer.WriteString(":")
	buffer.WriteString(strconv.Itoa(int(c.port)))
	buffer.WriteString("/")
	return buffer.String()
}

func (c *HttpClient) RetrieveElevations(e *message.ElevationQuery) message.ElevationQuery {
	log.Println("RetrieveElevations called!")
	var buffer bytes.Buffer
	buffer.WriteString(c.BaseEndpoint())
	buffer.WriteString("elevations/")
	data, err := proto.Marshal(e)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	var pbBuffer bytes.Buffer
	pbBuffer.Write(data)

	log.Printf("Making the following request: %s\n", buffer.String())

	res, err := http.Post(buffer.String(), "protobuf", &pbBuffer)
	if err != nil {
		log.Fatal(err)
	}
	elevations, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Received the following message: \n%s\n", elevations)
	elevationsRetrieved := &message.ElevationQuery{}
	err = proto.Unmarshal(elevations, elevationsRetrieved)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	return *elevationsRetrieved
}

func NewHttpClient(addr string, port uint32) *HttpClient {
	return &HttpClient{address: addr, port: port}
}
