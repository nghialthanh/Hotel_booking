package pkg

import (
	"fmt"
	"hotel-booking/pkg/db"
	"os"
)

type App struct {
	Name       string
	DBList     []*db.Client
	ServerList []ConfigServer
	hostname   string
}

func NewApplication(name string) *App {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "default-hostname"
	}

	app := &App{
		Name:       name,
		DBList:     []*db.Client{},
		ServerList: []ConfigServer{},
		hostname:   hostname,
	}
	return app
}

func (app *App) AddNewDB(config db.ConfigDB) {
	urlConnectDB := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.Username, config.Password, config.DBName)

	client := &db.Client{
		Name:   config.DBName,
		Url:    urlConnectDB,
		Config: &config,
	}

	app.DBList = append(app.DBList, client)
}

func (app *App) Run() error {
	if len(app.DBList) > 0 {
		for _, db := range app.DBList {
			err := db.Connect()
			if err != nil {
				fmt.Println("Connect DB " + db.Name + " error: " + err.Error())
				return err
			}
		}
	}
	return nil
}
