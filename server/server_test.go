package server

import (
	"testing"

	"github.com/acidleroy/elevation/message"
)

// Interface used to look up elevations with a query message
type ElevationLookup interface {
	// Retrieve elevations given an elveation query.
	// The input should only contain a list of latitudes and longitudes
	// and the output will container the elevations corresponding to each one
	// of those longitudes and latitudes. If there is any problem with the
	// query, then an error is returned.
	Retrieve(m *message.ElevationQuery) (message.ElevationQuery, error)
}

type DummyLookup struct{}

func (d *DummyLookup) Retrieve(m *message.ElevationQuery) (message.ElevationQuery, error) {

}

func TestDummyElevationLookup(t *testing.T) {
}
