package data

import (
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gilgames000/go-noskit/datastore"
)

var _ datastore.RawMapLoader = RawMapLoader{}

type RawMapLoader struct {
	mapsDirectory string
}

func NewRawMapLoader(mapsDirectory string) *RawMapLoader {
	return &RawMapLoader{mapsDirectory: mapsDirectory}
}

func (l RawMapLoader) Load(mapID int) (io.Reader, error) {
	f, err := os.Open(filepath.Join(l.mapsDirectory, strconv.Itoa(mapID)))
	if err != nil {
		return nil, err
	}

	return f, nil
}
