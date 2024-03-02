package config

import "time"

// App config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	// Redis    RedisConfig
	MongoDB MongoDB
	Cookie  Cookie
	// Store    Store
	Session Session
	// Metrics  Metrics
	// Logger   Logger
	// AWS      AWS
	// Jaeger   Jaeger
}

// Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

// Postgresql config
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PgDriver           string
}

// MongoDB config
type MongoDB struct {
	MongoURI string
}

// Cookie config
type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

// Session config
type Session struct {
	Prefix string
	Name   string
	Expire int
}

// // Load config file from given path
// func LoadConfig(filename string) (*viper.Viper, error) {
// 	v := viper.New()

// 	v.SetConfigName(filename)
// 	v.AddConfigPath(".")
// 	v.AutomaticEnv()
// 	if err := v.ReadInConfig(); err != nil {
// 		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
// 			return nil, errors.New("config file not found")
// 		}
// 		return nil, err
// 	}

// 	return v, nil
// }
