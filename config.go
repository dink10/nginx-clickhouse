package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/Sirupsen/logrus"
	"flag"
	"os"
	"strconv"
)

type Config struct {
	Settings struct {
		Interval int `yaml:"interval"`
		LogPath string `yaml:"log_path"`
	} `yaml:"settings"`
	ClickHouse struct {
		Db string `yaml:"db"`
		Table string `yaml:"table"`
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		Columns map[string]string `yaml:"columns"`
		Credentials struct{
			User string `yaml:"user"`
			Password string `yaml:"password"`
		} `yaml:"credentials"`
	} `yaml:"clickhouse"`
	Nginx struct {
		LogType string `yaml:"log_type"`
		LogFormat string `yaml:"log_format"`
	}
}

var configPath string

func init() {

	flag.StringVar(&configPath,"config_path", "config/config.yaml", "Config path.")

	flag.Parse()
}

func readConfig() *Config {

	config := Config{}

	logrus.Info("Reading config file: " + configPath)

	data, err := ioutil.ReadFile(configPath)

	if err != nil {
		logrus.Fatal("Config open error: ", err)
	}

	err = yaml.Unmarshal(data, &config)

	if err != nil {
		logrus.Fatal("Config read & unmarshal error: ", err)
	}

	return &config
}

func (c *Config) setEnvVariables() {

	// Settings

	if os.Getenv("LOG_PATH") != "" {
		c.Settings.LogPath = os.Getenv("LOG_PATH")
	}

	if os.Getenv("FLUSH_INTERVAL") != "" {
		c.Settings.Interval, _ = strconv.Atoi(os.Getenv("FLUSH_INTERVAL"))
	}

	// ClickHouse

	if os.Getenv("CLICKHOUSE_HOST") != "" {
		c.ClickHouse.Host = os.Getenv("CLICKHOUSE_HOST")
	}

	if os.Getenv("CLICKHOUSE_PORT") != "" {
		c.ClickHouse.Port = os.Getenv("CLICKHOUSE_PORT")
	}

	if os.Getenv("CLICKHOUSE_DB") != "" {
		c.ClickHouse.Db = os.Getenv("CLICKHOUSE_DB")
	}

	if os.Getenv("CLICKHOUSE_TABLE") != "" {
		c.ClickHouse.Table = os.Getenv("CLICKHOUSE_TABLE")
	}

	if os.Getenv("CLICKHOUSE_USER") != "" {
		c.ClickHouse.Credentials.User = os.Getenv("CLICKHOUSE_USER")
	}

	if os.Getenv("CLICKHOUSE_PASSWORD") != "" {
		c.ClickHouse.Credentials.Password = os.Getenv("CLICKHOUSE_PASSWORD")
	}

	// Nginx

	if os.Getenv("NGINX_LOG_TYPE") != "" {
		c.Nginx.LogType = os.Getenv("NGINX_LOG_TYPE")
	}

	if os.Getenv("NGINX_LOG_FORMAT") != "" {
		c.Nginx.LogFormat = os.Getenv("NGINX_LOG_FORMAT")
	}
}
