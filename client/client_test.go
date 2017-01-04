package client

import "testing"

func TestBaseEndpoint(t *testing.T) {
	hostname := "localhost"
	var port uint32 = 8080
	conn := NewHttpClient(hostname, port)
	if conn.address != "localhost" {
		t.Errorf("expected address to be %s, but instead go %s.", hostname, conn.address)
	}

	if conn.port != port {
		t.Errorf("Expected port to be %d, but instead got %d.", port, conn.port)
	}
	expectedName := "http://localhost:8080/"
	address := conn.BaseEndpoint()

	if expectedName != address {
		t.Errorf("Expected the addres to be %s, but instead got %s", expectedName, address)
	}
}
