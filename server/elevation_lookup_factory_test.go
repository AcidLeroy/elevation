package server

import "testing"

func TestElevationFactoryLookup(t *testing.T) {

	validFactoryMethods := []string{"dummy", "postgis"}

	for _, method := range validFactoryMethods {
		lookup, err := ElevationLookupFactory(method)
		if err != nil {
			t.Errorf("Could not create valid field %s: %s", method, err)
		}
		if method != lookup.Name() {
			t.Errorf("Expected %s type, but got %s", method, lookup.Name())
		}
	}
}
