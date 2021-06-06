package models

import (
	"encoding/json"
	"io/ioutil"
)

type httpconf struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}

type dbconf struct {
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	Password string `json:"password,omitempty"`
	User     string `json:"username,omitempty"`
	DBName   string `json:"dbname,omitempty"`
}

type Config struct {
	HTTP httpconf `json:"server,omitempty"`
	DB   dbconf   `json:"database,omitempty"`
}

func (c *Config) Load(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return err
	}
	return nil
}
