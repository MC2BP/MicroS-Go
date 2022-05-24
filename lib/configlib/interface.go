package configlib

type Configer interface {
	GetEnvironment() string
	GetApplicationID() int
	GetWebserver() Webserver
	GetService(service string) Service
	GetString(path string) string
	GetFloat(path string) float64
	GetInt(path string) int64
	GetBool(path string) bool
	GetObject(path string) interface{}
	GetSlice(path string) []interface{}
}
