package store

import (
	"fmt"
	"net/url"

	"github.com/shortedapp/shorted/services/watcher/pkg/index"
	"github.com/shortedapp/shorted/services/watcher/pkg/store/driver"
)

// Store represents the storage engine for the watcher index
type Index struct {
	driver.Driver
}

// Get retrieves the watch index from storage. An error is returned
// if the storage driver failed to fetch the index, or the
// index identified by the key, version pair does not exist.
func (s *Index) Get(id string) (*index.Watch, error) {
	return s.Driver.Get(id)
}

// Create creates a new storage entry holding the index. An
// error is returned if the storage driver failed to store the
// index, or a index with identical an key already exists.
// func (s *Storage) Create(idx *index.Watch) error {
// 	return s.Driver.Create(idx)
// }

// Update updates the index in storage. An error is returned if the
// storage backend fails to update the index or if the index
// does not exist.
func (s *Index) Update(id string, idx *index.Watch) error {
	return s.Driver.Update(id, idx)
}

func NewGCS(bucket string) (*Index, error) {
	d, err := driver.NewGCS(bucket)
	if err != nil {
		return &Index{}, fmt.Errorf("error initialising gcs storage driver: %v", err)
	}
	return &Index{
		Driver: d,
	}, nil
}
func New(u string) (*Index, error) {
	p, err := url.Parse(u)
	if err != nil {
		panic("unable to parse storage URL")
	}
	switch p.Scheme {
	case "gs":
		d, err := driver.NewGCS(u)
		if err != nil {
			return &Index{}, fmt.Errorf("error initialising gcs storage driver: %v", err)
		}
		return &Index{
			Driver: d,
		}, nil
	case "file":
		d, err := driver.NewFile(u)
		if err != nil {
			return &Index{}, fmt.Errorf("error initialising file storage driver: %v", err)
		}
		return &Index{
			Driver: d,
		}, nil
	default:
		panic("unknown storage driver")
	}
}