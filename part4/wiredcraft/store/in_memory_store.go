package store

import "sync"

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
