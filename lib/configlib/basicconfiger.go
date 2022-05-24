package configlib

import (
	"encoding/json"
	"errors"
	"flag"
	"os"
	"reflect"
	"strings"

	"github.com/MC2BP/MicroS-Go/lib/errorlib"
)

var (
	configFile = "./config/local.json"
	configEnv  = ""
)

func Init() {
	flag.StringVar(&configFile, "conf", "./config/local.json", "Specify the path to the configuration file")
	flag.StringVar(&configEnv, "envconf", "", "Specify the environment variable that contains the configuration")
}

func ReadConfig() Configer {
	flag.Parse()
	rawConfig, err := os.ReadFile(configFile)
	if errors.Is(err, os.ErrNotExist) {
		strConfig := os.Getenv(configEnv)
		if strConfig == "" {
			err = errorlib.Errf("Config from env '%s' was empty", configEnv)
			panic(err)
		}
		rawConfig = []byte(strConfig)
	} else if err != nil {
		err = errorlib.Errf("Failed to read config, err:", err)
		panic(err)
	}

	var config Config
	err = json.Unmarshal(rawConfig, &config)
	if err != nil {
		panic(err)
	}

	var interfaceConfig interface{}
	err = json.Unmarshal(rawConfig, &interfaceConfig)
	if err != nil {
		panic(err)
	}

	return &BasicConfiger{
		rawJson: rawConfig,
		config:  config,
		raw:     interfaceConfig,
	}
}

type BasicConfiger struct {
	rawJson []byte
	config  Config
	raw     interface{}
}

func (c *BasicConfiger) GetEnvironment() string {
	return c.config.Environment
}

func (c *BasicConfiger) GetApplicationID() int {
	return c.config.ID
}

func (c *BasicConfiger) GetWebserver() Webserver {
	return c.config.Webserver
}

func (c *BasicConfiger) GetService(service string) Service {
	return c.config.Services[service]
}

func (c *BasicConfiger) GetString(path string) string {
	v := c.GetObject(path)
	if v == nil {
		return ""
	}
	switch x := v.(type) {
	case string:
		return x
	default:
		return ""
	}
}

func (c *BasicConfiger) GetFloat(path string) float64 {
	v := c.GetObject(path)
	if v == nil {
		return 0
	}
	switch x := v.(type) {
	case float32:
		return float64(x)
	case float64:
		return x
	default:
		return 0
	}
}

func (c *BasicConfiger) GetInt(path string) int64 {
	v := c.GetObject(path)
	if v == nil {
		return 0
	}
	switch x := v.(type) {
	case int:
		return int64(x)
	case int64:
		return x
	default:
		return 0
	}
}

func (c *BasicConfiger) GetBool(path string) bool {
	v := c.GetObject(path)
	if v == nil {
		return false
	}
	switch x := v.(type) {
	case bool:
		return x
	default:
		return false
	}
}

func (c *BasicConfiger) GetSlice(path string) []interface{} {
	v := c.GetObject(path)
	if v == nil {
		return nil
	}
	switch x := v.(type) {
	case []interface{}:
		return x
	default:
		return nil
	}
}

func (c *BasicConfiger) GetObject(path string) interface{} {
	elements := strings.Split(path, ".")

	object := c.raw
	for _, element := range elements {
		r := reflect.ValueOf(object)
		f := r.FieldByName(element)
		if f.IsNil() {
			return ""
		}
		object = f.Interface()
	}
	return object
}
