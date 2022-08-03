package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(r *http.Request, target interface{}) error {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), target); err == nil {
			return nil
		} else {
			return fmt.Errorf("error while parsing body: %v", err)
		}
	} else {
		return fmt.Errorf("error while reading body: %v", err)
	}
}

func ParseUint(r *http.Request, key string) (uint64, error) {
	params := mux.Vars(r)
	if value, ok := params[key]; ok {
		if id, err := strconv.ParseUint(value, 10, 64); err == nil {
			return id, nil
		} else {
			return 0, fmt.Errorf("error while parsing %s: %v", key, err)
		}
	}
	return 0, fmt.Errorf("%s not found in params", key)
}

func ParseUint64(payload interface{}) (uint64, error) {
	if value, ok := payload.(float64); ok {
		return uint64(value), nil
	} else {
		return 0, fmt.Errorf("error while parsing %v", payload)
	}
}

func Contains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Write([]byte(message))
}

func GenerateResponseMap(message string) map[string]string {
	return map[string]string{"message": message}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
