package store

import "sync"

// inMemoryStore is created for unit test
// no need to have persistent storage for unit tests
type inMemoryStore struct {
	cache string
	mutex sync.RWMutex
}

func (s *inMemoryStore) Get() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.cache
}

func (s *inMemoryStore) Set(value string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache = value
	return nil
}

func SetupInMemoryStore(value string) {
	s = &inMemoryStore{
		cache: value,
	}
}
