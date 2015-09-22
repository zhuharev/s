package shortener

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Service struct {
	Store  Store
	Config *Config
}

func New(filepath string) (*Service, error) {
	cfg, e := NewConfig(filepath)
	if e != nil {
		return nil, e
	}
	return NewFromConfig(cfg)
}

func NewFromConfig(cfg *Config) (*Service, error) {
	s := new(Service)
	s.Config = cfg

	db, e := NewLevelDbStore(s.Config.Database.DataDir)
	if e != nil {
		return nil, e
	}
	s.Store = db

	return s, nil
}

func (s *Service) Run() {
	if s.Config.Web.Http {
		log.Printf("Listen %d\n", s.Config.Web.Port)
		if e := http.ListenAndServe(":"+fmt.Sprint(s.Config.Web.Port), s); e != nil {
			panic(e)
		}
	}
}

func randStr(length int) string {
	rand.Seed(time.Now().UnixNano())
	src := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = src[rand.Intn(len(src))]
	}
	return string(result)
}
