package index

import (
	"path"
	"sort"
	"time"

	"github.com/blang/semver"
	"github.com/shortedapp/shorted/services/watcher/internal/urlutil"
)

type Source struct {
	Name         string               `json:"name"`
	URL          string               `json:"url"`
	entries      map[string]Documents `json:"entries"`
	entriesCount int                  `json:"count"`
}
type SourceOption func(*Source)

func NewSource(opts ...SourceOption) *Source {
	s := &Source{
		entries: map[string]Documents{},
	}
	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *House as the argument
		opt(s)
	}
	return s
}

func WithSourceName(n string) SourceOption {
	return func(s *Source) {
		s.Name = n
	} 
}

func WithSourceURL(u string) SourceOption {
	return func(s *Source) {
		s.URL = u
	}
}

// Add entry into index
func (s *Source) Add(md *Metadata, filename, baseURL, digest string) {
	u := filename
	if baseURL != "" {
		var err error
		u, err = urlutil.URLJoin(baseURL, filename)
		if err != nil {
			u = path.Join(baseURL, filename)
		}
	}
	d := &Document{
		URL:      u,
		Metadata: md,
		Digest:   digest,
		Created:  time.Now(),
	}
	// Check if entry already exists, if so append otherwise set
	if ee, ok := s.entries[md.Name]; !ok {
		d.Version = semver.MustParse("1.0.0").String()
		s.entries[md.Name] = Documents{d}
	} else {
		s.entries[md.Name] = append(ee, d)
	}
}

func (s *Source) AddDocument(name string, d Documents) {
	if ee, ok := s.entries[name]; !ok {
		s.entries[name] = d
	} else {
		for _, doc := range d {
			s.entries[name] = append(ee, doc)
		}
	}
}

func (s Source) SortEntries() {
	for _, versions := range s.entries {
		sort.Sort(sort.Reverse(versions))
	}
}
