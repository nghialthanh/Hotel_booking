package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hotel-booking/pkg"
	"hotel-booking/pkg/db"
	"os"
)

func main() {

	// Parse config env
	var configMap map[string]string

	config := os.Getenv("config")

	decoded, err := base64.StdEncoding.DecodeString(config)
	if err != nil {
		fmt.Println("Decode config error: " + err.Error())
		return
	}
	err = json.Unmarshal(decoded, &configMap)
	if err != nil {
		fmt.Println("Parse JSON config string error: " + err.Error())
		return
	}

	app := pkg.NewApplication("hotel-booking")

	app.AddNewDB(db.ConfigDB{
		Host:     configMap["POSTGRES_HOST"],
		Username: configMap["POSTGRES_USER"],
		Password: configMap["POSTGRES_PASSWORD"],
		DBName:   configMap["POSTGRES_DB"],
		Port:     configMap["POSTGRES_PORT"],
		SSL:      false,
	})

	app.Run()

}
