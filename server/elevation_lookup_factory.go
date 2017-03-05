package server

import (
	"errors"
	"strings"

	"github.com/acidleroy/elevation/message"
)

// Interface used to look up elevations with a query message. With this interface,
// we can program looking up the elevation via SQL, Redis or other implementations
// that we decide will work well.
type ElevationLookup interface {
	// Retrieve elevations given an elveation query.
	// The input should only contain a list of latitudes and longitudes
	// and the output will container the elevations corresponding to each one
	// of those longitudes and latitudes. If there is any problem with the
	// query, then an error is returned.
	Retrieve(m *message.ElevationQuery) (message.ElevationQuery, error)

	Name() string
}

func ElevationLookupFactory(lookupType string) (ElevationLookup, error) {
	lookupType = strings.ToLower(lookupType)
	switch lookupType {
	case "dummy":
		return NewDummyLookup(), nil
	case "postgis":
		return NewPostGISLookup(), nil
	default:
		return nil, errors.New("Unknown elevation type requested!")
	}
}
