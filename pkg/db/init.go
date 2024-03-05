package db

import (
	"database/sql"
	"fmt"
)

type ConfigDB struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
	SSL      bool
}

type Client struct {
	Name   string
	Url    string
	Config *ConfigDB
}

func (c *Client) Connect() error {
	db, err := sql.Open("postgres", c.Url)
	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		return err
	}

	fmt.Println("Successfully connected!", c.Name)
	return nil
}
