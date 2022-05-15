package httplib

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetString(p string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[p]
}

func GetFloat(p string, r *http.Request) float64 {
	vars := mux.Vars(r)
	v, err := strconv.ParseFloat(vars[p], 64)
	if err != nil {
		return 0
	}
	return v
}

func GetInt(p string, r *http.Request) int {
	vars := mux.Vars(r)
	v, err := strconv.Atoi(vars[p])
	if err != nil {
		return 0
	}
	return v
}

func GetBool(p string, r *http.Request) bool {
	vars := mux.Vars(r)
	v, err := strconv.ParseBool(vars[p])
	if err != nil {
		return false
	}
	return v
}

func GetParameterString(p string, r *http.Request) string {
	query := r.URL.Query()
	v, has := query[p]
	if has && len(v) > 0 {
		return v[0]
	}
	return ""
}

func GetParameterFloat(p string, r *http.Request) float64 {
	v := GetParameterString(p, r)
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0
	}
	return f
}

func GetParameterInt(p string, r *http.Request) int {
	v := GetParameterString(p, r)
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}
	return i
}

func GetParameterBool(p string, r *http.Request) bool {
	v := GetParameterString(p, r)
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return b
}
