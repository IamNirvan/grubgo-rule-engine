package config

import (
	"os"
	"path"
	"strings"

	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Log struct {
	Level   string `mapstructure:"level" json:"level"`
	Methods bool   `mapstructure:"methods" json:"methods"`
}

type WebServer struct {
	Host    string `mapstructure:"host" json:"host"`
	Port    uint16 `mapstructure:"port" json:"port"`
	Timeout uint16 `mapstructure:"timeout" json:"timeout"`
}

type Database struct {
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Dbname   string `mapstructure:"dbname" json:"dbname"`
	Sslmode  string `mapstructure:"sslmode" json:"sslmode"`
}

type Config struct {
	Log       Log       `mapstructure:"log"`
	WebServer WebServer `mapstructure:"web_server"`
	Database  Database  `mapstructure:"database"`
}

func New() *Config {
	return &Config{
		// These will be overriden by config file or environment variables
		Log: Log{
			Level:   "info",
			Methods: true,
		},
		WebServer: WebServer{
			Host:    "localhost",
			Port:    8080,
			Timeout: 15,
		},
		Database: Database{
			User:     "postgres",
			Password: "postgres",
			Host:     "localhost",
			Port:     5432,
			Dbname:   "postgres",
			Sslmode:  "disable",
		},
	}
}

func LoadConfig() *Config {
	vpr := viper.NewWithOptions(viper.EnvKeyReplacer(strings.NewReplacer(".", "_")))

	vpr.SetConfigName("config")
	vpr.SetConfigType("yaml")

	vpr.AddConfigPath("./config")
	vpr.AddConfigPath(".")

	exe, err := os.Executable()
	if err != nil {
		log.Fatalf("cannot access executable directory, error : %v", err)
	}
	vpr.AddConfigPath(path.Dir(exe))

	if err := vpr.ReadInConfig(); err != nil {
		log.Fatalf("config file reading error : %v", err)
	}

	vpr.SetEnvPrefix("RULE_ENGINE")
	vpr.AutomaticEnv()

	config := New()
	if err := vpr.Unmarshal(config, viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(mapstructure.TextUnmarshallerHookFunc()))); err != nil {
		log.Fatalf("failed to decode configuration to struct with error: %v", err)
	}

	var level log.Level
	if err := level.UnmarshalText([]byte(config.Log.Level)); err != nil {
		log.Fatal("log level not valid, error:", err)
	}
	log.SetLevel(level)
	log.SetReportCaller(config.Log.Methods)

	return config
}
