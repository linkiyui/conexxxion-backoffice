package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	config      *Configuration
	initialized = false
)

type Configuration struct {
	environment      string
	LogsFile         string `yaml:"logs_file"`
	TemplatesPath    string `yaml:"templates_path"`
	Verbose          bool   `yaml:"verbose"`
	LogLevel         string `yaml:"log_level"`
	APIListenPort    int    `yaml:"api_listen_port"`
	GRPCListenPort   int    `yaml:"grpc_listen_port"`
	MetricsPort      int    `yaml:"metrics_port"`
	TranslationsPath string `yaml:"translations_path"`
	Database         `yaml:"database"`
	JWTOptions       `yaml:"jwt_options"`
}

func (c Configuration) IsDev() bool {
	return c.environment == "dev"
}

type Database struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	Username       string `yaml:"username"`
	Password       string `yaml:"password"`
	DatabaseName   string `yaml:"db_name"`
	MaxConnections int32  `yaml:"max_connections"`
}

type JWTOptions struct {
	Secret                string `yaml:"secret"`
	SignUpExpirationTime  int    `yaml:"signup_expiration_time"`
	LoginExpirationTime   int    `yaml:"login_expiration_time"`
	RefreshExpirationTime int    `yaml:"refresh_expiration_time"`
}

func Init(configFilePath string) {
	if initialized {
		return
	}
	defer func() {
		initialized = true
	}()
	config = defaultConfig()
	confFile := os.Getenv("CNXXXN_CONFIG_FILE")
	if configFilePath != "" {
		confFile = configFilePath
	}
	if confFile == "" {
		return
	}
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("reading config file: %s | error: %s\n", confFile, err.Error())
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("parsing config file: %s | error: %s\n", confFile, err.Error())
	}
}

func defaultConfig() *Configuration {
	c := &Configuration{
		environment:    "prod",
		LogsFile:       "",
		LogLevel:       "info",
		TemplatesPath:  "./static/templates",
		APIListenPort:  8080,
		GRPCListenPort: 9000,
		MetricsPort:    9090,

		TranslationsPath: "./static/translations",
		Database: Database{
			Host:           "127.0.0.1",
			Port:           5432,
			Username:       "postgres",
			Password:       "sasa",
			DatabaseName:   "conexxxion-backoffice",
			MaxConnections: 10,
		},

		JWTOptions: JWTOptions{
			Secret:                "supersecret",
			LoginExpirationTime:   12,
			RefreshExpirationTime: 3,
		},
	}
	if os.Getenv("CNXXXN_ENV") == "dev" {
		c.environment = "dev"
	}
	return c
}

func GetConfig() Configuration {
	if config == nil {
		config = defaultConfig()
	}
	return *config
}
