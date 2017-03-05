package server

import (
	"errors"

	"github.com/acidleroy/elevation/message"
)

type DummyLookup struct {
	defaultValue float64
}

func (d *DummyLookup) Retrieve(m *message.ElevationQuery) (message.ElevationQuery, error) {
	for _, element := range m.GetQueryList() {
		element.Altitude = d.defaultValue
	}
	return *m, nil
}

func (d *DummyLookup) Name() string {
	return "dummy"
}

func NewDummyLookup() ElevationLookup {
	return &DummyLookup{defaultValue: 111}
}

type PostGISLookup struct {
}

func (d *PostGISLookup) Retrieve(m *message.ElevationQuery) (message.ElevationQuery, error) {
	return *m, errors.New("PostGIS Retrieve method not yet implemented.")
}

func (d *PostGISLookup) Name() string {
	return "postgis"
}

func NewPostGISLookup() ElevationLookup {
	return &PostGISLookup{}
}
