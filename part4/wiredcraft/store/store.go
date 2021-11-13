package store

import (
	"fmt"
	"os"
	"sync"
	"wiredcraft/logger"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

// Store interface
// Get return the cache name, cache name init by storage
// Set if name is not equal to cached value, set to storage
type Store interface {
	Get() string
	Set(string) error
}

const NAME_HOLDER_KEY = "name_holder"
const NOT_FOUND_ERROR = "redigo: nil returned"

var s Store

type store struct {
	redisConn redis.Conn
	mutex     sync.RWMutex
	value     string
}

func (s *store) Get() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	return s.value
}

func (s *store) Set(newValue string) error {
	if s.value != newValue {
		s.mutex.Lock()
		defer s.mutex.Unlock()

		_, err := s.redisConn.Do("SET", NAME_HOLDER_KEY, newValue)
		if err != nil {
			logger.Logger().Error("error set to redis", zap.Any("error", err))
			return err
		}

		s.value = newValue
	}

	return nil
}

func InitValue() {
	var err error
	ss := &store{}

	password := "123456"
	host := "172.17.0.1"
	port := "6379"

	if os.Getenv("REDISPASS") != "" {
		password = os.Getenv("REDISPASS")
	}

	if os.Getenv("REDISHOST") != "" {
		host = os.Getenv("REDISHOST")
	}

	if os.Getenv("REDISHOST") != "" {
		port = os.Getenv("REDISPORT")
	}

	options := redis.DialPassword(password)
	ss.redisConn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", host, port), options)

	if err != nil {
		logger.Logger().Error("error connect to redis", zap.Any("error", err))
		os.Exit(1)
	}

	ss.value, err = redis.String(ss.redisConn.Do("GET", NAME_HOLDER_KEY))
	if err != nil && err.Error() != NOT_FOUND_ERROR {
		logger.Logger().Error("error query redis", zap.Any("error", err))
		os.Exit(1)
	}
	s = ss
}

func GetStore() Store {
	return s
}
