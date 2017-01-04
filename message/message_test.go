package message

import (
	"fmt"
	"testing"
)

func TestSerialize(t *testing.T) {
	test := &ElevationQuery{
		QueryList: []*ElevationQuery_LatLonAlt{{1.0, 2.0, 3.0}},
	}
	fmt.Println("test = ", test.QueryList[0].Latitude)
	if len(test.QueryList) != 5 {
		t.Errorf("Expected length = 1, instead go %d", len(test.QueryList))
	}
}
