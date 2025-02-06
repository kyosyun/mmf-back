package env

import (
	"sync"

	"github.com/caarlos0/env/v6"
)

var (
	_cfg *Environment
	once sync.Once
)

type Environment struct {
	DBConfig     DBConfig
	AmazonConfig AmazonConfig
}

type DBConfig struct {
	Host string `env:"DB_HOST,required"`
	Port string `env:"DB_PORT,required"`
	User string `env:"DB_USER,required"`
	Pass string `env:"DB_PASS,required"`
	Name string `env:"DB_NAME,required"`
}

type AmazonConfig struct {
	AccessKey  string `env:"AMAZON_ACCESS_KEY,required"`
	SecretKey  string `env:"AMAZON_SECRET_KEY,required"`
	PartnerTag string `env:"AMAZON_PARTNER_TAG,required"`
}

func NewEnvironment() *Environment {
	once.Do(func() {
		_cfg = &Environment{}

		if err := env.Parse(_cfg); err != nil {
			panic(err)
		}
	})
	return _cfg
}

func GetEnvironment() *Environment {
	if _cfg == nil {
		panic("environment not initialized")
	}
	return _cfg
}
