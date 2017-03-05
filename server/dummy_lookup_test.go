package server

import (
	"testing"

	"github.com/acidleroy/elevation/message"
)

func TestDummyElevationLookup(t *testing.T) {
	lookup := DummyLookup{defaultValue: 111}
	query := message.ElevationQuery{}
	list := &query.QueryList
	*list = append(*list, &message.ElevationQuery_LatLonAlt{Latitude: 1.0, Longitude: 2.0, Altitude: 0.0})
	*list = append(*list, &message.ElevationQuery_LatLonAlt{Latitude: 11.0, Longitude: 22.0, Altitude: 0.0})
	result, _ := lookup.Retrieve(&query)

	for _, element := range result.GetQueryList() {
		if element.Altitude != lookup.defaultValue {
			t.Errorf("Expected altitude to be %f, but instead got %f.", lookup.defaultValue, element.Altitude)
		}
	}

}
